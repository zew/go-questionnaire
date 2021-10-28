package qst

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/russross/blackfriday/v2"
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
	"PersonalLink":                   PersonalLink,
	"HasEuroQuestion":                ResponseTextHasEuro,
	"FederalStateAboveOrBelowMedian": FederalStateAboveOrBelowMedian,
	"PatLogos":                       PatLogos,
	"RenderStaticContent":            RenderStaticContent,
	"ErrorProxy":                     ErrorProxy,
}

var skipInputNames = map[string]map[string]bool{
	"fmt": {
		"selbst":   true,
		"contact":  true,
		"comment":  true,
		"finished": true,

		"rev_free":       true,
		"rev_free_label": true,

		"inff_2021__other_label": true,
		"inff_2022__other_label": true,
		"inff_2023__other_label": true,
	},
}

// Statistics returns the percentage of
// answers responded to.
// It is helper to ResponseStatistics().
func (q *QuestionnaireT) Statistics() (int, int, float64) {
	responses := 0
	counter := 0
	radioDoubles := map[string]int{}
	for _, p := range q.Pages {
		for _, gr := range p.Groups {
			for _, i := range gr.Inputs {
				if i.IsLayout() {
					continue
				}
				if i.Type == "hidden" {
					continue
				}
				if i.Type == "checkbox" { // I am not sure
					continue
				}
				// if i.Type == "textarea" { // textarea is always optional
				// 	continue
				// }
				if skipInputNames[q.Survey.Type][i.Name] {
					continue
				}
				if radioDoubles[i.Name] > 0 {
					continue
				}
				radioDoubles[i.Name]++
				counter++
				if i.Response != "" && i.Response != "0" {
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
	ct := q.Survey.Deadline
	// ct = ct.Truncate(time.Hour)
	cts := ct.Format("02.01.2006 15:04")
	nextDay := q.Survey.Deadline.Add(24 * time.Hour)
	nextDayS := nextDay.Format("02.01.2006")

	s1 := fmt.Sprintf(cfg.Get().Mp["percentage_answered"].Tr(q.LangCode), responses, inputs, pct)
	s2 := fmt.Sprintf(cfg.Get().Mp["survey_ending"].Tr(q.LangCode), cts, nextDayS)
	ret := s1 + s2
	// log.Print("ResponseStatistics: " + ret)
	return ret, nil
}

// PersonalLink returns the entry link
func PersonalLink(q *QuestionnaireT, inp *inputT, paramSet string) (string, error) {
	closed := !q.ClosingTime.IsZero()
	ret := ""
	if closed {
		ret = cfg.Get().Mp["finished_by_participant"].Tr(q.LangCode)
		ret = fmt.Sprintf(ret, q.ClosingTime.Format("02.01.2006 15:04"))
	} else {
		ret = cfg.Get().Mp["review_by_personal_link"].Tr(q.LangCode)
	}
	log.Printf("PersonalLink: closed is %v", closed)
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

// RenderStaticContent - http request time display of a markdown file
func RenderStaticContent(q *QuestionnaireT, inp *inputT, paramSet string) (string, error) {

	w1 := &strings.Builder{}
	err := RenderStaticContentInner(
		w1, paramSet, q.Survey.Type, q.LangCode,
	)
	if err != nil {
		log.Print(err)
	}

	return w1.String(), err

}

type staticPrefixT string                     // significant url path fragment
var packageDocPrefix = staticPrefixT("/doc/") // application singleton

// RenderStaticContentInner is a damn copy of tpl.RenderStaticContent
func RenderStaticContentInner(w io.Writer, subPth, site, lang string) error {

	var (
		bts []byte
		err error
	)

	// special file path: README.md is read directly from the app root via classic ioutil
	if strings.HasSuffix(subPth, "README.md") {
		bts, err = ioutil.ReadFile("./README.md")
		if err != nil {
			s := fmt.Sprintf("MarkdownH: cannot open README.md in app root: %v", err)
			log.Printf(s)
			return errors.Wrap(err, s)
		}
		// rewrite links in README.MD from app root
		//    ./app-bucket/content/somedir/my-img.png
		// to
		//          /urlprefix/doc/somedir/my-img.png
		//                    /doc/somedir/my-img.png  (without prefix)
		{
			needle := []byte("./app-bucket/content/")
			subst := []byte(cfg.PrefTS(string(packageDocPrefix)))
			bts = bytes.Replace(bts, needle, subst, -1)
		}

	} else {
		pth := path.Join(".", "content", site, lang, subPth) // not filepath; cloudio always has forward slash
		bts, err = cloudio.ReadFile(pth)
		if err != nil {
			if errors.Is(err, os.ErrNotExist) {
				bts, err = cloudio.ReadFile(path.Join(".", "content", site, subPth))
				if errors.Is(err, os.ErrNotExist) {
					bts, err = cloudio.ReadFile(path.Join(".", "content", subPth))
				}
			}
		}
		if err != nil {
			s := fmt.Sprintf("MarkdownH: cannot open markdown %v or upwards: %v", pth, err)
			log.Printf(s)
			return errors.Wrap(err, s)
		}

		{
			// static and dynamic link back
			needle1 := []byte("(./../../../../../../README.md")
			needle2 := []byte("(./../../../../../README.md")
			needle3 := []byte("(./../../../../README.md")
			subst := []byte("(" + cfg.Pref("/doc/README.md"))
			bts = bytes.Replace(bts, needle1, subst, -1)
			bts = bytes.Replace(bts, needle2, subst, -1)
			bts = bytes.Replace(bts, needle3, subst, -1)
		}

		{
			// relative links between static files dont work, if browser url has no trailing slash;
			// rewrite
			//                   ./linux-instructions.md
			// to
			//     ./urlprefix/doc/linux-instructions.md
			needle := []byte("(./")
			subst := []byte("(" + cfg.PrefTS("/doc/"))
			bts = bytes.Replace(bts, needle, subst, -1)
		}

	}

	// rewrite Links from static content to back application:
	//     {{AppPrefix}}
	// to
	//     /urlprefix/
	bts = bytes.Replace(bts, []byte("/{{AppPrefix}}"), []byte(cfg.Pref()), -1)

	// Since blackfriday version 1.52, bullet lists only work for UNIX line breaks
	// bts = bytes.ReplaceAll(bts, []byte("\r\n"), []byte("\n"))

	fmt.Fprint(w, "\n\t<div class='markdown'>\n")

	ext := path.Ext(subPth)
	w1 := &strings.Builder{}
	if ext == ".html" {
		fmt.Fprint(w1, string(bts)) // html direct
	} else {
		fmt.Fprint(w1, string(blackfriday.Run(bts))) // render markdown
	}

	hp := trl.HyphenizeText(w1.String())

	fmt.Fprint(w, hp)
	fmt.Fprint(w, "\n\t</div>  <!-- markdown -->\n")

	// output += "<br>\n<br>\n<br>\n<p style='font-size: 75%;'>\nRendered by russross/blackfriday</p>\n" // Inconspicuous rendering marker

	return nil

}

// ErrorProxy - shows errors for inputs named like paramSet
//
func ErrorProxy(q *QuestionnaireT, inp *inputT, paramSet string) (string, error) {
	return "", nil
}
