package review

import (
	"bufio"
	"fmt"
	"os"
)

func GetReviewers(reviewersFromCmd []string) []string {
	fmt.Println("Resolving reviewers")
	var resolvedReviewers []string

	var validateAndAdd = func(reviewer string) bool {
		if IsValidEmail(reviewer) {
			resolvedReviewers = append(resolvedReviewers, reviewer)
			return true
		}
		return false
	}

	if len(reviewersFromCmd) != 0 {
		for _, reviewerFromCmd := range reviewersFromCmd {
			if !validateAndAdd(reviewerFromCmd) {
				teamKey := "teams." + reviewerFromCmd
				if Config.Exists(teamKey) {
					reviewersFromCfg := Config.Strings(teamKey)
					for _, reviewerFromCfg := range reviewersFromCfg {
						validateAndAdd(reviewerFromCfg)
					}
				}
			}
		}
	}

	fmt.Println("Opening REVIEWERS")
	fReviewers, err := os.Open("REVIEWERS")
	if err != nil {
		fmt.Println("Cannot", err)
	} else {
		fmt.Println("Reading REVIEWERS")
		scanner := bufio.NewScanner(fReviewers)
		for scanner.Scan() {
			reviewer := scanner.Text()
			validateAndAdd(reviewer)
		}
	}
	defer fReviewers.Close()

	nReviewers := len(resolvedReviewers)
	if nReviewers == 0 {
		fmt.Println("No valid reviewers to add")
		return nil
	} else {
		fmt.Println("Adding", nReviewers, "reviewers")
	}

	for i, reviewer := range resolvedReviewers {
		resolvedReviewers[i] = "r=" + reviewer
	}

	return resolvedReviewers
}
