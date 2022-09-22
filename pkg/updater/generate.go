package updater

import (
	"fmt"
	"github.com/spf13/cobra"
)

var (
	generateCmd = &cobra.Command{
		Use:     "generate",
		Aliases: []string{"gen", "g"},
		Short:   "Generates metadata files for releases or patch",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("inDir: %s\n", inDir)
			fmt.Printf("outDir: %s\n", outDir)
		},
	}

	inDir  string
	outDir string
)

func initGenerate() {
	generateCmd.Flags().StringVarP(&inDir, "in", "i", ".", "release or patch directory")
	generateCmd.Flags().StringVarP(&outDir, "out", "o", ".", "metadata file output directory")
	rootCmd.AddCommand(generateCmd)
}
