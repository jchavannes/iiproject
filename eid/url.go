package eid

import (
	"strings"
	"regexp"
)

func ConvertShortEidUrlIntoFull(eidUrl string) string {
	if ! strings.Contains(eidUrl, "@") {
		return eidUrl
	}

	reg := regexp.MustCompile(`^[^@]+@`)
	username := reg.FindString(eidUrl)
	username = username[:len(username) - 1]
	domain := reg.ReplaceAllString(eidUrl, "")
	realEidUrl := domain + "/id/" + username
	return realEidUrl
}

func ConvertFullEidUrlIntoShort(eidUrl string) string {
	reg := regexp.MustCompile(`^(?P<domain>[^/]+)/id/(?P<name>[^/]+)$`)
	if ! reg.MatchString(eidUrl) {
		return eidUrl

	}
	shortEidUrl := reg.ReplaceAllString(eidUrl, "${name}@${domain}")
	return shortEidUrl
}
