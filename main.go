package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/fatih/color"
)

type Commit struct {
	Hash    string
	Author  string
	Message string
}

func main() {
	// 创建一个带有颜色的 Logger
	logger := log.New(color.Output, "", log.LstdFlags)
	// 设置日志输出颜色
	// logger.SetPrefix(color.GreenString("[INFO] "))

	logger.SetFlags(logger.Flags() &^ (log.Ldate | log.Ltime)) // 移除默认的日期和时间标记
	// 设置标题为绿色，值为黄色
	logTitleColor := color.New(color.FgGreen).PrintfFunc()
	logValueColor := color.New(color.FgYellow).PrintfFunc()

	logger.Println(color.GreenString("\nohgit v1.0.0\n"))

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
	if err != nil {
		log.Fatal(err)
	}
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
	logTitleColor("Stat Range: ")

	logValueColor("%s - %s (%d days)\n\n", startDateStr, endDateStr, days)
	commits := strings.Split(string(output), "\n")
	if len(commits) == 0 || (len(commits) == 1 && commits[0] == "") {
		logValueColor("No results.")
		return
	} else {
		logTitleColor("Number of commits: ")
		logValueColor("%d\n\n", len(commits))
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

				logTitleColor("Hash: ")
				logValueColor("%s\n", commit.Hash)
				logTitleColor("Author: ")
				logValueColor("%s\n", commit.Author)
				logTitleColor("Message: ")
				logValueColor("%s\n", commit.Message)
				println("")
			}
		}
	}
}
