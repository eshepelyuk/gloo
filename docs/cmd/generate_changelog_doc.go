package main

import (
	"context"
	"fmt"
	"os"

	"github.com/google/go-github/v31/github"
	"github.com/rotisserie/eris"
	. "github.com/solo-io/gloo/docs/cmd/changelogutils"
	. "github.com/solo-io/go-utils/versionutils"
	"github.com/spf13/cobra"
	"golang.org/x/oauth2"
)

func main() {
	ctx := context.Background()
	app := rootApp(ctx)
	if err := app.Execute(); err != nil {
		fmt.Printf("unable to run: %v\n", err)
		os.Exit(1)
	}
}

type options struct {
	ctx              context.Context
	HugoDataSoloOpts HugoDataSoloOpts
}

type HugoDataSoloOpts struct {
	product string
	version string
	// if set, will override the version when rendering the
	callLatest bool
	noScope    bool
}

func rootApp(ctx context.Context) *cobra.Command {
	opts := &options{
		ctx: ctx,
	}
	app := &cobra.Command{
		Use: "docs-util",
		RunE: func(cmd *cobra.Command, args []string) error {

			return nil
		},
	}
	app.AddCommand(changelogMdFromGithubCmd(opts))
	app.AddCommand(minorReleaseChangelogMdFromGithubCmd(opts))

	app.PersistentFlags().StringVar(&opts.HugoDataSoloOpts.version, "version", "", "version of docs and code")
	app.PersistentFlags().StringVar(&opts.HugoDataSoloOpts.product, "product", "gloo", "product to which the docs refer (defaults to gloo)")
	app.PersistentFlags().BoolVar(&opts.HugoDataSoloOpts.noScope, "no-scope", false, "if set, will not nest the served docs by product or version")
	app.PersistentFlags().BoolVar(&opts.HugoDataSoloOpts.callLatest, "call-latest", false, "if set, will use the string 'latest' in the scope, rather than the particular release version")

	return app
}

func changelogMdFromGithubCmd(opts *options) *cobra.Command {
	app := &cobra.Command{
		Use:   "gen-changelog-md",
		Short: "generate a markdown file from Github Release pages API",
		RunE: func(cmd *cobra.Command, args []string) error {
			if os.Getenv(skipChangelogGeneration) != "" {
				return nil
			}
			return generateChangelogMd(args)
		},
	}
	return app
}

func minorReleaseChangelogMdFromGithubCmd(opts *options) *cobra.Command {
	app := &cobra.Command{
		Use:   "gen-minor-releases-changelog-md",
		Short: "generate an aggregated changelog markdown file for each minor release version",
		RunE: func(cmd *cobra.Command, args []string) error {
			if os.Getenv(skipChangelogGeneration) != "" {
				return nil
			}
			return generateMinorReleaseChangelog(args)
		},
	}
	return app
}

const (
	latestVersionPath = "latest"
)

const (
	glooDocGen              = "gloo"
	glooEDocGen             = "glooe"
	skipChangelogGeneration = "SKIP_CHANGELOG_GENERATION"
)

const (
	glooOpenSourceRepo = "gloo"
	glooEnterpriseRepo = "solo-projects"
)

var (
	InvalidInputError = func(arg string) error {
		return eris.Errorf("invalid input, must provide exactly one argument, either '%v' or '%v', (provided %v)",
			glooDocGen,
			glooEDocGen,
			arg)
	}
	MissingGithubTokenError = func() error {
		return eris.Errorf("Must either set GITHUB_TOKEN or set %s environment variable to true", skipChangelogGeneration)
	}
)

// Generates changelog for releases as fetched from Github
// Github defaults to a chronological order
func generateChangelogMd(args []string) error {
	if len(args) != 1 {
		return InvalidInputError(fmt.Sprintf("%v", len(args)-1))
	}
	client := github.NewClient(nil)
	target := args[0]
	var repo string
	switch target {
	case glooDocGen:
		repo = glooOpenSourceRepo
	case glooEDocGen:
		repo = glooEnterpriseRepo
		ctx := context.Background()
		if os.Getenv("GITHUB_TOKEN") == "" {
			return MissingGithubTokenError()
		}
		ts := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
		)
		tc := oauth2.NewClient(ctx, ts)
		client = github.NewClient(tc)
	default:
		return InvalidInputError(target)
	}

	allReleases, err := GetAllReleases(client, repo, false)
	if err != nil {
		return err
	}

	for _, release := range allReleases {
		fmt.Printf("### %v\n\n", release.GetTagName())
		fmt.Printf("%v", release.GetBody())
	}
	return nil
}

// Performs additional processing to generate changelog grouped and ordered by release version
func generateMinorReleaseChangelog(args []string) error {
	if len(args) != 1 {
		return InvalidInputError(fmt.Sprintf("%v", len(args)-1))
	}
	target := args[0]
	var (
		err error
	)
	switch target {
	case glooDocGen:
		err = generateGlooChangelog()
	case glooEDocGen:
		err = generateGlooEChangelog()
	default:
		return InvalidInputError(target)
	}

	return err
}

// Fetches Gloo Open Source releases and orders them by version
func generateGlooChangelog() error {
	client := github.NewClient(nil)
	allReleases, err := GetAllReleases(client, glooOpenSourceRepo, true)
	if err != nil {
		return err
	}

	minorReleaseMap, err := ParseReleases(allReleases, true)
	if err != nil {
		return err
	}
	printVersionOrderReleases(minorReleaseMap)
	return nil
}

// Fetches Gloo Enterprise releases, merges in open source release notes, and orders them by version
func generateGlooEChangelog() error {
	// Initialize Auth
	ctx := context.Background()
	if os.Getenv("GITHUB_TOKEN") == "" {
		return MissingGithubTokenError()
	}
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	// Get all Gloo OSS release changelogs
	enterpriseReleases, err := GetAllReleases(client, glooEnterpriseRepo, true)
	if err != nil {
		return err
	}
	openSourceReleases, err := GetAllReleases(client, glooOpenSourceRepo, true)
	for _, i := range openSourceReleases {
		println(i.GetTagName())
	}
	return fmt.Errorf("some error")
	if err != nil {
		return err
	}
	minorReleaseMap, err := MergeEnterpriseOSSReleases(enterpriseReleases, openSourceReleases)
	if err != nil {
		return err
	}

	printVersionOrderReleases(minorReleaseMap)
	return nil
}

// Outputs changelogs in markdown format
func printVersionOrderReleases(minorReleaseMap map[Version]string) {
	var versions []Version
	for minorVersion, _ := range minorReleaseMap {
		versions = append(versions, minorVersion)
	}
	SortReleaseVersions(versions)
	for _, version := range versions {
		println(version.String())
		body := minorReleaseMap[version]
		fmt.Printf("\n\n### v%d.%d\n\n", version.Major, version.Minor)
		fmt.Printf("%v", body)
	}
}
