package qst

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"sort"
	"strings"

	"github.com/zew/go-questionnaire/pkg/trl"
)

type optionT struct {
	Key      string
	Val      trl.S //  template.HTML
	Selected bool

	ValLC template.HTML `json:"-"` // helper

}

// DropdownT represents a HTML dropdown control
// Methods need to return a string, so we can use them in templates
type DropdownT struct {
	Name       string
	AutoSubmit bool // onchange this.Form.Submit() is suppressed
	Disabled   bool

	// management of 'style' and 'class' HTML attributes
	Attrs map[template.HTMLAttr]template.HTMLAttr

	LC      string // for rendering the labels inside the template below
	Options []optionT

	NameJavaScriptExpression template.JSStr `json:"-"` // helper
}

// SetName - for usage in templates
func (d *DropdownT) SetName(s string) string {
	d.Name = s
	return ""
}

// SetAutoSubmit - for usage in templates
func (d *DropdownT) SetAutoSubmit(b bool) string {
	d.AutoSubmit = b
	return ""
}

// SetDisabled - for usage in templates
func (d *DropdownT) SetDisabled(b bool) string {
	d.Disabled = b
	return ""
}

//
// Attributes management
//

// HasAttr checks whether an attribute key exists
func (d *DropdownT) HasAttr(k string) bool {
	if _, ok := d.Attrs[template.HTMLAttr(k)]; ok {
		return true
	}
	return false
}

// SetAttr - setting or appending an attribute
func (d *DropdownT) SetAttr(k string, vi interface{}) string {
	sep := " " // separator for CSS class
	if strings.TrimSpace(k) == "style" {
		sep = ";"
	}
	v := strings.TrimSpace(fmt.Sprintf("%v", vi)) // transform integer to string
	if !strings.HasSuffix(v, sep) {
		v += sep
	}
	if d.Attrs == nil {
		d.Attrs = map[template.HTMLAttr]template.HTMLAttr{}
	}
	if vOld, ok := d.Attrs[template.HTMLAttr(k)]; ok && vOld != template.HTMLAttr("") {
		d.Attrs[template.HTMLAttr(k)] += template.HTMLAttr(v) // append
	} else {
		d.Attrs[template.HTMLAttr(k)] = template.HTMLAttr(v) // set anew
	}
	return ""
}

// RemoveAttrVal - removing the value in an attribute
func (d *DropdownT) RemoveAttrVal(k, v string) string {
	if vOld, ok := d.Attrs[template.HTMLAttr(k)]; ok {
		s := string(vOld)
		s = strings.Replace(s, v+";", "", -1)
		s = strings.Replace(s, v+" ", "", -1)
		s = strings.Replace(s, v, "", -1)
		d.Attrs[template.HTMLAttr(k)] = template.HTMLAttr(s)
	}
	return ""
}

// RemoveAllAttrs removes all attributes
func (d *DropdownT) RemoveAllAttrs() string {
	d.Attrs = map[template.HTMLAttr]template.HTMLAttr{}
	return "" // dummy for usage in templates
}

//
// Options management
//

// Add adds an option returns selected key
func (d *DropdownT) Add(k string, v trl.S) string {
	o := optionT{}
	o.Key = k
	o.Val = v
	d.Options = append(d.Options, o)
	return ""
}

// AddPleaseSelect adds a default option at the top
func (d *DropdownT) AddPleaseSelect(v trl.S) {
	leadOpt := []optionT{{Key: "", Val: v}} // i.e. "please choose"
	(*d).Options = append(leadOpt, (*d).Options...)
}

// Selected returns selected key
func (d *DropdownT) Selected() string {
	for i := 0; i < len(d.Options); i++ {
		if d.Options[i].Selected {
			return d.Options[i].Key
		}
	}
	return ""
}

// Select an option by key
func (d *DropdownT) Select(selectKey string) string {
	for i := 0; i < len(d.Options); i++ {
		if d.Options[i].Key == selectKey {
			d.Options[i].Selected = true
		} else {
			d.Options[i].Selected = false
		}
	}
	return ""
}

//
// Sorting stuff
//

func (d *DropdownT) Len() int           { return len(d.Options) }
func (d *DropdownT) Swap(i, j int)      { d.Options[i], d.Options[j] = d.Options[j], d.Options[i] }
func (d *DropdownT) Less(i, j int) bool { return d.Options[i].Val[d.LC] < d.Options[j].Val[d.LC] }

// SortOptionsByLabel for quest generation time
func (d *DropdownT) SortOptionsByLabel() {
	sort.Sort(d)
}

//
// Template stuff
//

var tplStr = `
	<select 
			name='{{ .Name }}'  id='{{ .Name }}'
				
			{{- range $attr, $val := .Attrs}}
				{{$attr}}='{{$val}}'
			{{end -}}

			{{- if .AutoSubmit}}
				onchange='console.log(this.form.{{.NameJavaScriptExpression}}.options[this.form.{{.NameJavaScriptExpression}}.selectedIndex].value); this.form.submit();'
			{{end -}}

			{{- if .Disabled }}
				disabled
			{{end -}}

	>

		{{$outer := .}}
		{{range $Option := .Options -}}
			<!-- keep the ugly formatting of the end if -->
			<option value="{{ $Option.Key }}" {{ if eq $Option.Selected true }}selected{{end}} >{{$Option.ValLC}}</option>
		{{- end}}
	</select>
`

var tpl = template.New("dd")

func init() {
	var err error
	tpl, err = tpl.Parse(tplStr)
	if err != nil {
		panic(err)
	}
}

// Render to io.Writer
func (d *DropdownT) Render(w io.Writer) {
	d.NameJavaScriptExpression = template.JSStr(d.Name)
	for i := 0; i < len(d.Options); i++ {
		d.Options[i].ValLC = template.HTML(d.Options[i].Val[d.LC])
	}
	// log.Printf("%#v", d.Options)
	err := tpl.Execute(w, d)
	if err != nil {
		log.Printf("error rendering dropdown %v: %v", d.Name, err)
	}
}

// RenderStr to string
func (d *DropdownT) RenderStr() string {
	sb := &strings.Builder{}
	d.Render(sb)
	// log.Print(sb.String())
	return sb.String()
}
