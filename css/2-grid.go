package css

import (
	"fmt"
	"strings"
)

type GridContainerStyle struct {
	AutoFlow string `json:"auto_flow,omitempty"` // column | row | dense | dense column | dense row

	AutoColumns string `json:"auto_columns,omitempty"` // default beyond TemplateColumns
	AutoRows    string `json:"auto_rows,omitempty"`    // default beyond TemplateRows

	TemplateColumns string `json:"template_columns,omitempty"` // i.e. minmax(4rem, 2fr) [col-title] minmax(4rem, 2fr) minmax(4rem, 2fr)
	TemplateRows    string `json:"template_rows,omitempty"`
	JustifyContent  string `json:"justify_content,omitempty"` // main axis   - all items inside container - space-around | space-between
	JustifyItems    string `json:"justify_items,omitempty"`   // main axis   - item inside its 'cell' - stretch | baseline | center | start | end
	AlignContent    string `json:"align_content,omitempty"`   // second axis - all items inside container - space-around | space-between
	AlignItems      string `json:"align_items,omitempty"`     // second axis - item inside its 'cell' - stretch | baseline | center | start | end

	ColumnGap string `json:"column_gap,omitempty"`
	RowGap    string `json:"row_gap,omitempty"`
}

func gridContainerStyleExample1() GridContainerStyle {
	return GridContainerStyle{
		AutoFlow:        "column",
		AutoColumns:     "minmax(auto,  300px)",
		AutoRows:        "minmax(100px, auto)",
		TemplateColumns: "[col-img] 6fr [col-title] 4fr [col-menu] 1fr",
		TemplateRows:    "repeat(auto-fill, minmax(100px, 1f  ))",
		JustifyContent:  "space-around",
		JustifyItems:    "start",
		AlignContent:    "space-around",
		AlignItems:      "start",
	}
}

func gridContainerStyleExample1Want() string {
	return `	grid-auto-flow: column;
	grid-auto-columns: minmax(auto,  300px);
	grid-auto-rows: minmax(100px, auto);
	grid-template-columns: [col-img] 6fr [col-title] 4fr [col-menu] 1fr;
	grid-template-rows: repeat(auto-fill, minmax(100px, 1f  ));
	justify-content: space-around;
	justify-items: start;
	align-content: space-around;
	align-items: start;
`
}

// CSS renders styles
func (gcs GridContainerStyle) CSS() string {
	s := &strings.Builder{}
	if gcs.AutoFlow != "" {
		fmt.Fprintf(s, "\tgrid-auto-flow: %v;\n", gcs.AutoFlow)
	}
	if gcs.AutoColumns != "" {
		fmt.Fprintf(s, "\tgrid-auto-columns: %v;\n", gcs.AutoColumns)
	}
	if gcs.AutoRows != "" {
		fmt.Fprintf(s, "\tgrid-auto-rows: %v;\n", gcs.AutoRows)
	}
	if gcs.TemplateColumns != "" {
		fmt.Fprintf(s, "\tgrid-template-columns: %v;\n", gcs.TemplateColumns)
	}
	if gcs.TemplateRows != "" {
		fmt.Fprintf(s, "\tgrid-template-rows: %v;\n", gcs.TemplateRows)
	}
	if gcs.JustifyContent != "" {
		fmt.Fprintf(s, "\tjustify-content: %v;\n", gcs.JustifyContent)
	}
	if gcs.JustifyItems != "" {
		fmt.Fprintf(s, "\tjustify-items: %v;\n", gcs.JustifyItems)
	}
	if gcs.AlignContent != "" {
		fmt.Fprintf(s, "\talign-content: %v;\n", gcs.AlignContent)
	}
	if gcs.AlignItems != "" {
		fmt.Fprintf(s, "\talign-items: %v;\n", gcs.AlignItems)
	}
	if gcs.ColumnGap != "" {
		fmt.Fprintf(s, "\tgrid-column-gap: %v;\n", gcs.ColumnGap)
	}
	if gcs.RowGap != "" {
		fmt.Fprintf(s, "\tgrid-row-gap: %v;\n", gcs.RowGap)
	}
	return s.String()
}

type GridItemStyle struct {
	JustifySelf string `json:"justify_self,omitempty"` // main axis   - item inside its 'cell' - stretch | baseline | center | start | end
	AlignSelf   string `json:"align_self,omitempty"`   // second axis - item inside its 'cell' - stretch | baseline | center | start | end
	Col         string `json:"col,omitempty"`
	Row         string `json:"row,omitempty"`
	Order       int    `json:"order,omitempty"`
}

func gridItemStyleExample1() GridItemStyle {
	return GridItemStyle{
		JustifySelf: "start",
		AlignSelf:   "stretch",
		Col:         "col-menu/span 1",
		Row:         "2/-1",
		Order:       12,
	}
}

func gridItemStyleExample1Want() string {
	return `	justify-self: start;
	align-self: stretch;
	grid-column: col-menu/span 1;
	grid-row: 2/-1;
	order: 12;
`
}

// CSS renders styles
func (gis GridItemStyle) CSS() string {
	s := &strings.Builder{}
	if gis.JustifySelf != "" {
		fmt.Fprintf(s, "\tjustify-self: %v;\n", gis.JustifySelf)
	}
	if gis.AlignSelf != "" {
		fmt.Fprintf(s, "\talign-self: %v;\n", gis.AlignSelf)
	}
	if gis.Col != "" {
		fmt.Fprintf(s, "\tgrid-column: %v;\n", gis.Col)
	}
	if gis.Row != "" {
		fmt.Fprintf(s, "\tgrid-row: %v;\n", gis.Row)
	}
	if gis.Order != 0 {
		fmt.Fprintf(s, "\torder: %v;\n", gis.Order)
	}
	return s.String()
}
