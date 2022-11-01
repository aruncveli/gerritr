package action

import (
	"fmt"
	"gerritr/pkg/git"
	"gerritr/pkg/review"
	"os"
	"runtime"
	"strings"
)

func Push(branch string, state string, msg string, reviewers []string) []byte {

	var result []byte
	if notPatchFlow() && msg != "" {
		var err error
		result, err = git.Commit(msg)
		if err != nil {
			fmt.Printf("Cannot commit \n%s", err)
			os.Exit(1)
		}
	}

	if branch == "" {
		branch = git.GetUpstreamBranch()
	}

	branchRef := fmt.Sprintf("HEAD:refs/for/%s", branch)
	options := []string{state}

	var refSpec strings.Builder
	refSpec.WriteString(branchRef)

	if notPatchFlow() {
		options = append(options, review.GetReviewers(reviewers)...)
	}

	refSpec.WriteString("%" + strings.Join(options[:], ","))

	fmt.Println("Pushing to", branchRef)
	pushOutput, err := git.Push(refSpec.String())
	if err != nil {
		fmt.Printf("Cannot push to %s\n%s\n%s", branchRef, pushOutput, err)
		os.Exit(1)
	}
	result = append(result, pushOutput...)
	return result
}

func notPatchFlow() bool {
	_, file, _, _ := runtime.Caller(2)
	return !strings.Contains(file, "patch")
}
