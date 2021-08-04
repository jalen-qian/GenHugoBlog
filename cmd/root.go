package cmd

import (
	"sort"

	"github.com/jalen-qian/GenHugoBlog/logging"

	"github.com/urfave/cli/v2"
)

func NewApp() *cli.App {
	app := &cli.App{
		Name:        "GenHugoBlog",
		HelpName:    "GenHugoBlog",
		Usage:       "GenHugoBlog 用于帮助根据模板生成Hugo博客网站的文章",
		Description: "运行在k8s集群,提供对集群的管理",
		Flags:       []cli.Flag{},
		Commands: []*cli.Command{
			genCmd, envCmd,
		},
		Before: func(c *cli.Context) error {
			var err error
			logging.GLog, err = logging.NewLogger(c.String("log-level"))
			return err
		},
	}

	app.Version = "v1.0.0"

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	return app
}
