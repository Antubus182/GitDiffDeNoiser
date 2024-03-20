package main

import (
	"fmt"
	"html"
	"os/exec"
	"regexp"
	"strings"
)

func RunDiff(d DiffData) string {
	// gitdiff <sha> <sha>

	v1 := d.Sha1
	v2 := d.Sha2
	directroy := d.Dir
	//difcom := "git diff 903fc35001c923faf396ccf299bf81287094e926 5d0f6fd61ac2bc5f4dd9844dfdfdcb570fed2079"
	//fmt.Println("excucute the following command: ")
	//fmt.Println(difcom)

	cmd := exec.Command("git", "diff", v1, v2) //the command to be executed
	cmd.Dir = directroy                        //the directory in which the command should be executed
	stdout, err := cmd.Output()                //The actual execution of the command

	if err != nil {
		fmt.Println(err.Error())
		return ""
	}

	arrayofStings := strings.Split((string(stdout)), "\n")
	addcounter := 0
	delcounter := 0
	//fmt.Print(string(stdout))
	fmt.Println("Counting added lines: ")
	for _, str := range arrayofStings {
		if strings.Contains(str, "+") {
			addcounter++
		} else if strings.Contains(str, "-") {
			delcounter++
		}
	}

	fmt.Println(addcounter)
	paratags := StringToPara(arrayofStings)
	return paratags
}

func StringToPara(lines []string) string {
	para := ""
	for _, str := range lines {
		str = EscapeHTML(str)
		para += "<p>" + str + "</p>"
	}

	return para
}

func VerifyInputs(inputs DiffData) bool {
	re := regexp.MustCompile("^[0-9a-fA-F]{40}$")
	if !re.MatchString(inputs.Sha1) || !re.MatchString(inputs.Sha2) || inputs.Sha1 == inputs.Sha2 {
		return false
	}
	return true
}

func EscapeHTML(raw string) string {
	return html.EscapeString(raw)
}
