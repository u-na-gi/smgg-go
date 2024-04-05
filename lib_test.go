package smgg

import (
	"fmt"
	"testing"
)

func TestXxx(t *testing.T) {
	t.Run("TestXxx", func(t *testing.T) {
		sourcePaths, err := WalkingCurrentProject()
		if err != nil {
			t.Fatal(err)
		}
		for _, sourcePath := range sourcePaths {
			fmt.Println("sourcePath", sourcePath)
		}
	})

	t.Run("TestXxx34242", func(t *testing.T) {
		sourcePaths, err := WalkingCurrentProject()
		if err != nil {
			t.Fatal(err)
		}
		byPackageGenSource, err := AggregateByPackageName(sourcePaths)
		if err != nil {
			t.Fatal(err)
		}

		for _, sourceBuilder := range byPackageGenSource {
			source := sourceBuilder.Source
			for _, sourcePath := range sourceBuilder.SourcePaths {
				crs, err := source.CreateSource(sourcePath)
				if err != nil {
					t.Fatal(err)
				}
				sourceBuilder.Source = crs
			}
			sourceBuilder.Generate()

		}
	})

}
