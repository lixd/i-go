package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"strings"
)

// func main() {
// 	expr, _ := parser.ParseExpr("a * -1")
// 	fmt.Printf("%#v\n", expr)
// }

var codeTemplate = `
func getReqStruct(r *http.Request) (*{{requestStructName}}, error) {
	r.ParseForm()
	var reqStruct = &{{requestStructName}}{}

	// bind data
	{{bindData}}j

	// bind partial example
	// reqStruct.{{fieldName}} =
	// {{transToFieldType}}(r.Form['{{fieldTagFormName}}'])


	if bindErr != nil {
		return nil, err
	}

	// validate data
	{{validateData}}

	// validate partial example
	// validateErr = validate(reqStruct.{{fieldName}}, validateStr)
	// if validateErr != nil
	// return nil, err


	return reqStruct, nil
}
`

func getTag(input string) []structTag {
	var out []structTag
	var tagStr = input
	tagStr = strings.Replace(tagStr, "`", "", -1)
	tagStr = strings.Replace(tagStr, "\"", "", -1)
	tagList := strings.Split(tagStr, " ")
	for _, val := range tagList {
		tmpArr := strings.Split(val, ":")
		st := structTag{}
		st.key = tmpArr[0]
		st.values = strings.Split(tmpArr[1], ",")
		out = append(out, st)
	}

	return out
}

type structTag struct {
	key    string
	values []string
}

func main() {
	fset := token.NewFileSet()
	// if the src parameter is nil, then will auto read the second filepath file
	f, err := parser.ParseFile(fset, "./example.go", nil, parser.Mode(0))
	if err != nil {
		panic(err)
	}
	//	ast.Print(fset, f.Decls[0])

	tagList := getTag(f.Decls[0].(*ast.GenDecl).Specs[0].(*ast.TypeSpec).Type.(*ast.StructType).Fields.List[0].Tag.Value)
	fieldName := f.Decls[0].(*ast.GenDecl).Specs[0].(*ast.TypeSpec).Type.(*ast.StructType).Fields.List[0].Names[0].Name
	fieldType := f.Decls[0].(*ast.GenDecl).Specs[0].(*ast.TypeSpec).Type.(*ast.StructType).Fields.List[0].Type.(*ast.Ident).Name
	requestStructName := f.Decls[0].(*ast.GenDecl).Specs[0].(*ast.TypeSpec).Name.Name
	fmt.Println(tagList)           // [{form [star]} {validation [gte=1 lte=5]} {doc [formData]}]
	fmt.Println(fieldName)         // Star
	fmt.Println(fieldType)         // int
	fmt.Println(requestStructName) // CollectRequest
}
