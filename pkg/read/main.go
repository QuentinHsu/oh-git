package read

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/QuentinHsu/ohgit/pkg/logger"

	"github.com/fatih/color"
)

func Git() []string {
	logger := &logger.Logger{}
	repoPath := flag.String("path", "", "repository path")
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
	localTimezone, err := time.LoadLocation("") // 加载系统的时区
	if err != nil {
		log.Fatal(err)
	}
	cmdIsGitWorkSpace := exec.Command("git", "rev-parse", "--is-inside-work-tree")
	cmdIsGitWorkSpace.Dir = *repoPath
	err = cmdIsGitWorkSpace.Run()
	if err != nil {
		if _, ok := err.(*exec.ExitError); ok {
			logger.Error("\nThe current directory is not a git repository.\n")
			os.Exit(1)
		}
		logger.Error(err.Error())
	}
	fmtStrDay := "2006-01-02 15:04:05"
	endDate := time.Now().In(localTimezone).Add(24 * time.Hour).Truncate(24 * time.Hour).Add(-time.Second)
	endDateStr := endDate.Format(fmtStrDay)

	startDate := endDate.AddDate(0, 0, (-*statDays)).Add(+time.Second)
	startDateStr := startDate.Format(fmtStrDay)

	cmdArgs := []string{"log", "--pretty=format:%H|%an|%s|%ci", fmt.Sprintf("--since=%s", startDateStr), fmt.Sprintf("--until=%s", endDateStr)}
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

	logger.Info("\nStat Range: ", color.FgGreen)
	logger.Info(fmt.Sprintf("%s - %s (%d days)\n\n", startDateStr, endDateStr, days), color.FgYellow)

	commits := strings.Split(string(output), "\n")
	return commits
}
