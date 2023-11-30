package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
)

type Commit struct {
	Hash    string
	Author  string
	Message string
}

func main() {
	repoPath := flag.String("p", "", "repository path")
	statDays := flag.Int("stat-day", 7, "number of days to include in the stats")
	flag.Parse()
	if *repoPath == "" {
		wd, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
		repoPath = &wd
	}
	cmd := exec.Command("git", "log", "--pretty=format:%H|%an|%s", fmt.Sprintf("--since=%d.days.ago", *statDays))
	cmd.Dir = *repoPath
	output, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	startDate := time.Now().AddDate(0, 0, -*statDays).Format("2006-01-02")
	endDate := time.Now().Format("2006-01-02")
	duration := time.Now().Sub(time.Now().AddDate(0, 0, -*statDays))
	// 将小时数转换为天数
	days := int(duration.Hours() / 24)
	fmt.Printf("Stat Range: %s - %s (%d days)\n\n", startDate, endDate, days)
	commits := strings.Split(string(output), "\n")
	fmt.Printf("Number of commits: %d\n\n", len(commits))
	if len(commits) == 0 || (len(commits) == 1 && commits[0] == "") {
		fmt.Println("\nNo results.")
		return
	}
	for _, commit := range commits {
		if commit != "" {
			fields := strings.Split(commit, "|")
			if len(fields) == 3 {
				commit := Commit{
					Hash:    fields[0],
					Author:  fields[1],
					Message: fields[2],
				}
				fmt.Printf("Hash: %s\nAuthor: %s\nMessage: %s\n\n", commit.Hash, commit.Author, commit.Message)
			}
		}
	}
}
