package main

import (
	"fmt"
	"github.com/sanguohot/medichain/util"
	"github.com/urfave/cli"
	"os"
	"time"
)

func InitApp() error {
	app := cli.NewApp()
	app.Name = "medichain"
	app.Usage = "command line for medichain!"
	app.Version = "1.0.1"
	app.Compiled = time.Now()
	app.Authors = []cli.Author{
		cli.Author{
			Name:  "Sanguohot",
			Email: "hw535431@163.com",
		},
	}
	app.Action = func(c *cli.Context) error {
		fmt.Printf("%s-%s", app.Name, app.Version)
		fmt.Printf("\n%s", app.Usage)
		return nil
	}
	app.Commands = nil

	return app.Run(os.Args)
}

func main() {
	err := InitApp()
	if err != nil {
		zap.Logger.Fatal(err.Error())
	}
}