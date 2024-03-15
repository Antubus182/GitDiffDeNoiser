package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func RunDiff() {

	fmt.Println("Git diff")
	fmt.Println("Cli tool to run git diff")
	fmt.Println("gitdiff <sha> <sha>")
	fmt.Println("excucute the following command: ")
	difcom := "git diff 903fc35001c923faf396ccf299bf81287094e926 5d0f6fd61ac2bc5f4dd9844dfdfdcb570fed2079"
	fmt.Println(difcom)
	v1 := "903fc35001c923faf396ccf299bf81287094e926"
	v2 := "5d0f6fd61ac2bc5f4dd9844dfdfdcb570fed2079"
	directroy := "C:\\Users\\niels.ten.thije\\OneDrive - Zeton BV\\Documenten\\lets-go\\Code\\Lets-go"
	fmt.Println(directroy)
	cmd := exec.Command("git", "diff", v1, v2)
	cmd.Dir = directroy
	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	arrayofStings := strings.Split((string(stdout)), "\n")
	addcounter := 0
	delcounter := 0
	fmt.Print(string(stdout))
	fmt.Println("Counting added lines: ")
	for _, str := range arrayofStings {
		if strings.Contains(str, "+") {
			addcounter++
		} else if strings.Contains(str, "-") {
			delcounter++
		}
	}
	fmt.Println(addcounter)
}
