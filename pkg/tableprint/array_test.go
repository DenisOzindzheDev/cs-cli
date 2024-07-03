package tableprint_test

// import (
// 	"bytes"
// 	"testing"

// 	"github.com/DenisOzindzheDev/cs-cli/pkg/tableprint"
// 	"github.com/stretchr/testify/assert"
// )

// func TestTablePrintA(t *testing.T) {
// 	data := map[string]interface{}{
// 		"Name": "John Doe",
// 		"Age":  30,
// 	}

// 	var buf bytes.Buffer
// 	tableprint.TablePrintA(data, "Key", "Value")

// 	expected := "+------+----------+\n| KEY  | VALUE    |\n+------+----------+\n| Age  | 30       |\n| Name | John Doe |\n+------+----------+\n"
// 	assert.Equal(t, expected, buf.String())
// }
