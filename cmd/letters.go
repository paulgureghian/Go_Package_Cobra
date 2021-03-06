package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	countFlagLetters int
	langFlagLetters  string
)

// lettersCmd represents the letters command
var lettersCmd = &cobra.Command{
	Use:   "letters",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("letters mode")
		fmt.Println("--count:", countFlagLetters)
		fmt.Println("--lang:", langFlagLetters)
		fmt.Println("--verbose:", verbose)
	},
}

func init() {
	rootCmd.AddCommand(lettersCmd)

	lettersCmd.Flags().IntVarP(
		&countFlagLetters, "count", "c", 0,
		"A count of random letters",
	)

	lettersCmd.MarkFlagRequired("count")

	lettersCmd.Flags().StringVarP(
		&langFlagLetters, "lang", "l", "en",
		"A language. Optional",
	)
}
