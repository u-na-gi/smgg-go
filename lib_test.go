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

}
