/*
Package struct2form creates an HTML input form for a given struct type;
field info is taken from the json struct tag;
additional attributes are taken from the form struct tag;
options for input[select] need to be provided by the caller;

ParseMultipartForm() and ExtractUploadedFile() are helper funcs
to extract file upload data.

Example

type loadAssetFromExcel struct {
	DateLayout  string `json:"date_layout,omitempty"  form:"accesskey='t',maxlength='16',size='16',pattern='[0-9\\.\\-/]{10}',placeholder='2006/01/02 15:04'"` // 2006-01-02 15:04
	Verbose     bool   `json:"verbose,omitempty"      form:"accesskey='v'"`
	DateColumn  int    `json:"date_column,omitempty"  form:"accesskey='a',min=0,max='100',maxlength='2',size='2'"`
	ValueColumn int    `json:"value_column,omitempty" form:"accesskey='l',min=0,max='100',maxlength='2',size='2'"`
	Upload      []byte `json:"upload,omitempty"       form:"accesskey='u',accept='.xlsx'"`
}


The package does not provide parsing request forms into a struct type.
For this, we recommend
	"github.com/go-playground/form"
since it accepts json tags despite containing  ,omitempty.
Is also tolerates superfluous request fields -
thus the submit button does not need a struct field.

Example
	dec := form.NewDecoder()
	dec.SetTagName("json") // recognizes and ignores ,omitempty
	err = dec.Decode(&frm, r.Form)
	...
	fmt.Fprint(w, struct2form.HTML(frm))


Accesskeys are not inside the label, but inside the input tag.


TODO:
Floats with various precision

Optional hiding of spinners for integers/floats

Support for date and time

*/
package struct2form

import (
	"bytes"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
	"strings"
	"unicode"
)

// S2FT contains formatting options for converting a struct into a HTML form
type S2FT struct {
	Indent       int
	ShowSubmit   bool
	ShowHeadline bool
}

// S2F is the default formatter
var S2F = &S2FT{
	Indent:       80,
	ShowSubmit:   false,
	ShowHeadline: false,
}

type option struct {
	Key, Val string
}

type options []option

// creating <option val='...'>...</option> tags
func (opts options) HTML(selected string) string {
	w := &bytes.Buffer{}
	for _, o := range opts {
		if o.Key == selected {
			fmt.Fprintf(w, "\t<option value='%v' selected >%v</option>\n", o.Key, o.Val)
		} else {
			fmt.Fprintf(w, "\t<option value='%v'          >%v</option>\n", o.Key, o.Val)
		}
	}
	return w.String()
}

// AddOptions is used by the caller to prepare multiple option key-values
// to be passed into HTML() func
func AddOptions(previous map[string]options, name string, keys, values []string) map[string]options {
	if previous == nil {
		previous = map[string]options{}
	}
	for i, key := range keys {
		previous[name] = append(previous[name], option{key, values[i]})
	}
	return previous
}

// golang type to html input type
func toInputType(t string) string {
	switch t {
	case "bool":
		return "checkbox"
	case "string":
		return "text"
	case "int", "float64":
		return "number"
	case "[]uint8":
		return "file"
	}
	return "text"
}

// retrieving a special 'form' tag from the struct
func structTag(tags, key string) string {
	tagss := strings.Split(tags, ",")
	for _, a := range tagss {
		aLow := strings.ToLower(a)
		if strings.HasPrefix(aLow, key) {
			kv := strings.Split(a, "=")
			if len(kv) == 2 {
				return strings.Trim(kv[1], "'")
			}
		}
	}
	return ""
}

// convert all 'form' struct tags to  html input attributes
func structTagsToAttrs(tags string) string {
	tagss := strings.Split(tags, ",")
	ret := ""
	for _, t := range tagss {
		t = strings.TrimSpace(t)
		tl := strings.ToLower(t) // tag lower
		switch {
		case strings.HasPrefix(tl, "accesskey"): // goes into input, not into label
			ret += " " + t
		case strings.HasPrefix(tl, "size="): // visible width of input field
			ret += " " + t
		case strings.HasPrefix(tl, "maxlength="): // digits of input data
			ret += " " + t
		case strings.HasPrefix(tl, "max="): // for input number
			ret += " " + t
		case strings.HasPrefix(tl, "min="): // for input number
			ret += " " + t
		case strings.HasPrefix(tl, "pattern="): // client side validation; i.e. date layout [0-9\\.\\-/]{10}
			ret += " " + t
		case strings.HasPrefix(tl, "placeholder="): // a watermark showing expected input; i.e. 2006/01/02 15:04
			ret += " " + t
		case strings.HasPrefix(tl, "options="): // for select inputs - string key for the provided list of options
			ret += " " + t
		case strings.HasPrefix(tl, "accept="): // file upload extension
			ret += " " + t
		case strings.HasPrefix(tl, "onchange"): // file upload extension
			ret += " " + "onchange='javascript:this.form.submit();'"
		}

	}
	return ret
}

// for example 'Date layout' with accesskey 't' becomes 'Da<u>t</u>e layout'
func accessKeyify(s, attrs string) string {
	ak := structTag(attrs, "accesskey")
	if ak == "" {
		return s
	}
	akr := rune(ak[0])
	akrUp := unicode.ToUpper(akr)

	s2 := []rune{}
	found := false
	// log.Printf("-%s- -%s-", s, ak)
	for _, ru := range s {
		// log.Printf("\tcomparing %#U to %#U - %#U", ru, akr, akrUp)
		if (ru == akr || ru == akrUp) && !found {
			s2 = append(s2, '<', 'u', '>')
			s2 = append(s2, ru)
			s2 = append(s2, '<', '/', 'u', '>')
			found = true
			continue
		}
		s2 = append(s2, ru)
	}
	return string(s2)
}

// labelize converts struct field names and json field names
// to human readable format:
// bond_fund => Bond fund
// bondFund  => Bond fund
// bondFUND  => Bond fund
//
// notice rare edge case: BONDFund would be converted to 'BONDF und'
func labelize(s string) string {
	rs := make([]rune, 0, len(s))
	previousUpper := false
	for i, char := range s {
		if i == 0 {
			rs = append(rs, unicode.ToUpper(char))
			previousUpper = true
		} else {
			if char == '_' {
				char = ' '
			}
			if unicode.ToUpper(char) == char {
				if !previousUpper {
					rs = append(rs, ' ')
					rs = append(rs, unicode.ToLower(char))
				} else {
					rs = append(rs, unicode.ToLower(char))
				}
				previousUpper = true
			} else {
				rs = append(rs, char)
				previousUpper = false
			}
		}
	}
	return string(rs)
}

// ParseMultipartForm parses an HTTP request form
// with file attachments
func ParseMultipartForm(r *http.Request) error {

	if r.Method == "GET" {
		return nil
	}

	const _24K = (1 << 20) * 24
	err := r.ParseMultipartForm(_24K)
	if err != nil {
		log.Printf("Parse multipart form error: %v\n", err)
		return err
	}
	return nil
}

// ExtractUploadedFile extracts a file from an HTTP POST request.
// It needs the request form to be prepared with ParseMultipartForm.
func ExtractUploadedFile(r *http.Request) (bts []byte, fname string, err error) {

	if r.Method == "GET" {
		return
	}

	_, fheader, err := r.FormFile("upload")
	if err != nil {
		log.Printf("Error unpacking upload bytes from post request: %v\n", err)
		return
	}

	fname = fheader.Filename
	log.Printf("Uploaded filename = %+v", fname)

	rdr, err := fheader.Open()
	if err != nil {
		log.Printf("Error opening uploaded file: %v\n", err)
		return
	}
	defer rdr.Close()

	bts, err = ioutil.ReadAll(rdr)
	if err != nil {
		log.Printf("Error reading uploaded file: %v\n", err)
		return
	}

	log.Printf("Extracted %v bytes from uploaded file", len(bts))
	return

}

// HTML takes a struct instance
// and turns it into an HTML form.
func (sft *S2FT) HTML(intf interface{}, optSelectOptions ...map[string]options) template.HTML {

	needSubmit := false // only select with onchange:submit() ?

	selectOptions := map[string]options{}
	if len(optSelectOptions) > 0 {
		selectOptions = optSelectOptions[0]
	}

	ifVal := reflect.ValueOf(intf)
	// ifVal = ifVal.Elem() // de reference
	if ifVal.Kind().String() != "struct" {
		return template.HTML(fmt.Sprintf("struct2form.HTML() - first arg must be struct - is %v", ifVal.Kind()))
	}

	w := &bytes.Buffer{}

	if sft.ShowHeadline {
		fmt.Fprintf(w, "<h3 style='margin-left:%vpx'>%v</h3>\n", 2, labelize(ifVal.Type().Name()))
	}

	//
	uploadPostForm := false
	for i := 0; i < ifVal.NumField(); i++ {
		tp := ifVal.Field(i).Type().Name() // primitive type name: string, int
		if ifVal.Type().Field(i).Type.Kind() == reflect.Slice {
			tp = "[]" + ifVal.Type().Field(i).Type.Elem().Name()
		}
		if toInputType(tp) == "file" {
			uploadPostForm = true
			break
		}
	}

	if uploadPostForm {
		fmt.Fprint(w, "<form class='control-form'  method='post' enctype='multipart/form-data'>\n")
	} else {
		fmt.Fprint(w, "<form class='control-form'>\n")
	}

	// Render fields
	for i := 0; i < ifVal.NumField(); i++ {

		fldName := ifVal.Type().Field(i).Name // i.e. Name, Birthdate

		if strings.HasPrefix(fldName, "separator") && len(fldName) == len("separator")+2 {
			fmt.Fprintf(w, "<hr>")
			continue
		}

		if fldName[0:1] != strings.ToUpper(fldName[0:1]) {
			// skip unexported
			continue
		}

		inpName := ifVal.Type().Field(i).Tag.Get("json") // i.e. date_layout
		inpName = strings.Replace(inpName, ",omitempty", "", -1)
		frmLabel := labelize(inpName)

		attrs := ifVal.Type().Field(i).Tag.Get("form")

		if attrs == "-" {
			continue
		}

		val := ifVal.Field(i)
		tp := ifVal.Field(i).Type().Name() // primitive type name: string, int
		if ifVal.Type().Field(i).Type.Kind() == reflect.Slice {
			tp = "[]" + ifVal.Type().Field(i).Type.Elem().Name() // []byte => []uint8
		}

		// Label
		fmt.Fprintf(w,
			"<label for='%s' style='display:inline-block; min-width: %vpx; text-align: right;' >%v</label>\n",
			inpName, sft.Indent, accessKeyify(frmLabel, attrs),
		)

		// Input
		if strings.Contains(attrs, "options") {
			fmt.Fprintf(w, "<select name='%v' id='%v' %v />\n", inpName, inpName, structTagsToAttrs(attrs))
			fmt.Fprint(w, selectOptions[inpName].HTML(val.String()))
			fmt.Fprint(w, "</select><br>\n")
			continue
		}
		if toInputType(tp) == "checkbox" {
			needSubmit = true
			checked := ""
			if val.Bool() {
				checked = "checked"

			}
			fmt.Fprintf(w, "<input type='%v' name='%v' id='%v' value='%v' %v %v />\n", toInputType(tp), inpName, inpName, "true", checked, structTagsToAttrs(attrs))
			fmt.Fprintf(w, "<input type='hidden' name='%v' value='false' />\n", inpName)
			sfx := structTag(attrs, "suffix")
			if sfx != "" {
				fmt.Fprintf(w, "<span style='font-size:85%%;'>&nbsp;%s</span>", sfx)
			}
			fmt.Fprintf(w, "<br>\n")
			continue
		}
		if toInputType(tp) == "file" {
			needSubmit = true
			// <input type="file"   name="upload" id="upload" value="ignored.json" accept=".json" >
			fmt.Fprintf(w, "<input type='%v' name='%v' id='%v' value='%v' %v /><br>\n", toInputType(tp), inpName, inpName, "ignored.json", structTagsToAttrs(attrs))
			continue
		}

		needSubmit = true
		fmt.Fprintf(w, "<input type='%v' name='%v' id='%v' value='%v' %v />", toInputType(tp), inpName, inpName, val, structTagsToAttrs(attrs))
		sfx := structTag(attrs, "suffix")
		if sfx != "" {
			fmt.Fprintf(w, "<span style='font-size:85%%;'>&nbsp;%s</span>", sfx)
		}
		fmt.Fprintf(w, "<br>\n")

	}

	if needSubmit || sft.ShowSubmit {
		// Name should *not* be 'submit' to avoid error on this.form.submit() of 'submit is not a function' stackoverflow.com/questions/833032/
		fmt.Fprintf(w, "<button  type='submit' name='btnSubmit' value='1' accesskey='s' style='padding: 4px 16px;border-radius: 4px;margin-left: %vpx;' ><b>S</b>ubmit</button><br>\n", sft.Indent+8)
	} else {
		fmt.Fprintf(w, "<input   type='hidden' name='btnSubmit' value='1'\n")

	}

	fmt.Fprint(w, "</form>\n")
	fmt.Fprint(w, "<br>\n")

	return template.HTML(w.String())
}

// HTML takes a struct instance
// and uses the default formatter
// to turns it into an HTML form.
func HTML(intf interface{}, optSelectOptions ...map[string]options) template.HTML {
	return S2F.HTML(intf, optSelectOptions...)
}
