package main

import (
	"fmt"
	"os"
	"sort"
	"text/template"
)

var tepl1 = `func (m *{{.ModelName}}) HMapKey() string {
	return fmt.Sprintf("{{.TableName}}:{{.EntityDBID}}:%v", m.{{.EntityID}})
}`

var tepl2 = `
func (m *{{.ModelName}}) PK() (pk string){
	fmtStr:="{{.TableName}}
	{{- range $i, $v := .PKFields -}}
		:%v
	{{- end -}}
	"
	return fmt.Sprintf(fmtStr
		{{- range $i, $v := .PKFields -}}
			, m.{{$v}}
		{{- end -}}
		)
}
`

func main() {

	{
		data := map[string]interface{}{
			"ModelName":  "A",
			"TableName":  "t1",
			"EntityDBID": "id",
			"EntityID":   "ID",
		}
		tmpl, _ := template.New("test").Parse(tepl1)
		tmpl.Execute(os.Stdout, data)
		fmt.Println("")
	}

	{
		data := map[string]interface{}{
			"ModelName": "A",
			"TableName": "t1",
			"PKFields":  sort.StringSlice{"ID", "SubID"},
		}
		tmpl, _ := template.New("test").Parse(tepl2)
		tmpl.Execute(os.Stdout, data)
		fmt.Println("")
	}
}
