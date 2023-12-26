package qst

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/zew/go-questionnaire/pkg/cfg"
	"github.com/zew/go-questionnaire/pkg/cloudio"
	"github.com/zew/go-questionnaire/pkg/trl"
)

// dynamic funcs return a dynamic piece of content
//
// compare CompositeFuncT, validatorT
type dynFuncT func(*QuestionnaireT, *inputT, string) (string, error)

var dynFuncs = map[string]dynFuncT{
	"ResponseStatistics":             ResponseStatistics,
	"DeadlineAndPublication":         DeadlineAndPublication,
	"LinkBack":                       LinkBack,
	"PersonalLink":                   PersonalLink,
	"PermaLink":                      PermaLink,
	"HasEuroQuestion":                ResponseTextHasEuro,
	"FederalStateAboveOrBelowMedian": FederalStateAboveOrBelowMedian,
	"PatLogos":                       PatLogos,
	"RenderStaticContent":            RenderStaticContent,
	"ErrorProxy":                     ErrorProxy,
	"knebSlightlyDistinctLabel":      knebSlightlyDistinctLabel,
	"knebsDownloadURL":               knebDownloadURL,
}

func isOther(inpName string) bool {

	if strings.HasSuffix(inpName, "__other") {
		return true
	}

	if strings.HasSuffix(inpName, "__other_label") {
		return true
	}

	return false
}

var skipInputNames = map[string]map[string]bool{
	"fmt": {
		"selbst":   true,
		"contact":  true,
		"comment":  true,
		"finished": true,

		"rev_free":       true,
		"rev_free_label": true,

		// 2021-11
		"fmr_comment": true,

		// 2023-12 - kneb
		"qz1_download": true,
		"qz2_comment":  true,
	},
}

// Statistics returns the percentage of
// answers responded to.
// It is helper to ResponseStatistics().
func (q *QuestionnaireT) Statistics() (int, int, float64) {

	responses := 0
	counter := 0
	radioDoubles := map[string]int{}

	for pageIdx, p := range q.Pages {

		if !q.IsInNavigation(pageIdx) {
			continue
		}

		for _, gr := range p.Groups {
			for _, i := range gr.Inputs {
				if i.IsLayout() {
					continue
				}
				if i.Type == "hidden" {
					continue
				}

				if isOther(i.Name) {
					continue
				}

				// checkboxes are set to '0' on submit
				//    <input type='hidden' value='0'...

				// textareas are considered mandatory
				// unless configured in skipInputNames[]

				if skipInputNames[q.Survey.Type][i.Name] {
					continue
				}

				if radioDoubles[i.Name] > 0 {
					continue
				}
				radioDoubles[i.Name]++

				counter++

				if i.Response != "" {
					responses++
				}
			}
		}

	}

	pct := 100 * float64(responses) / float64(counter)
	if pct > 100 {
		pct = 100
	}

	return responses, counter, pct
}

// ResponseStatistics returns the percentage of
// answers responded to.
func ResponseStatistics(q *QuestionnaireT, inp *inputT, paramSet string) (string, error) {
	responses, inputs, pct := q.Statistics()
	ret := fmt.Sprintf(cfg.Get().Mp["percentage_answered"].Tr(q.LangCode), responses, inputs, pct)
	// log.Print("ResponseStatistics(): " + ret)
	return ret, nil
}

// DeadlineAndPublication
func DeadlineAndPublication(q *QuestionnaireT, inp *inputT, paramSet string) (string, error) {
	ct := q.Survey.Deadline
	// ct = ct.Truncate(time.Hour)
	cts := ct.Format("02.01.2006 15:04")
	nextDay := q.Survey.Deadline.Add(24 * time.Hour)
	nextDayS := nextDay.Format("02.01.2006")

	ret := fmt.Sprintf(cfg.Get().Mp["survey_ending"].Tr(q.LangCode), cts, nextDayS)
	// log.Print("DeadlineAndPublication: " + ret)
	return ret, nil
}

// LinkBack
func LinkBack(q *QuestionnaireT, inp *inputT, paramSet string) (string, error) {
	ret := cfg.Get().Mp["link_to_previous_page"].Tr(q.LangCode)
	return ret, nil
}

// PersonalLink is only a reminder, does not contain the personal link
func PersonalLink(q *QuestionnaireT, inp *inputT, paramSet string) (string, error) {
	closed := !q.ClosingTime.IsZero()
	ret := ""
	if closed {
		ret = cfg.Get().Mp["finished_by_participant"].Tr(q.LangCode)
		ret = fmt.Sprintf(ret, q.ClosingTime.Format("02.01.2006 15:04"))
	} else {
		ret = cfg.Get().Mp["review_by_personal_link"].Tr(q.LangCode) + cfg.Get().Mp["link_to_previous_page"].Tr(q.LangCode)
	}
	log.Printf("PersonalLink: closed is %v", closed)
	return ret, nil
}

// PermaLink returns the perma link
func PermaLink(q *QuestionnaireT, inp *inputT, paramSet string) (string, error) {
	closed := !q.ClosingTime.IsZero()
	ret := ""
	if closed {
		ret = cfg.Get().Mp["finished_by_participant"].Tr(q.LangCode)
		ret = fmt.Sprintf(ret, q.ClosingTime.Format("02.01.2006 15:04"))
	} else {
		permaLink, ok := q.Attrs["permalink"]
		if ok {
			lnk := cfg.Get().AbsoluteLink() + "/d/" + permaLink
			// log.Printf("lnk: %v", lnk)
			template := cfg.Get().Mp["review_by_permalink"].Tr(q.LangCode)
			ret = fmt.Sprintf(template, lnk, lnk)
		}
	}
	// log.Printf("PermaLink: %v", ret)
	return ret, nil
}

// ResponseTextHasEuro yields texts => want to keep € - want to have €
func ResponseTextHasEuro(q *QuestionnaireT, inp *inputT, paramSet string) (string, error) {

	if q.Attrs == nil {

	}

	attr1, ok1 := q.Attrs["euro-member"]
	attr2, ok2 := q.Attrs["country"] // country of residence - not language - ISO

	if !ok1 || !ok2 {
		return "Question requires known euro-membership and residence code.", nil
	}

	cntry := trl.Countries[attr2]

	cntry["en"] = strings.Replace(cntry["en"], "Czech Republic", "Czechia", -1)
	cntry["de"] = strings.Replace(cntry["de"], "Tschechische Republik", "Tschechien", -1)
	cntry["fr"] = strings.Replace(cntry["fr"], "République tchèque", "Tchéquie", -1)
	cntry["it"] = strings.Replace(cntry["it"], "Repubblica Ceca", "Cechia", -1)

	hl := trl.S{
		"de": "Wirtschaftlicher Nutzen des Euro<br>",
		"en": "Economic benefits of the euro<br>",
		"fr": "Avantages économiques de l'euro<br>",
		"it": "Benefici economici dell'Euro<br>",
	}
	desc := ""
	ret := ""

	if attr1 == "yes" {
		s1 := trl.S{
			"de": fmt.Sprintf("Den Euro in %v als die offizielle Währung zu haben, ist wirtschaftlich vorteilhaft.",
				cntry["de"]),
			"en": fmt.Sprintf("Having the euro in %v as the official currency is economically beneficial.",
				cntry["en"]),
			"fr": fmt.Sprintf("Avoir l'euro en %v comme monnaie officielle est économiquement avantageux.",
				cntry["fr"]),
			"it": fmt.Sprintf("Avere l'Euro come valuta ufficiale in %v è economicamente vantaggioso.",
				cntry["it"]),
		}
		desc = s1[q.LangCode]

	} else {
		s1 := trl.S{
			"de": fmt.Sprintf("Den Euro in %v als offizielle Währung einzuführen, wäre wirtschaftlich vorteilhaft. ",
				cntry["de"]),
			"en": fmt.Sprintf("Introducing the euro in %v as the official currency would be economically beneficial.",
				cntry["en"]),
			"fr": fmt.Sprintf("L'introduction de l'euro dans %v en tant que monnaie officielle serait économiquement avantageuse.",
				cntry["fr"]),
			"it": fmt.Sprintf("Introdurre l'Euro come valuta ufficiale in %v sarebbe economicamente vantaggioso.",
				cntry["it"]),
		}
		desc = s1[q.LangCode]
	}

	ret = fmt.Sprintf("<b> %v </b> %v", hl[q.LangCode], desc)

	return ret, nil

}

// FederalStateAboveOrBelowMedian returns "besser" or "schlechter";
// depending on the user's federal state education ranking
func FederalStateAboveOrBelowMedian(q *QuestionnaireT, inp *inputT, paramSet string) (string, error) {

	attr1, ok := q.Attrs["aboveOrBelowMedian"]

	if !ok {
		return "Question requires known euro-membership and residence code.", nil
	}
	return attr1, nil

}

// PatLogos - only for the img src URLs
func PatLogos(q *QuestionnaireT, inp *inputT, paramSet string) (string, error) {

	return fmt.Sprintf(
		`
		<div class="uni-logos  logo-imgs-in-content">
			<img src="%v"  style="width:61%%;"  alt=""  >
			<img src="%v"  style="width:33%%;"  alt=""  >
			<img src="%v"  style="width:50%%;"  alt=""  >
			<img src="%v"  style="width:44%%;"  alt=""  >
			<img src="%v"  style="width:28%%;"  alt=""  >
		</div>
		
		<br>
		
		`,
		cfg.Pref("/img/pat/uni-mannheim-wide.png"),
		cfg.Pref("/img/pat/uni-koeln.png"),
		cfg.Pref("/img/pat/uni-muenster.png"),
		cfg.Pref("/img/pat/uni-zurich.png"),
		cfg.Pref("/img/pat/zew.png"),
	), nil

}

// this is a copy of tpl.packageDocPrefix
// maybe we should move it into the config
var packageDocPrefix = "/doc/"

// RenderStaticContent - http request time display of a markdown file
func RenderStaticContent(q *QuestionnaireT, inp *inputT, paramSet string) (string, error) {

	w1 := &strings.Builder{}
	// log.Printf("RenderStaticContent: %v", paramSet)
	err := cloudio.RenderStaticContent(w1, paramSet, q.Survey.Type, q.LangCode, packageDocPrefix)
	if err != nil {
		log.Print(err)
	}

	return w1.String(), err

}

// ErrorProxy - shows errors for inputs named like paramSet
func ErrorProxy(q *QuestionnaireT, inp *inputT, paramSet string) (string, error) {
	return "", nil
}

// knebSlightlyDistinctLabel yields distinct labels depending on treatment
func knebSlightlyDistinctLabel(q *QuestionnaireT, inp *inputT, paramSet string) (string, error) {

	if q.UserIDInt()%2 == 0 {
		return `Wie motiviert sind Sie, sich        mit dem Thema „Sparen und investieren“ zu befassen?`, nil
	} else {
		return `Wie motiviert sind Sie, sich weiter mit dem Thema „Sparen und investieren“ zu befassen?`, nil
	}

}

// Shows different URL, depending on previous answer
func knebDownloadURL(q *QuestionnaireT, inp *inputT, paramSet string) (string, error) {

	urls := map[string]string{
		"weber":              "https://www.arero.de/fileadmin/user_upload/07_downloads/genial_einfach_investieren_ebook.pdf",
		"zewexpertise":       "https://ftp.zew.de/pub/zew-docs/ZEWKurzexpertisen/ZEW_Kurzexpertise2303.pdf",
		"verbraucherzentral": "https://www.verbraucherzentrale.de/alles-zur-geldanlage-das-muessen-sie-dazu-wissen-18777",
	}

	inpSrc := q.ByName("qz1_download")
	if inpSrc == nil {
		return ``, nil
		// return `Ihr Download wurde nicht gefunden.`, nil
	}
	if inpSrc.Response == "" || inpSrc.Response == "0" || inpSrc.Response == "nothanks" {
		return ``, nil
		// return `no thank you - or no selection at all.`, nil
	}

	return fmt.Sprintf(`<a target='blank' href=%v>Ihr Dankeschön-Download<a>`, urls[inpSrc.Response]), nil

}
