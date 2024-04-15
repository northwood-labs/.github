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
	ctx      = context.Background()
	isPublic bool

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

	repo, _, err := client.Repositories.Get(ctx, "northwood-labs", fRepo)
	if err != nil {
		exiterrorf.ExitErrorf(errors.Wrap(err, "failed to get repository"))
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
		exiterrorf.ExitErrorf(errors.Wrap(err, "failed to edit repository"))
	}

	fmt.Println("OK")
}
