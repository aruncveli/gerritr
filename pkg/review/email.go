package review

import (
	"regexp"
	"strings"
)

const cfgKey = "allowedEmailDomains"

var cfgExists bool
var defaultEmailPattern = regexp.MustCompile(`.+@.+\..+`)
var allowedEmailDomains []string

func IsValidEmail(reviewer string) bool {
	if defaultEmailPattern.MatchString(reviewer) {
		if cfgExists {
			for _, allowedEmailDomain := range allowedEmailDomains {
				if strings.HasSuffix(reviewer, allowedEmailDomain) {
					return true
				}
			}
		} else {
			return true
		}
	} else {
		return false
	}
	return false
}

func SetAllowedEmailDomains() {
	cfgExists = Config.Exists(cfgKey)
	if cfgExists {
		for _, v := range Config.Strings(cfgKey) {
			allowedEmailDomains = append(allowedEmailDomains, "@"+v)
		}
	}
}
