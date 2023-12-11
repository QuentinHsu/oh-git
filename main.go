package main

import (
	"fmt"
	"strings"

	"github.com/QuentinHsu/ohgit/pkg/info"
	"github.com/QuentinHsu/ohgit/pkg/logger"
	"github.com/QuentinHsu/ohgit/pkg/read"

	"github.com/fatih/color"
)

var ()

type Commit struct {
	Hash       string
	Author     string
	Message    string
	CommitDate string
}

func main() {
	logger := &logger.Logger{}

	logger.Info(" Welcome to ohgit ", color.BgGreen)

	logger.Info(fmt.Sprintf("\nVersion: %s, Release: %s\n", info.Version, info.Release), color.FgGreen)

	commits := read.Git()
	if len(commits) == 0 || (len(commits) == 1 && commits[0] == "") {
		logger.Warn("No results.\n")
		return
	} else {
		logger.Label("Number of commits: ")
		logger.Value(fmt.Sprintf("%d\n\n", len(commits)))
	}
	for _, commit := range commits {
		if commit != "" {
			fields := strings.Split(commit, "|")
			if len(fields) == 4 {
				commit := Commit{
					Hash:       fields[0],
					Author:     fields[1],
					Message:    fields[2],
					CommitDate: fields[3],
				}

				logger.Label("Hash:       ")
				logger.Value(fmt.Sprintf("%s\n", commit.Hash))
				logger.Label("Author:     ")
				logger.Value(fmt.Sprintf("%s\n", commit.Author))
				logger.Label("Message:    ")
				logger.Value(fmt.Sprintf("%s\n", commit.Message))
				logger.Label("CommitDate: ")
				logger.Value(fmt.Sprintf("%s\n\n", commit.CommitDate))

			}
		}
	}
}
