package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

func main() {
	output := make(map[string][]string)
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter date:\n ")
	text, _ := reader.ReadString('\n')
	since := "--since=" + text
	cmd := exec.Command("git", "log", since, "--pretty=format:\"%an#%s\"")
	stdoutStderr, _ := cmd.Output()
	stdoutStdStr := string(stdoutStderr)
	lines := strings.Split(stdoutStdStr, "\n")
	ticketNumPattern := regexp.MustCompile(`[A-Z]{3}-[0-9]{3,4}`)
	for _, line := range lines {
		original := strings.Split(line, "#")
		ticketNum := ticketNumPattern.FindStringSubmatch(original[1])
		if len(ticketNum) > 0 {
			num := ticketNum[0]
			if output[original[0]] == nil {
				output[original[0]] = append(output[original[0]], num)
			} else {
				if !Find(output[original[0]], num) {
					output[original[0]] = append(output[original[0]], num)
				}
			}
		}
	}

	for key, element := range output {
		fmt.Println(strings.Replace(key, "\"", "", -1), "=>", element, "\n")
	}
}

func Find(slice []string, val string) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}
