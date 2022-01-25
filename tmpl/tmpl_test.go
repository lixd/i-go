package tmpl

import (
	"os"
	"testing"
	"text/template"
)

// 更新时，维护CodeComments中的错误码和注释关系，并修改Start、End即可

var CodeComments = map[int64]string{
	100: "参数错误",
	101: "参数为空",
	102: "伪造参数攻击",
	103: "请求重放攻击",
	104: "异常设备",
	105: "异常IP",
	106: "异常用户",
	107: "TODO",
	108: "TODO",
	109: "TODO",
}

const (
	Start   = 100
	End     = 110
	Package = "vcode"
)

type code struct {
	Number  int64
	Comment string
}

const VCodeTmpl = `
package {{.Package}}
{{/*生成一组const*/}}
const(
	{{range .Data}} C{{.Number}}={{.Number}} // {{.Comment}}
	{{end -}}
)
{{/*定义一个变量来存储MaxID，便于后续在 for range 中使用*/}}
{{- $MaxID := .MaxID}}
{{/*生成对应List列表*/ -}}
var(
	List = []string{
{{- range $i,$v :=.Data -}}
	{{/*判断是最后一个值时，不添加逗号*/}}"{{.Number}}"{{if lt $i $MaxID -}} , {{- end -}}
{{end}}}
)

`

func TestGenerateVCode(t *testing.T) {
	tmpl, err := template.New("test").Parse(VCodeTmpl)
	if err != nil {
		panic(err)
	}

	list := make([]code, 0, End-Start)
	for i := int64(Start); i < End; i++ {
		item := code{
			Number:  i,
			Comment: CodeComments[i],
		}
		list = append(list, item)
	}

	err = tmpl.Execute(os.Stdout, struct {
		Data    []code
		MaxID   int
		Package string
	}{
		Data:    list,
		MaxID:   End - Start - 1, // id从0开始，因此需要减1
		Package: Package,
	})
	if err != nil {
		panic(err)
	}
}

func TestGenerateCode(t *testing.T) {
	tmpl, err := template.New("RespCode").Parse(`
const(
{{range .}} C{{.}}={{.}}
{{end}})
`)
	if err != nil {
		panic(err)
	}

	list := make([]int64, 0, 10)
	for i := int64(1000); i < 1010; i++ {
		list = append(list, i)
	}

	err = tmpl.Execute(os.Stdout, list)
	if err != nil {
		panic(err)
	}
}
