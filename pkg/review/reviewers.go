package review

import (
	"bufio"
	"fmt"
	"os"
)

const aliasPrefix = "alias."
const reviewersFilename = "REVIEWERS"

/*
Resolves reviewers from the input parameters and REVIEWERS file

For each string in the input slice,
  - If it is an email ID, adds as is
  - If it is not an email ID, assumes it to be a alias. Tries to resolve the alias from config

Then, reads REVIEWERS file and adds each line. Returns a slice of resolved reviewers, with each element prefixed with "r=", as per [the format required by Gerrit].

[the format required by Gerrit]: https://gerrit-documentation.storage.googleapis.com/Documentation/3.6.2/user-upload.html#reviewers
*/
func ResolveReviewers(reviewersInput []string) []string {

	fmt.Println("Resolving reviewers")
	var resolvedReviewers []string

	if len(reviewersInput) != 0 {
		for _, reviewer := range reviewersInput {
			if IsEmail(reviewer) {
				resolvedReviewers = append(resolvedReviewers, reviewer)
			} else {
				alias := aliasPrefix + reviewer
				if Config.Exists(alias) {
					reviewersFromCfg := Config.Strings(alias)
					resolvedReviewers = append(resolvedReviewers, reviewersFromCfg...)
				}
			}
		}
	}

	fmt.Println("Opening", reviewersFilename)
	fReviewers, err := os.Open(reviewersFilename)
	if err != nil {
		fmt.Println("Cannot", err)
	} else {
		fmt.Println("Reading", reviewersFilename)
		scanner := bufio.NewScanner(fReviewers)
		for scanner.Scan() {
			reviewer := scanner.Text()
			resolvedReviewers = append(resolvedReviewers, reviewer)
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
