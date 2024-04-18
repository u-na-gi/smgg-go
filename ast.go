package smgg

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"

	primitivedefaultvaluego "github.com/u-na-gi/primitive-default-value-go"
	"golang.org/x/tools/go/packages"
)

type StructsInfo struct {
	PackageName string
	StructNames []string
	Conditions  map[string]map[string]*ConditionType
}

type ConditionType struct {
	FieldValue string
	IsPointer  bool
}

func a(fieldType ast.Expr) {
	switch ft := fieldType.(type) {
	case *ast.Ident:
		fmt.Println("--------------111")
		fmt.Println(ft.Name)
		fmt.Println("--------------111")
		switch ft.Name {
		case "string":
			fmt.Println("string")
		case "int", "int8", "int16", "int32", "int64",
			"uint", "uint8", "uint16", "uint32", "uint64", "byte", "rune":
			fmt.Println("int")
		case "float32", "float64":
			fmt.Println("float")
		case "bool":
			fmt.Println("bool")
		case "complex64", "complex128":
			fmt.Println("complex")
		default:
			fmt.Println("default")
		}
	default:
		fmt.Println("default")
	}
}

func ParseAstFromFile(filePath string) (*StructsInfo, error) {
	cfg := &packages.Config{
		Mode:  packages.NeedImports | packages.NeedTypes | packages.NeedTypesInfo,
		Tests: false,
	}

	pkgs, err := packages.Load(cfg, filePath)
	if err != nil {
		log.Fatalf("failed to load packages: %v", err)
	}

	if len(pkgs) == 0 {
		log.Fatalf("no packages found")
	}

	// あとはこっちでとったフィールド名と合体させるといいかなあ
	for _, pkg := range pkgs {
		for _, obj := range pkg.TypesInfo.Defs {

			if obj == nil {
				continue
			}
			fmt.Println("------------------")
			fmt.Printf("obj: %v, type: %#v \n", obj.Name(), obj.Type().Underlying().String())
			fmt.Println("------------------")

		}
	}

	fset := token.NewFileSet()

	// 指定したファイルのソースコードを解析し、ASTを生成
	node, err := parser.ParseFile(fset, filePath, nil, parser.ParseComments)
	if err != nil {
		return nil, err
	}

	packageName := node.Name.Name
	structNames := make([]string, 0)
	conditions := make(map[string]map[string]*ConditionType, 0)

	ast.Inspect(node, func(n ast.Node) bool {
		ts, ok := n.(*ast.TypeSpec)
		if ok {
			tss, ok := ts.Type.(*ast.StructType)
			if ok {
				structNames = append(structNames, ts.Name.Name)
				for _, field := range tss.Fields.List {
					a(field.Type)
					zerv := primitivedefaultvaluego.DefaultValue(field.Type)
					if conditions[ts.Name.Name] == nil {
						conditions[ts.Name.Name] = map[string]*ConditionType{}
					}
					if conditions[ts.Name.Name][field.Names[0].Name] == nil {
						conditions[ts.Name.Name][field.Names[0].Name] = &ConditionType{
							FieldValue: zerv,
							IsPointer:  false,
						}
					} else {
						conditions[ts.Name.Name][field.Names[0].Name].FieldValue = zerv
					}

					// ポインタかどうか確認する
					// 続き
					// ポインタならポインタであるフラグを立てる(true)
					// ポインタでない場合はfalse
					// これがfalseかつ、値がnilの場合、aliasの可能性が高いので、aliasであるかもしれないものとして扱う
					if _, ok := field.Type.(*ast.StarExpr); ok {
						conditions[ts.Name.Name][field.Names[0].Name].IsPointer = true
					}

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
