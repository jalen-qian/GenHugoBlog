package gen

import (
	"os"
	"text/template"
	"time"
)

/**
文章生成器
*/

type Generator struct {
	FileName string
	// Title 标题
	Title string `template:"title"`
	// Date 文章发布时间
	Date time.Time `template:"date"`
	// LastMod 上次修改时间
	LastMod time.Time `template:"lastmod"`
	// Tags 标签
	Tags []string `template:"tags"`
	// Categories 分类
	Categories []string `template:"categories"`
	// Author 作者
	Author string `template:"author"`
}

// GenerateFile 初始化模板文件
func (g *Generator) GenerateFile() error {
	templateGen, err := template.New("tmp").ParseFiles("./templates/template.md")
	if err != nil {
		return err
	}
	f, err := os.OpenFile(g.FileName, os.O_RDWR|os.O_APPEND, os.ModePerm)
	if err != nil {
		return err
	}
	return templateGen.Execute(f, map[string]interface{}{
		"title" : g.Title,
		"author" : g.Author,
		"date" : time.Now().String(),
		"lastmod": time.Now().String(),
		"tags" : `["a","b"]`,
		"categories" : `["a","b"]`,
	})
}
