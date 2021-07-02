package cmd

import (
	"fmt"
	"github.com/jalen-qian/GenHugoBlog/logging"
	"github.com/jalen-qian/GenHugoBlog/pkg/gen"
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
	"strings"
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
		&cli.StringSliceFlag{
			Name:  "tags",
			Usage: "标签",
			Value: nil,
		},
		&cli.StringSliceFlag{
			Name:  "cates",
			Usage: "分类",
			Value: nil,
		},
		&cli.StringFlag{
			Name:  "filename",
			Usage: "文件名称，默认用标题名称",
			Value: "",
		},
	},
	Usage: "生成器,用于根据模板生成文章",
	Action: func(c *cli.Context) (err error) {
		fmt.Println("请确保在环境变量[HUGO_TEMPLATE]配置模板文件的位置，如您未配置此环境变量，将从当前目录下的 [/templates/template.md] 读取，请确保此文件存在")
		title := c.String("title")
		if title == "" {
			titles := make([]string, 0)
			if err = gen.EnterMsg(&titles, "please enter your title:"); err != nil {
				logging.GLog.Error("scan title failed", zap.Error(err))
				return err
			}
			title = strings.Join(titles, " ")
		}
		author := c.String("author")
		if author == "" {
			authors := make([]string, 0)
			if err = gen.EnterMsg(&authors, "please enter your author:"); err != nil {
				logging.GLog.Error("scan title failed", zap.Error(err))
				return err
			}
			author = strings.Join(authors, " ")
		}
		tags := c.StringSlice("tags")
		if len(tags) == 0 {
			if err = gen.EnterMsg(&tags, "please enter tags:"); err != nil {
				logging.GLog.Error("scan tags failed", zap.Error(err))
				return err
			}
		}
		cates := c.StringSlice("cates")
		if len(cates) == 0 {
			if err = gen.EnterMsg(&cates, "please enter categories:"); err != nil {
				logging.GLog.Error("scan categories failed", zap.Error(err))
				return err
			}
		}
		filename := c.String("filename")
		if filename == "" {
			filename = fmt.Sprintf("./%s.md", title)
		}
		generator := gen.Generator{
			FileName:   filename,
			Title:      title,
			Author:     author,
			Tags:       tags,
			Categories: tags,
		}
		err = generator.GenerateFile()
		if err != nil {
			logging.GLog.Error("gen failed", zap.Error(err))
			return err
		}
		return err
	},
}
