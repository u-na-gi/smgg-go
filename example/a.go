package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"text/template"

	primitivedefaultvaluego "github.com/u-na-gi/primitive-default-value-go"
)

// コード生成のためのテンプレート

const ptempl = `package {{.PackageName}}
`
const funcStartTmpl = `
// MergeUpsert{{.TypeName}} merges the source into the target. If an entry exists only in the source, it's added to the target.
// If an entry exists in both, the source's value overwrites the target's.
// Entries present in the target but not in the source are preserved.
// Both target and source must be of the same type.
func MergeUpsert{{.TypeName}}(target {{.TypeName}}, source {{.TypeName}}) {{.TypeName}} {
`

const funcUpsertAllowDeleteStartTmpl = `
// MergeUpsertAllowDelete{{.TypeName}} merges the source object into the target object. If a value from the source object
// does not exist in the target object, it will be added. If the target object already contains
// a value from the source object, it will be overwritten with the value from the source object.
// Both the target and source objects must be of the same type.
func MergeUpsertAllowDelete{{.TypeName}}(target {{.TypeName}}, source {{.TypeName}}) {{.TypeName}} {
`

const funcUpsertConditionTmpl = `
	if source.{{.FieldName}} != {{.ZeroValue}} {
		target.{{.FieldName}} = source.{{.FieldName}}
	}
`

const funcUpsertAllowDeleteTmpl = `
	target.{{.FieldName}} = source.{{.FieldName}}
`

const funcEndTmpl = `
	return target
}
`

func main() {
	filePath := "example/source/source.go"
	// トークンの位置情報を保持するファイルセットを生成
	fset := token.NewFileSet()

	// 指定したファイルのソースコードを解析し、ASTを生成
	node, err := parser.ParseFile(fset, filePath, nil, parser.ParseComments)
	if err != nil {
		fmt.Println("Error parsing file:", err)
		return
	}

	// package名を表示
	fmt.Println("Package:", node.Name.Name)
	packageName := node.Name.Name
	// fmt.Println(packageName)

	structNames := make([]string, 0)
	conditions := make(map[string]map[string]string, 0)

	// ASTをトラバースしてtype定義を探す
	ast.Inspect(node, func(n ast.Node) bool {
		// type定義を見つけた場合
		ts, ok := n.(*ast.TypeSpec)
		if ok {
			// TypeがStructTypeであるかチェック
			tss, ok := ts.Type.(*ast.StructType)
			if ok {
				// structの名前と定義を表示
				fmt.Printf("Struct: %s\n", ts.Name.Name)
				structNames = append(structNames, ts.Name.Name)
				for _, field := range tss.Fields.List {
					fmt.Println(field.Names[0].Name)
					fmt.Println("field.Type", field.Type)
					// typeによってゼロ値を設定
					zerv := primitivedefaultvaluego.DefaultValue(field.Type)
					fmt.Println("zerv", zerv)
					if conditions[ts.Name.Name] == nil {
						conditions[ts.Name.Name] = map[string]string{}
					}
					conditions[ts.Name.Name][field.Names[0].Name] = zerv
				}
				// typeの名前と定義を表示
				fmt.Printf("Type: %s, Definition: %s\n", ts.Name.Name, ts.Type)
			}
		}
		return true // 子ノードも訪れるためにtrueを返す
	})

	fmt.Println("structNames", structNames)
	fmt.Println("conditions", conditions)

	// // テンプレートをパース
	t, err := template.New("code").Parse(ptempl)
	if err != nil {
		fmt.Println("Error parsing template:", err)
		return
	}

	// バッファを作成し、テンプレートを実行してコードを生成
	var buf bytes.Buffer
	if err := t.Execute(&buf, struct {
		PackageName string
	}{
		PackageName: packageName,
	}); err != nil {
		fmt.Println("Error executing template:", err, "packageName", packageName)
		return
	}

	// 生成したコードを出力（またはファイルに書き込む）
	source := buf.String()

	// バッファを作成し、テンプレートを実行してコードを生成
	for _, typeName := range structNames {
		// 関数1生成用のテンプレートをパース
		t, err = template.New("code").Parse(funcStartTmpl)
		if err != nil {
			fmt.Println("Error parsing template:", err)
			return
		}
		fmt.Println("source", source)

		buf.Reset()
		if err := t.Execute(&buf, struct {
			TypeName string
		}{
			TypeName: typeName,
		}); err != nil {
			fmt.Println("Error executing template:", err, "typeName", typeName)
			return
		}

		source += buf.String()
		fmt.Println("source", source)

		for range conditions {
			c, err := template.New("code").Parse(funcUpsertConditionTmpl)
			if err != nil {
				fmt.Println("Error parsing template:", err)
				return
			}

			for fieldName, zeroValue := range conditions[typeName] {
				buf.Reset()
				fmt.Println("fieldName", fieldName)
				fmt.Println("zeroValue", zeroValue)

				if err := c.Execute(&buf, struct {
					FieldName string
					ZeroValue string
				}{
					FieldName: fieldName,
					ZeroValue: zeroValue,
				}); err != nil {
					fmt.Println("Error executing template:", err, "fieldName", fieldName, "zeroValue", zeroValue)
					return
				}

				source += buf.String()
				// fmt.Println("source", source)

			}
		}

		e, err := template.New("code").Parse(funcEndTmpl)
		if err != nil {
			fmt.Println("Error parsing template:", err)
			return
		}

		buf.Reset()
		if err := e.Execute(&buf, struct{}{}); err != nil {
			fmt.Println("Error executing template:", err)
			return
		}

		source += buf.String()

		// 関数2生成用のテンプレートをパース
		t2, err := template.New("code").Parse(funcUpsertAllowDeleteStartTmpl)
		if err != nil {
			fmt.Println("Error parsing template:", err)
			return
		}

		buf.Reset()
		if err := t2.Execute(&buf, struct {
			TypeName string
		}{
			TypeName: typeName,
		}); err != nil {
			fmt.Println("Error executing template:", err, "typeName", typeName)
			return
		}

		source += buf.String()

		for range conditions {
			c, err := template.New("code").Parse(funcUpsertAllowDeleteTmpl)
			if err != nil {
				fmt.Println("Error parsing template:", err)
				return
			}

			for fieldName := range conditions[typeName] {
				buf.Reset()
				if err := c.Execute(&buf, struct {
					FieldName string
				}{
					FieldName: fieldName,
				}); err != nil {
					fmt.Println("Error executing template:", err)
					return
				}
				source += buf.String()
			}

		}

		// end

		buf.Reset()
		if err := e.Execute(&buf, struct{}{}); err != nil {
			fmt.Println("Error executing template:", err)
			return
		}

		source += buf.String()

	}

	fmt.Println("source", source)

	// // 実際のアプリケーションでは、生成したコードをファイルに書き出すことも可能
	err = os.WriteFile("example/source/output.gen.go", []byte(source), 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
}
