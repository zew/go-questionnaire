package qst

import (
	"fmt"
	"html/template"
	"strings"
	"time"

	"github.com/zew/go-questionnaire/cfg"
	"github.com/zew/go-questionnaire/trl"
)

// ParamT contains changing parameters to a questionnaire
type ParamT struct {
	Name string `json:"name,omitempty"` // i.e. main_refinance_rate_ecb
	Val  string `json:"val,omitempty"`  // i.e. "01.02.2018: 3.2%"
	// Challenge string `json:"challenge,omitempty"` // i.e.  Set field 'main_refinance_rate_ecb' to `01.02.2018: 3.2%` as in `main refinance rate of the ECB (01.02.2018: 3.2%)`"
}

// Param returns the value of a questionnaire param
func (s *surveyT) Param(name string) (string, error) {
	for _, p := range s.Params {
		if p.Name == name {
			return p.Val, nil
		}
	}
	return "?", fmt.Errorf("Param %q not found", name)
}

// surveyT stores the interval components of a questionnaire wave.
// For quarterly intervals, it needs to be extended
type surveyT struct {
	Type    string `json:"type,omitempty"`    // The type identifier, i.e. "fmt" or "cep"
	Variant string `json:"variant,omitempty"` // A variation

	Org  trl.S `json:"org,omitempty"`  // organization, i.e. Unicef
	Name trl.S `json:"name,omitempty"` // full name, i.e. programming languages survey

	Year  int        `json:"year,omitempty"`  // The wave year
	Month time.Month `json:"month,omitempty"` // The wave month, 1-based, int

	Deadline time.Time `json:"deadline,omitempty"` // No more responses accepted

	Params []ParamT `json:"params,omitempty"` // I.e. NASDAQ at the time begin of the wave - being used in the wording of the questions
}

// NewSurvey returns a survey based on current time
func NewSurvey(tp string) surveyT {

	if tp == "" {
		panic("survey must have a type")
	}
	s := surveyT{Type: tp}

	t := time.Now()
	if t.Day() > 20 {
		t = t.AddDate(0, 1, 0)
	}
	s.Year = t.Year()
	s.Month = t.Month()

	// see GenerateQuestionnaireTemplates() for time parsing with a locale
	s.Deadline = time.Date(s.Year, s.Month, 28, 23, 59, 59, 0, cfg.Get().Loc).In(cfg.Get().Loc) // This is arbitrary

	s.Params = []ParamT{}

	return s
}

// String is the default identifier
func (s surveyT) String() string {
	return s.Type + "-" + s.WaveID()
}

// WaveID returns the year-month in standard format yyyy-mm
func (s surveyT) WaveID() string {
	// Notice the month +1
	// It is necessary, even though the spec says 'January = 1'
	t := time.Date(s.Year, s.Month+1, 0, 0, 0, 0, 0, cfg.Get().Loc)
	return t.Format("2006-01")
}

// WaveIDPretty is a pretty identifier
func (s surveyT) WaveIDPretty() string {
	if s.Year == 0 {
		return ""
	}
	// Notice the month +1
	// It is necessary, even though the spec says 'January = 1'
	t := time.Date(s.Year, s.Month+1, 0, 0, 0, 0, 0, cfg.Get().Loc)
	// return t.Format("January 2006") // yields English month names.
	return t.Format("2006-01")
}

// TemplateLogoText for display in HTML
func (s surveyT) TemplateLogoText(langCode string) template.HTML {

	ret := ""

	if s.WaveIDPretty() != "" {
		ret = fmt.Sprintf(`
			<span class="survey-org"        >%v </span>
			<span class="survey-name"       >%v </span>
			<span class="survey-wave-id" > - %v</span>
		`,
			s.Org.TrSilent(langCode),
			trl.HyphenizeText(s.Name.TrSilent(langCode)),
			s.WaveIDPretty(),
		)

	} else {
		ret = fmt.Sprintf(`
			<span class="survey-org"        >%v </span>
			<span class="survey-name"       >%v </span>
		`,
			cfg.Get().Mp["app_org"].TrSilent(langCode),
			trl.HyphenizeText(cfg.Get().Mp["app_label"].TrSilent(langCode)),
		)

	}

	return template.HTML(ret)

}

func dropDown(vals []string, selected string) string {

	opts := ""
	for _, val := range vals {
		isSelected := ""
		if val == selected {
			isSelected = "selected"
		}
		opts += fmt.Sprintf("\t\t<option value='%v' %v >%v</option>\n", val, isSelected, val)
	}

	str := `
	<select name="type">
		%v
	</select>`
	str = fmt.Sprintf(str, opts)
	return str
}

// HTMLForm renders an HTML edit form
// for survey data
func (s *surveyT) HTMLForm(questTypes []string, errStr string) string {

	/*
		<datalist id="time-entries">
			<option>01.01.2030 00:00 CEST</option>
			<option>12.04.2021 17:15 CEST</option>
		</datalist>

	*/

	ret := `<!DOCTYPE html>
<html>
<head>
    <meta charset="utf-8" />
    <title>Generate JSON questionnaires</title>
	<style>
		* {font-family: monospace;}
		.survey-edit-form span {
			display: inline-block;
			min-width: 140px;
		}

	</style>
</head>
<body>
		<b>Generate JSON questionnaires</b><br>

		%v

        <form method="POST" class="survey-edit-form"  style='white-space:pre' >
            Type     %v
            Year     <input type="text" name="year"      value="%v"  />
            Month    <input type="text" name="month"     value="%v"  />

                     '01.01.2030 00:00 CEST' for indefinite   
                     '12.04.2021 17:15 CEST' for concrete     
            Deadline <input type="text" name="deadline"  value="%v" placeholder="dd.mm.yyyy hh:mm CEST"   xxlist="time-entries" size=30 /> 

%v
                     <input type="submit" name="submit" id="submit"  value="Submit" accesskey="s"  /> <br>
		</form>
		
        <script> document.getElementById('submit').focus(); </script>
            %v
            %v
            %v    	
            %v    	
</body>
</html>

		`

	if s == nil {
		*s = NewSurvey("fmt")
	}
	dd := dropDown(questTypes, s.Type)
	dd = strings.TrimSpace(dd)

	kv := ""

	nP := 2
	if len(s.Params) < nP {
		for i := len(s.Params); i < nP; i++ {
			s.Params = append(s.Params, ParamT{})
		}
	}
	for i := 0; i < nP; i++ {
		kv += fmt.Sprintf(
			"\t\tParam%v ",
			i,
		)
		kv += fmt.Sprintf(
			"<input type='text' name='param_keys[%v]' placeholder='name%v'  value='%v' />",
			i, i, s.Params[i].Name,
		)
		kv += fmt.Sprintf(
			"<input type='text' name='param_vals[%v]' placeholder='val%v'   value='%v' /><br>",
			i, i, s.Params[i].Val,
		)
	}

	link1 := fmt.Sprintf(
		"<a href='%v/generate-hashes?survey_id=%v&wave_id=%v&p=%v' target='_blank' >Generate hash logins</a> - securing login with a base64 encoded MD5 hash<br>",
		cfg.Pref(), s.Type, s.WaveID(), "1",
	)
	link2 := fmt.Sprintf(
		"<a href='%v/generate-hash-ids?start=10000&stop=10020&host=https://survey2.zew.de' target='_blank' >Generate hash IDs for direct login</a> - ultra short login URL - requires matching DirectLoginRanges in config<br>",
		cfg.Pref(),
	)
	link3 := fmt.Sprintf(
		"<a href='%v/shufflings-to-csv?start=1000&stop=1020' target='_blank' >Shufflings to CSV</a><br>",
		cfg.Pref(),
	)
	link4 := fmt.Sprintf(
		"<a href='%v/create-anonymous-id' target='_blank' >Create anonymous ID</a> - requires anonymous_survey_id and matching direct_login_range<br>",
		cfg.Pref(),
	)

	ret = fmt.Sprintf(
		ret,
		errStr,
		dd,
		s.Year,
		fmt.Sprintf("%02d", int(s.Month)+0), // month
		s.Deadline.In(cfg.Get().Loc).Format("02.01.2006 15:04 MST"),
		kv,
		link1,
		link2,
		link3,
		link4,
	)
	return ret

}
