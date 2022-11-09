package action

import (
	"github.com/aruncveli/gerritr/pkg/git"
)

/*
Amends the last commit and pushes it. Returns the combined output of [commit --amend] and [push] commands.

[commit --amend]: https://git-scm.com/docs/git-commit#Documentation/git-commit.txt---amend
[push]: https://git-scm.com/docs/git-push
*/
func Patch(branch string, state string, msg string) []byte {
	amendOutput := git.Amend(msg)
	pushOutput := Push(branch, state, msg, nil)
	return append(amendOutput, pushOutput...)
}
