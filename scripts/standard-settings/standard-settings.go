package main

import (
	"context"
	"fmt"
	"os"

	"github.com/google/go-github/v61/github"
	"github.com/northwood-labs/golang-utils/exiterrorf"
	"github.com/pkg/errors"
	flag "github.com/spf13/pflag"
)

var (
	ctx = context.Background()

	fRepo string
)

func main() {
	ghToken := os.Getenv("GITHUB_TOKEN")
	if ghToken == "" {
		exiterrorf.ExitErrorf(errors.New("GITHUB_TOKEN is required"))
	}

	flag.StringVarP(&fRepo, "repo", "r", "", "Repository to update inside NWL organization")
	flag.Parse()

	client := github.NewClient(nil).WithAuthToken(ghToken)

	_, _, err := client.Repositories.Edit(ctx, "northwood-labs", fRepo, &github.Repository{
		AllowAutoMerge:    github.Bool(true),
		AllowMergeCommit:  github.Bool(false),
		AllowRebaseMerge:  github.Bool(false),
		AllowSquashMerge:  github.Bool(true),
		AllowUpdateBranch: github.Bool(true),
		Archived:          github.Bool(false),
		// DefaultBranch:       github.String("main"),
		DeleteBranchOnMerge: github.Bool(true),
		HasIssues:           github.Bool(true),
		HasProjects:         github.Bool(false),
		HasWiki:             github.Bool(false),
		IsTemplate:          github.Bool(false),
		SecurityAndAnalysis: &github.SecurityAndAnalysis{
			SecretScanning: &github.SecretScanning{
				Status: github.String("enabled"),
			},
			SecretScanningPushProtection: &github.SecretScanningPushProtection{
				Status: github.String("enabled"),
			},
			SecretScanningValidityChecks: &github.SecretScanningValidityChecks{
				Status: github.String("enabled"),
			},
		},
	})
	if err != nil {
		exiterrorf.ExitErrorf(errors.Wrap(err, "failed to edit repository"))
	}

	fmt.Println("OK")
}
