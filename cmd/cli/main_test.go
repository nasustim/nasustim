package main_test

import (
	"testing"

	cli "github.com/nasustim/nasustim/cmd/cli"
)

func Test_Run(t *testing.T) {
	t.Run("exit successfully", func(t *testing.T) {
		actual := cli.Run("../../tmpl/README.go.md", "../../README.md")
		if actual != cli.EXIT_SUCCESS {
			t.Error("unexpected result")
		}
	})
}
