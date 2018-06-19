package main

import (
	"fmt"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/urfave/cli"

	"regexp"
)

const (
	regexp_phone = "^((13[0-9])|(14[5|7])|(15([0-3]|[5-9]))|(18[0,5-9]))\\d{8}$"
)

func validate(mobileNum string) bool {
	r := regexp.MustCompile(regexp_phone)
	return r.MatchString(mobileNum)
}

var p = "15966638724"

type Config struct {
	Hosts     []string `json:"hosts"`
	Port      int      `json:"port"`
	User      string   `json:"user"`
	Key       string   `json:"ssh-key"`
	Password  string   `json:"password"`
	Source    string   `json:"source"`
	Target    string   `json:"target"`
	Delete    bool     `json:"delete"`
	Chmod     string   `json:"chmod"`
	Chown     string   `json:"chown"`
	Recursive bool     `json:"recursive"`
	Include   []string `json:"include"`
	Exclude   []string `json:"exclude"`
	Filter    []string `json:"filter"`
	Script    []string `json:"script"`
}

// Plugin in Rsync
type Plugin struct {
	Config Config
}

// build number set at compile-time
var build = "0"

// Version set at compile-time
var Version string

func Main_slice() {
	if Version == "" {
		Version = fmt.Sprintf("1.3.1+%s", build)
	}

	app := cli.NewApp()
	app.Name = "Drone-rsync"
	app.Usage = "Rsync to Remote Hosts"
	app.Copyright = "Copyright (c) 2017 rinetd"
	app.Authors = []cli.Author{
		{
			Name:  "rinetd",
			Email: "sdlylshl@gmail.com",
		},
	}
	app.Action = run
	app.Version = Version
	app.Flags = []cli.Flag{

		// 		User:      c.String("user"),
		// 		Key:       c.String("ssh-key"),
		// 		Password:  c.String("password"),
		// 		Source:    c.String("source"),
		// 		Target:    c.String("target"),
		// 		Recursive: c.Bool("recursive"),
		// 		Delete:    c.Bool("delete"),
		// 		Chmod:     c.String("chmod"),
		// 		Chown:     c.String("chown"),
		// 		Include:   c.StringSlice("include"),
		// 		Exclude:   c.StringSlice("exclude"),
		// 		Filter:    c.StringSlice("filter"),
		// 		Script:    c.StringSlice("script"),
		cli.StringSliceFlag{
			Name:   "hosts,H",
			Usage:  "connect to host",
			EnvVar: "PLUGIN_HOSTS,SSH_HOST",
		},
		cli.IntFlag{
			Name:   "port,p",
			Usage:  "connect to port ",
			EnvVar: "PLUGIN_PORT,SSH_PORT",
			Value:  22,
		},
		cli.StringFlag{
			Name:   "username,user,u",
			Usage:  "connect as user ",
			EnvVar: "PLUGIN_USERNAME,PLUGIN_USER,SSH_USERNAME",
			Value:  "root",
		},
		cli.StringFlag{
			Name:   "ssh-key,key",
			Usage:  "private ssh key",
			EnvVar: "PLUGIN_SSH_KEY,PLUGIN_KEY,SSH_KEY",
		},
		cli.StringFlag{
			Name:   "password,P",
			Usage:  "user password",
			EnvVar: "PLUGIN_PASSWORD,SSH_PASSWORD",
		},
		cli.StringFlag{
			Name:   "source",
			Usage:  "source",
			EnvVar: "PLUGIN_SOURCE",
		},
		cli.StringFlag{
			Name:   "target",
			Usage:  "target",
			EnvVar: "PLUGIN_TARGET",
		},

		cli.BoolFlag{
			Name:   "recursive",
			Usage:  "recursive mode",
			EnvVar: "PLUGIN_RECURSIVE",
		},
		cli.BoolFlag{
			Name:   "delete",
			Usage:  "delete mode",
			EnvVar: "PLUGIN_DELETE",
		},
		cli.StringFlag{
			Name:   "chmod",
			Usage:  "chmod commands",
			EnvVar: "PLUGIN_CHMOD",
		},
		cli.StringFlag{
			Name:   "chown",
			Usage:  "chown commands",
			EnvVar: "PLUGIN_CHOWN",
		},
		cli.StringSliceFlag{
			Name:   "include",
			Usage:  "include commands",
			EnvVar: "PLUGIN_INCLUDE",
		},
		cli.StringSliceFlag{
			Name:   "exclude",
			Usage:  "exclude commands",
			EnvVar: "PLUGIN_EXCLUDE",
		},
		cli.StringSliceFlag{
			Name:   "filter",
			Usage:  "filter commands",
			EnvVar: "PLUGIN_FILTER",
		},
		cli.StringSliceFlag{
			Name:   "script,s",
			Usage:  "execute commands",
			EnvVar: "PLUGIN_SCRIPT,SSH_SCRIPT",
		},
		// cli.BoolFlag{
		// 	Name:   "sync",
		// 	Usage:  "sync mode",
		// 	EnvVar: "PLUGIN_SYNC",
		// },
		// cli.DurationFlag{
		// 	Name:   "timeout,t",
		// 	Usage:  "connection timeout",
		// 	EnvVar: "PLUGIN_TIMEOUT,SSH_TIMEOUT",
		// },
		// cli.IntFlag{
		// 	Name:   "command.timeout,T",
		// 	Usage:  "command timeout",
		// 	EnvVar: "PLUGIN_COMMAND_TIMEOUT,SSH_COMMAND_TIMEOUT",
		// 	Value:  60,
		// },

		// cli.StringFlag{
		// 	Name:  "env-file",
		// 	Usage: "source env file",
		// },

		// cli.StringSliceFlag{
		// 	Name:   "secrets",
		// 	Usage:  "plugin secret",
		// 	EnvVar: "PLUGIN_SECRETS",
		// },
		// cli.StringSliceFlag{
		// 	Name:   "envs",
		// 	Usage:  "Pass envs",
		// 	EnvVar: "PLUGIN_ENVS",
		// },
		// cli.BoolFlag{
		// 	Name:   "debug",
		// 	Usage:  "debug mode",
		// 	EnvVar: "PLUGIN_DEBUG",
		// },
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Println("drone-ssh error: ", err)
		os.Exit(1)
	}
}

func run(c *cli.Context) error {
	// if c.String("env-file") != "" {
	// 	_ = godotenv.Load(c.String("env-file"))
	// }

	plugin := Plugin{
		Config: Config{
			Hosts:     c.StringSlice("hosts"),
			Port:      c.Int("port"),
			User:      c.String("user"),
			Key:       c.String("ssh-key"),
			Password:  c.String("password"),
			Source:    c.String("source"),
			Target:    c.String("target"),
			Recursive: c.Bool("recursive"),
			Delete:    c.Bool("delete"),
			Chmod:     c.String("chmod"),
			Chown:     c.String("chown"),
			Include:   c.StringSlice("include"),
			Exclude:   c.StringSlice("exclude"),
			Filter:    c.StringSlice("filter"),
			Script:    c.StringSlice("script"),
		},
	}
	fmt.Println("pl%vgin", plugin)

	return plugin.Exec()
	// return nil
}

func (p Plugin) Exec() error {
	return nil
}
