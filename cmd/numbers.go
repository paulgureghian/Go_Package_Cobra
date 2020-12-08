package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	countFlagNumbers int
	rangeFlagNumbers []string
)

// numbersCmd represents the numbers command
var numbersCmd = &cobra.Command{
	Use:   "numbers",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("numbers called")
		fmt.Println("--count:", countFlagNumbers)
		fmt.Println("--range:", rangeFlagNumbers)
		fmt.Println("--verbose:", verbose)
	},
}

func init() {
	rootCmd.AddCommand(numbersCmd)

	numbersCmd.Flags().IntVarP(
		&countFlag, "count", "c", 0,
		"A count of random numbers",
	)

	numbersCmd.MarkFlagRequired("count")

	numbersCmd.Flags().StringSliceVarP(
		&rangeFlag, "range", "r", []string{"1:100"},
		"Range of numbers. Optional",
	)
}