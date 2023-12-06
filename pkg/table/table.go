package table

import (
	"fmt"

	"github.com/jedib0t/go-pretty/table"
	"github.com/jedib0t/go-pretty/text"
)

type HAlign text.Align

const (
	AlignCenter  HAlign = HAlign(text.AlignCenter)
	AlignRight   HAlign = HAlign(text.AlignRight)
	AlignLeft    HAlign = HAlign(text.AlignLeft)
	AlignJustify HAlign = HAlign(text.AlignJustify)
)

type Row []any

type RenderParam struct {
	Title           string
	TitleAlign      HAlign
	EnableNumbering bool
	Header          []string
	DataSiggle      Row
	DataList        []Row
	DataALign       HAlign
}

func Render(param RenderParam) error {
	if len(param.DataList) > 0 && len(param.DataSiggle) > 0 {
		return fmt.Errorf("err data must be either data singgle or data list. cannot both")
	}

	if len(param.DataSiggle) > 0 && len(param.DataSiggle) != len(param.Header) {
		return fmt.Errorf("header and data must be in the same lenght")
	}

	if len(param.DataList) > 0 {
		for _, val := range param.DataList {
			if len(val) != len(param.Header) {
				return fmt.Errorf("header and data must be in the same lenght")
			}
		}
	}

	t := table.NewWriter()
	t.SetAutoIndex(param.EnableNumbering)
	t.SetTitle(param.Title)
	t.Style().Title.Align = text.Align(param.TitleAlign)

	headers := table.Row{}
	columnConfig := []table.ColumnConfig{}
	for idx, val := range param.Header {
		headers = append(headers, val)
		columnConfig = append(columnConfig, table.ColumnConfig{Number: idx + 1, Align: text.Align(param.DataALign), AlignHeader: text.Align(param.DataALign), VAlign: text.VAlignMiddle})
	}
	t.SetColumnConfigs(columnConfig)
	t.AppendHeader(headers)

	switch {
	case len(param.DataSiggle) > 0:
		var row table.Row
		for _, data := range param.DataSiggle {
			row = append(row, data)
		}
		t.AppendRow(row)

	case len(param.DataList) > 0:
		var rows []table.Row
		for _, val := range param.DataList {
			rows = append(rows, table.Row(val))
		}
		t.AppendRows(rows)
	}

	fmt.Printf("%v\n\n", t.Render())
	return nil
}

// func EntityToRow
