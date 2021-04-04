package main

import (
	"os"
	"testing"

	"github.com/gregoryv/wolf"
)

func Test_main_help(t *testing.T) {
	os.Args = append(os.Args, "--help")
	main()
}

func Test_run_help(t *testing.T) {
	cmd := wolf.NewTCmd("starplan", "--help")
	run(cmd)
	if cmd.ExitCode != 0 {
		t.Error(cmd.Dump(), cmd.ExitCode)
	}
}

func Test_run_badarg(t *testing.T) {
	cmd := wolf.NewTCmd("starplan", "--no-such")
	run(cmd)
	if cmd.ExitCode != 1 {
		t.Error(cmd.Dump(), cmd.ExitCode)
	}
}

func Test_run(t *testing.T) {
	cmd := wolf.NewTCmd("starplan", "--bind", "80")
	run(cmd)
	if cmd.ExitCode != 1 {
		t.Error(cmd.Dump(), cmd.ExitCode)
	}
}
