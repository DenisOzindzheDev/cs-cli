package tableprint

import (
	"io"

	"github.com/olekukonko/tablewriter"
)

func TablePrintStringMap(w io.Writer, data map[string]string, columns ...string) {
	t := tablewriter.NewWriter(w)
	t.SetHeader(columns)
	for k, v := range data {
		t.Append([]string{k, v})
	}
	t.Render()
}
