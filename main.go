package main

import (
	"errors"
	"os"
	"os/exec"

	"github.com/gabe565/moreutils/cmd"
	"github.com/gabe565/moreutils/internal/cmdutil"
	"github.com/gabe565/moreutils/internal/util"
)

var version = "beta"

func main() {
	root := cmd.New(os.Args[0], cmdutil.WithVersion(version))
	root.SilenceErrors = true
	if err := root.Execute(); err != nil {
		var execErr *exec.ExitError
		if errors.As(err, &execErr) {
			os.Exit(execErr.ExitCode())
		}

		var exitCodeErr *util.ExitCodeError
		if errors.As(err, &exitCodeErr) {
			os.Exit(exitCodeErr.ExitCode())
		}

		root.PrintErrln(root.ErrPrefix(), err.Error())
		os.Exit(1)
	}
}
