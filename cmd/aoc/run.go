/*

Copyright (c) 2021 - Present. Blend Labs, Inc. All rights reserved
Blend Confidential - Restricted

*/

package aoc

import (
	"aoc/pkg/aoc"
	"log"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(runCmd)
}

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Runs the provided aoc algorithm",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		res1, res2 := aoc.Run(args[0], args[1])
		log.Printf("Part 1: %d, Part 2: %d\n", res1, res2)
	},
}
