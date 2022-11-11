package action

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/aruncveli/gerritr/pkg/git"
	"github.com/aruncveli/gerritr/pkg/review"
)

/*
Adds reviewers and pushes to origin

  - If the branch parameter is empty, checks if main or master branch is present in the repository
  - If the message parameter is non-empty, creates a new commit first, with the staged files and the given message

Returns the combined output of [commit] and [push] commands

[commit]: https://git-scm.com/docs/git-commit
[push]: https://git-scm.com/docs/git-push
*/
func Push(branch string, state string, msg string, reviewers []string) []byte {
	var commitOutput []byte
	if notPatchFlow() && msg != "" {
		commitOutput = git.Commit(msg)
	}

	if branch == "" {
		branch = git.GetUpstreamBranch()
	}

	options := []string{state}
	if notPatchFlow() {
		options = append(options, review.Resolve(reviewers)...)
	}
	optionsStr := strings.Join(options[:], ",")

	refSpec := fmt.Sprintf("HEAD:refs/for/%s%%%s", branch, optionsStr)
	pushOutput := git.Push(refSpec)

	return append(commitOutput, pushOutput...)
}

func notPatchFlow() bool {
	_, file, _, _ := runtime.Caller(2)
	return !strings.Contains(file, "patch")
}
