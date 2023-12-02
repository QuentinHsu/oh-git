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
	statDays := flag.Int("stat-day", 1, "number of days to include in the stats")
	filterUser := flag.String("user", "", "filter commits by user")
	flag.Parse()
	if *repoPath == "" {
		wd, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
		repoPath = &wd
	}
	loc, err := time.LoadLocation("") // 加载系统的时区
	fmtStrDay := "2006-01-02 15:04:05"
	endDate := time.Now().In(loc).Add(24 * time.Hour).Truncate(24 * time.Hour).Add(-time.Second)
	endDateStr := endDate.Format(fmtStrDay)

	startDate := endDate.AddDate(0, 0, (-*statDays)).Add(+time.Second)
	startDateStr := startDate.Format(fmtStrDay)

	cmdArgs := []string{"log", "--pretty=format:%H|%an|%s", fmt.Sprintf("--since=%s", startDateStr), fmt.Sprintf("--until=%s", endDateStr)}
	if *filterUser != "" {
		cmdArgs = append(cmdArgs, fmt.Sprintf("--author=%s", *filterUser))
	}
	cmd := exec.Command("git", cmdArgs...)
	cmd.Dir = *repoPath
	output, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	duration := endDate.Sub(startDate)
	// 将小时数转换为天数
	days := int(duration.Hours()/24) + 1
	fmt.Printf("Stat Range: %s - %s (%d days)\n\n", startDateStr, endDateStr, days)
	commits := strings.Split(string(output), "\n")
	if len(commits) == 0 || (len(commits) == 1 && commits[0] == "") {
		fmt.Println("No results.")
		return
	} else {
		fmt.Printf("Number of commits: %d\n\n", len(commits))
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
