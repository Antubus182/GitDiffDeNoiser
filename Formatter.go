package main

import (
	"fmt"
	"regexp"
	"strings"
)

func FormatDiff(line string) string {
	regRemoved := regexp.MustCompile("(?m)^[-]{1}([^-]|$)")
	regAdded := regexp.MustCompile("(?m)^[+]{1}([^+]|$)")
	regStart := regexp.MustCompile(`(?m)^(diff --git)`)
	//formatted := ""
	if strings.Contains(line, "LastEditTime") {
		fmt.Println("Timestamp")
		line = "<p class='timeline'>" + line + "</p>"
	} else if regStart.MatchString(line) {
		line = "</div><div class='startsection'>" + line
	} else if regAdded.MatchString(line) {
		fmt.Println("Green Line")
		line = "<p class='addedline'>" + line + "</p>"
	} else if regRemoved.MatchString(line) {
		fmt.Println("Red line")
		line = "<p class='removedline'>" + line + "</p>"
	} else {
		line = "<p>" + line + "</p>"
	}

	return line
}

func FileCount(line string) int {
	if strings.Contains(line, "diff --git") {
		return 1
	} else {
		return 0
	}
}

func Verify(inputs DiffData) bool {
	re := regexp.MustCompile("^[0-9a-fA-F]{40}$")
	if !re.MatchString(inputs.Sha1) || !re.MatchString(inputs.Sha2) || inputs.Sha1 == inputs.Sha2 {
		return false
	}
	return true
}
