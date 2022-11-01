package action

import (
	"fmt"
	"gerritr/pkg/git"
	"os"
)

func Patch(branch string, state string, msg string) []byte {

	fmt.Println("Amending the last commit")
	result, err := git.Amend(msg)
	if err != nil {
		fmt.Printf("Cannot amend the review\n%s\n%s", result, err)
		os.Exit(1)
	}

	pushOutput := Push(branch, state, msg, nil)
	result = append(result, pushOutput...)
	return result
}
