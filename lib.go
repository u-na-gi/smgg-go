package smgg

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type SourceBuilder struct {
	PackagePath string
	FileName    string
	Source      string
}

// ファイルパスが渡されたとする
func CreateSource(sourcePath string) (*SourceBuilder, error) {

	sinfo, err := ParseAstFromFile(sourcePath)
	if err != nil {
		return nil, err
	}
	tp := NewTemplateParser()
	packageName := sinfo.PackageName
	tp.ParseAndUpdateSource(Ptempl, TeplatePackageArg{PackageName: packageName})
	structNames := sinfo.StructNames
	conditions := sinfo.Conditions

	for _, structName := range structNames {
		tp.ParseAndUpdateSource(FuncUpsertStartTmpl, TemplateFuncStructArg{TypeName: structName})
		for fieldName, zeroValue := range conditions[structName] {
			tp.ParseAndUpdateSource(FuncUpsertConditionTmpl, TemplateUpsertConditionStructArg{
				FieldName: fieldName,
				ZeroValue: zeroValue,
			})
		}
		tp.ParseAndUpdateSource(FuncEndTmpl, TemplateFuncStructArg{TypeName: structName})

		tp.ParseAndUpdateSource(FuncUpsertAllowDeleteStartTmpl, TemplateFuncStructArg{TypeName: structName})
		for fieldName, zeroValue := range conditions[structName] {
			tp.ParseAndUpdateSource(FuncUpsertConditionTmpl, TemplateUpsertConditionStructArg{
				FieldName: fieldName,
				ZeroValue: zeroValue,
			})
		}
		tp.ParseAndUpdateSource(FuncEndTmpl, TemplateFuncStructArg{TypeName: structName})

	}

	// 返り値はpackageが存在するパス, ファイル名, 生成したコード
	return &SourceBuilder{
		PackagePath: filepath.Dir(sourcePath),
		FileName:    fmt.Sprintf("%s.generated.go", packageName),
		Source:      tp.GetSource(),
	}, nil
}

func WalkingCurrentProject() ([]string, error) {
	var files []string

	// 現在のディレクトリから再帰的に検索
	err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err // アクセスに失敗した場合は、エラーを返します
		}
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".go") {
			files = append(files, path) // 拡張子が .go のファイルのみを追加
		}
		return nil
	})

	if err != nil {
		fmt.Println("Error walking through directories:", err)
		return nil, err
	}

	return files, nil
}

// SourceBuilderを受け取り、ファイルを生成する
func (sb *SourceBuilder) Generate() error {
	file, err := os.Create(filepath.Join(sb.PackagePath, sb.FileName))
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(sb.Source)
	if err != nil {
		return err
	}
	return nil
}
