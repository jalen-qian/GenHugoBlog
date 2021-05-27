package cmd

import (
	"github.com/jalen-qian/GenHugoBlog/logging"
	"github.com/jalen-qian/GenHugoBlog/pkg/gen"
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
)

var genCmd = &cli.Command{
	Name: "gen",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "title",
			Usage: "文章的标题",
			Value: "",
		},
		&cli.StringFlag{
			Name:  "author",
			Usage: "文章的作者",
			Value: "Jalen",
		},
	},
	Usage: "生成器,用于根据模板生成文章",
	Action: func(c *cli.Context) error {
		generator := gen.Generator{
			FileName: "./aaa.md",
			Title:    c.String("title"),
			Author:   c.String("author"),
		}
		err := generator.GenerateFile()
		if err != nil {
			logging.GLog.Error("gen failed", zap.Error(err))
		}
		return err
	},
}
