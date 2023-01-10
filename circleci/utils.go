package circleci

import (
	"regexp"
)

var (
	githubRegex, _    = regexp.Compile("^https://github")
	bitbucketRegex, _ = regexp.Compile("^https://bitbucket")
)

func Slugify(vcsUrl, vcsUserName, vcsRepoName string) (string, string) {
	var vcsSlug string
	githubMatch := githubRegex.MatchString(vcsUrl)
	if githubMatch {
		vcsSlug = "gh"
	} else {
		bitbucketMatch := bitbucketRegex.MatchString(vcsUrl)
		if bitbucketMatch {
			vcsSlug = "bb"
		}
	}

	var organizationSlug string
	if vcsSlug != "" {
		organizationSlug = vcsSlug + "/" + vcsUserName
	}

	if vcsRepoName == "" {
		return organizationSlug, ""
	}

	var projectSlug string
	if organizationSlug != "" {
		projectSlug = organizationSlug + "/" + vcsRepoName
	}

	return organizationSlug, projectSlug
}
