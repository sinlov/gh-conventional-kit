package constant

import "github.com/urfave/cli/v2"

type BadgeConfig struct {
	NoCommonBadges bool
	GolangBadges   bool
	RustBadges     bool
	RustVersion    string
	NodeBadges     bool
	NpmPackage     string
	DockerUser     string
	DockerRepo     string
}

func BadgeFlags() []cli.Flag {
	return []cli.Flag{
		&cli.BoolFlag{
			Name:  "no-common-badges",
			Value: false,
			Usage: "no badges common subcommand for this repo",
		},

		&cli.BoolFlag{
			Name:  "golang",
			Usage: "golang badges for this repo",
		},

		&cli.BoolFlag{
			Name:  "rust",
			Usage: "rust badges for this repo",
		},
		&cli.StringFlag{
			Name:  "rust-version",
			Usage: "rust version badges for this repo",
		},

		&cli.BoolFlag{
			Name:  "node",
			Usage: "node badges for this repo, default blank is invalid",
		},

		&cli.StringFlag{
			Name:  "npm",
			Usage: "npm badges for this repo, default blank is invalid",
			Value: "",
		},

		&cli.StringFlag{
			Name:  "docker-user",
			Usage: "docker user for this repo, default blank is invalid",
			Value: "",
		},

		&cli.StringFlag{
			Name:  "docker-repo",
			Usage: "docker repo for this repo, --docker-user must be effective",
			Value: "",
		},
	}
}

func BindBadgeConfig(c *cli.Context) *BadgeConfig {
	return &BadgeConfig{
		NoCommonBadges: c.Bool("no-common-badges"),
		GolangBadges:   c.Bool("golang"),
		RustBadges:     c.Bool("rust"),
		NodeBadges:     c.Bool("node"),
		RustVersion:    c.String("rust-version"),
		NpmPackage:     c.String("npm"),
		DockerUser:     c.String("docker-user"),
		DockerRepo:     c.String("docker-repo"),
	}
}
