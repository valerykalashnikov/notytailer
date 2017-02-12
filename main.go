package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strings"

	"github.com/hpcloud/tail"
)

func main() {

	// var pattern = `HTTP\/1.0" (4\d{2}|5\d{2}|6\d{2})`

	var subject, recepient, pattern, logFile string

	flag.StringVar(&subject, "subject", "", "subject of email notifications")
	flag.StringVar(&recepient, "recepient", "", "recepient of email notifications")
	flag.StringVar(&pattern, "pattern", "", "regexp to find the line in the log file")
	flag.StringVar(&logFile, "logfile", "", "regexp to find the line in the log file")

	flag.Parse()

	var flagIsEmpty = (subject == "" || recepient == "" || pattern == "" || logFile == "")

	if flagIsEmpty {
		fmt.Println(usageOutput)
		os.Exit(1)
	}

	t, err := tail.TailFile(logFile, tail.Config{Follow: true, ReOpen: true})

	if err != nil {
		log.Fatal("Unable to tail log file", "-", err)
	}

	for line := range t.Lines {
		matched, err := regexp.MatchString(pattern, line.Text)

		if err != nil {
			log.Println("Unable to match line", "-", err)
		}

		if matched {
			sendEmail(subject, recepient, line.Text)
		}
	}
}

func sendEmail(subject, recepient, message string) {
	cmd := exec.Command("mail", "-s", subject, recepient)
	cmd.Stdin = strings.NewReader(message)
	err := cmd.Run()
	if err != nil {
		log.Println("Error when running mail comand", "-", err)
	}
	fmt.Println("Sent to", recepient, ":", message)
}

const usageOutput = `Usage: ./tailer -subject 'This is the subject' -recepient 'recepient@example.com' -pattern 'your_regex_here' -logfile '/var/log/example.log'`
