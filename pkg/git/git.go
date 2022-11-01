package git

import (
	"fmt"
	"os"
	"os/exec"
)

var ExecName = "git"

var branches = [...]string{"main", "master"}
var checkBranchArgs = []string{"rev-parse", "--verify"}

const commit = "commit"
const msgFlag = "--message"

var commitWithMsgArgs = []string{commit, msgFlag}
var pushArgs = []string{"push", "origin"}

var amendArgs = []string{commit, "--amend"}
var amendNoEditArgs = append(amendArgs, "--no-edit")
var amendWithMsgArgs = append(amendArgs, msgFlag)

func run(arg ...string) ([]byte, error) {
	return exec.Command(ExecName, arg...).CombinedOutput()
}

func Commit(msg string) ([]byte, error) {
	arg := append(commitWithMsgArgs, msg)
	return run(arg...)
}

func Push(refSpec string) ([]byte, error) {
	arg := append(pushArgs, refSpec)
	return run(arg...)
}

func Amend(msg string) ([]byte, error) {
	if msg == "" {
		return run(amendNoEditArgs...)
	}
	arg := append(amendWithMsgArgs, msg)
	return run(arg...)
}

func GetUpstreamBranch() string {

	for _, branch := range branches {
		arg := append(checkBranchArgs, branch)
		_, err := run(arg...)
		if err == nil {
			fmt.Println("Setting branch as", branch)
			return branch
		}
	}

	fmt.Println("Target branch is not specified and not resolvable")
	os.Exit(1)
	return ""
}
