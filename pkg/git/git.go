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

// Tries to create a commit and exits if fails
func Commit(msg string) []byte {
	arg := append(commitWithMsgArgs, msg)
	output, err := run(arg...)
	if err != nil {
		fmt.Println("Cannot commit")
		fmt.Printf("%s", output)
		os.Exit(1)
	}
	return output
}

// Tries to push to origin and exits if fails
func Push(refSpec string) []byte {
	fmt.Println("Pushing with spec", refSpec)
	arg := append(pushArgs, refSpec)
	output, err := run(arg...)
	if err != nil {
		fmt.Println("Cannot push")
		fmt.Printf("%s", output)
		os.Exit(1)
	}
	return output
}

func amendWithMsg(msg string) []byte {
	arg := append(amendWithMsgArgs, msg)
	output, err := run(arg...)
	if err != nil {
		fmt.Println("Cannot amend")
		fmt.Printf("%s", output)
		os.Exit(1)
	}
	return output
}

func amendNoEdit() []byte {
	output, err := run(amendNoEditArgs...)
	if err != nil {
		fmt.Println("Cannot amend")
		fmt.Printf("%s", output)
		os.Exit(1)
	}
	return output
}

/*
Tries to amend the last commit and exits if fails. If msg parameter is non-empty, replaces the commit message.
*/
func Amend(msg string) []byte {
	if msg == "" {
		fmt.Println("Amending without editing the commit message")
		return amendNoEdit()
	}
	fmt.Println("Amending with message", msg)
	return amendWithMsg(msg)
}

/*
Checks if main or master branch is present in the repository - main before master. Exits if neither is present.

Uses [rev-parse --verify].

[rev-parse --verify]: https://git-scm.com/docs/git-rev-parse#Documentation/git-rev-parse.txt---verify
*/
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
