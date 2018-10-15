package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "recnf"
	app.Email = "https://github.com/rinetd"
	app.Usage = `Translate YAML, JSON, TOML, ...
	 jy -t toml file.json
	`
	app.Version = "1.0.0"
	app.Action = Action

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "i, input",
			Value: "json",
			Usage: "input Type: json, yaml, toml, xml",
		},
		cli.StringFlag{
			Name:  "s, src",
			Value: "json",
			Usage: "input Type: ./conf.json, conf.yaml, toml, xml",
		},
		cli.StringFlag{
			Name:  "o, output",
			Value: "yaml",
			Usage: "output  file : json, yaml, toml, xml",
		},
		cli.StringFlag{
			Name:  "t, type",
			Value: "yaml",
			Usage: "output Type: json, yaml, toml, xml",
		},
	}

	app.Run(os.Args)
}

// 说明: cli 中的 args 是不包含参数的！！使用的前提是选项必须放在开头 超级棒。

func Action(c *cli.Context) {
	fmt.Println(c.NArg())
	// ./recnf -i json /data/src.go dst.go
	if c.NArg() > 0 {
		fmt.Println(c.Args()[c.NArg()-1])
		for index, value := range c.Args() {
			fmt.Println(index, value)
		}
	}

}
