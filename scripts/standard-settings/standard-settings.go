package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/charmbracelet/log"
	"github.com/google/go-github/v61/github"
	flag "github.com/spf13/pflag"
)

var (
	ctx      = context.Background()
	isPublic bool

	fRepo string

	logger = log.NewWithOptions(os.Stderr, log.Options{
		ReportTimestamp: true,
		TimeFormat:      time.Kitchen,
		Prefix:          "standard-settings",
	})
)

func main() {
	ghToken := os.Getenv("GITHUB_TOKEN")
	if ghToken == "" {
		logger.Fatal("GITHUB_TOKEN is required")
	}

	flag.StringVarP(&fRepo, "repo", "r", "", "Repository to update inside NWL organization")
	flag.Parse()

	client := github.NewClient(nil).WithAuthToken(ghToken)

	repo, _, err := client.Repositories.Get(ctx, "northwood-labs", fRepo)
	if err != nil {
		logger.Fatalf("failed to get repository: %s", err.Error())
	}

	isPublic = repo.GetVisibility() == "public"

	_, _, err = client.Repositories.Edit(ctx, "northwood-labs", fRepo, &github.Repository{
		AllowAutoMerge:      github.Bool(true),
		AllowMergeCommit:    github.Bool(false),
		AllowRebaseMerge:    github.Bool(false),
		AllowSquashMerge:    github.Bool(true),
		AllowUpdateBranch:   github.Bool(true),
		Archived:            github.Bool(false),
		DeleteBranchOnMerge: github.Bool(true),
		HasIssues:           github.Bool(true),
		HasProjects:         github.Bool(false),
		HasWiki:             github.Bool(false),
		IsTemplate:          github.Bool(false),
		// SecurityAndAnalysis: &github.SecurityAndAnalysis{
		// 	AdvancedSecurity: &github.AdvancedSecurity{
		// 		Status: github.String(func() string {
		// 			if isPublic {
		// 				return "enabled"
		// 			}

		// 			return "disabled"
		// 		}()),
		// 	},
		// 	SecretScanning: &github.SecretScanning{
		// 		Status: github.String(func() string {
		// 			if isPublic {
		// 				return "enabled"
		// 			}

		// 			return "disabled"
		// 		}()),
		// 	},
		// 	SecretScanningPushProtection: &github.SecretScanningPushProtection{
		// 		Status: github.String(func() string {
		// 			if isPublic {
		// 				return "enabled"
		// 			}

		// 			return "disabled"
		// 		}()),
		// 	},
		// 	SecretScanningValidityChecks: &github.SecretScanningValidityChecks{
		// 		Status: github.String(func() string {
		// 			if isPublic {
		// 				return "enabled"
		// 			}

		// 			return "disabled"
		// 		}()),
		// 	},
		// },
	})
	if err != nil {
		logger.Fatalf("failed to edit repository: %s", err.Error())
	}

	fmt.Println("OK")
}
