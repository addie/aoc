/*

Copyright (c) 2021 - Present. Blend Labs, Inc. All rights reserved
Blend Confidential - Restricted

*/

package clipboard

import "os/exec"

var (
	copyCmdArgs = "pbcopy"
)

// WriteAll write string to clipboard
func WriteAll(text string) error {
	return writeAll(text)
}

func writeAll(text string) error {
	copyCmd := getCopyCommand()
	in, err := copyCmd.StdinPipe()
	if err != nil {
		return err
	}

	if err := copyCmd.Start(); err != nil {
		return err
	}
	if _, err := in.Write([]byte(text)); err != nil {
		return err
	}
	if err := in.Close(); err != nil {
		return err
	}
	return copyCmd.Wait()
}

func getCopyCommand() *exec.Cmd {
	return exec.Command(copyCmdArgs)
}
