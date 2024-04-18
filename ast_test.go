package smgg

import (
	"fmt"
	"testing"
)

func TestParseAlias(t *testing.T) {
	t.Run("aliasがあっても処理ができる", func(t *testing.T) {
		filename := "example/source/source2.go"

		res, err := ParseAstFromFile(filename)
		if err != nil {
			t.Fatal(err)
		}
		fmt.Printf("res: %#+v\n", res)

		expected := &StructsInfo{
			PackageName: "source",
			StructNames: []string{"Bird"},
			Conditions: map[string]map[string]*ConditionType{
				"Bird": {
					"Name": {
						FieldValue: "string",
						IsPointer:  true,
					},
				},
			},
		}

		if res.PackageName != expected.PackageName {
			t.Errorf("PackageName: got %s, expected %s", res.PackageName, expected.PackageName)
		}
		if len(res.StructNames) != len(expected.StructNames) {
			t.Errorf("StructNames: got %v, expected %v", res.StructNames, expected.StructNames)
		}
		for i, v := range res.StructNames {
			if v != expected.StructNames[i] {
				t.Errorf("StructNames[%d]: got %s, expected %s", i, v, expected.StructNames[i])
			}
		}
		for k, v := range res.Conditions {
			fmt.Printf("k: %s,  v: %#v\n", k, v)
			for kk, vv := range v {
				fmt.Printf("kk: %s,  vv: %#v\n", kk, vv)
			}
		}
	})
}
