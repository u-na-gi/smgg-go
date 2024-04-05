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

	t.Run("TestXxx", func(t *testing.T) {
		sourcePaths, err := WalkingCurrentProject()
		if err != nil {
			t.Fatal(err)
		}
		for _, sourcePath := range sourcePaths {
			sb, err := CreateSource(sourcePath)
			if err != nil {
				panic(err)
			}

			if err := sb.Generate(); err != nil {
				panic(err)
			}
		}
	})

}
