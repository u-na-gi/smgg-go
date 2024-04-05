package smggclilib

import (
	"github.com/spf13/cobra"
	smgg "github.com/u-na-gi/smgg-go"
)

var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "Generate a new resource",
	Long:  `Generate a new resource. `,
	Run: func(cmd *cobra.Command, args []string) {
		sourcePaths, err := smgg.WalkingCurrentProject()
		if err != nil {
			panic(err)
		}
		for _, sourcePath := range sourcePaths {
			sb, err := smgg.CreateSource(sourcePath)
			if err != nil {
				panic(err)
			}

			if err := sb.Generate(); err != nil {
				panic(err)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(genCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
