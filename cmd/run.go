package cmd

import (
	"aoc/pkg/aoc"
	"fmt"
	"reflect"

	"github.com/spf13/cobra"
)

var Demo bool

func init() {
	runCmd.Flags().BoolVarP(&Demo, "demo", "d", true, "Use demo input")
	rootCmd.AddCommand(runCmd)
}

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Runs the provided aoc algorithm",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		res1, res2 := aoc.Run(args[0], args[1], Demo)
		printResult(res1)
		printResult(res2)
	},
}

func printResult(t any) {
	switch reflect.TypeOf(t).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(t)
		for i := 0; i < s.Len(); i++ {
			fmt.Println(s.Index(i))
		}
	default:
		fmt.Printf("%v\n", t)
	}
}
