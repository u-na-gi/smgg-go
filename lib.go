package smgg

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type SourceBuilder struct {
	PackageName string
	SourcePaths []string
	PackagePath string
	FileName    string
	Source      GenSourcer
}

// package: [ソースパス]にする必要あり

func AggregateByPackageName(sourcePaths []string) (map[string]*SourceBuilder, error) {
	packageMap := make(map[string]*SourceBuilder)
	for _, sourcePath := range sourcePaths {
		sinfo, err := ParseAstFromFile(sourcePath)
		if err != nil {
			return nil, err
		}
		packageName := sinfo.PackageName
		if _, ok := packageMap[packageName]; !ok {
			tp := NewTemplateParser()
			tp.ParseAndUpdateSource(Ptempl, TeplatePackageArg{PackageName: packageName})
			packageMap[packageName] = &SourceBuilder{
				PackageName: packageName,
				PackagePath: filepath.Dir(sourcePath),
				FileName:    fmt.Sprintf("%s.smgg.gen.go", packageName),
				SourcePaths: []string{},
				Source:      tp,
			}
		}
		packageMap[packageName].SourcePaths = append(packageMap[packageName].SourcePaths, sourcePath)
	}

	return packageMap, nil
}

// ファイルパスが渡されたとする
func (tp *genSourcer) CreateSource(sourcePath string) (GenSourcer, error) {

	sinfo, err := ParseAstFromFile(sourcePath)
	if err != nil {
		return nil, err
	}
	structNames := sinfo.StructNames
	conditions := sinfo.Conditions

	for _, structName := range structNames {
		tp.ParseAndUpdateSource(FuncUpsertStartTmpl, TemplateFuncStructArg{TypeName: structName})
		for fieldName, zeroValue := range conditions[structName] {
			tp.ParseAndUpdateSource(FuncUpsertConditionTmpl, TemplateUpsertConditionStructArg{
				FieldName: fieldName,
				ZeroValue: zeroValue.FieldValue,
			})
		}
		tp.ParseAndUpdateSource(FuncEndTmpl, TemplateFuncStructArg{TypeName: structName})

		tp.ParseAndUpdateSource(FuncUpsertAllowDeleteStartTmpl, TemplateFuncStructArg{TypeName: structName})
		for fieldName, zeroValue := range conditions[structName] {
			tp.ParseAndUpdateSource(FuncUpsertAllowDeleteTmpl, TemplateUpsertConditionStructArg{
				FieldName: fieldName,
				ZeroValue: zeroValue.FieldValue,
			})
		}
		tp.ParseAndUpdateSource(FuncEndTmpl, TemplateFuncStructArg{TypeName: structName})

	}

	// 返り値はpackageが存在するパス, ファイル名, 生成したコード
	return tp, nil
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
		return nil, err
	}

	return files, nil
}

// SourceBuilderを受け取り、ファイルを生成する
func (sb *SourceBuilder) Generate() error {
	// ファイルが存在する場合、一度削除する
	fpath := filepath.Join(sb.PackagePath, sb.FileName)
	err := os.Remove(fpath)
	if err != nil && !os.IsNotExist(err) {
		// ファイルが存在しない以外のエラーの場合は、エラーを返す
		return err
	}
	file, err := os.Create(fpath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(sb.Source.GetSource())
	if err != nil {
		return err
	}
	return nil
}
