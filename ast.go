package structmergingcodegeneratorgo

import (
	"go/ast"
	"go/parser"
	"go/token"

	primitivedefaultvaluego "github.com/u-na-gi/primitive-default-value-go"
)

type StructsInfo struct {
	PackageName string
	StructNames []string
	Conditions  map[string]map[string]string
}

func ParseAstFromFile(filePath string) (*StructsInfo, error) {
	fset := token.NewFileSet()

	// 指定したファイルのソースコードを解析し、ASTを生成
	node, err := parser.ParseFile(fset, filePath, nil, parser.ParseComments)
	if err != nil {
		return nil, err
	}

	packageName := node.Name.Name
	structNames := make([]string, 0)
	conditions := make(map[string]map[string]string, 0)

	ast.Inspect(node, func(n ast.Node) bool {
		ts, ok := n.(*ast.TypeSpec)
		if ok {
			tss, ok := ts.Type.(*ast.StructType)
			if ok {
				structNames = append(structNames, ts.Name.Name)
				for _, field := range tss.Fields.List {
					zerv := primitivedefaultvaluego.DefaultValue(field.Type)
					if conditions[ts.Name.Name] == nil {
						conditions[ts.Name.Name] = map[string]string{}
					}
					conditions[ts.Name.Name][field.Names[0].Name] = zerv
				}
			}
		}
		return true
	})

	return &StructsInfo{
		PackageName: packageName,
		StructNames: structNames,
		Conditions:  conditions,
	}, nil
}
