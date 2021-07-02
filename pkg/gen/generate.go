package gen

import (
	"fmt"
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
	templateFilePath := os.Getenv("HUGO_TEMPLATE")
	if templateFilePath == "" {
		templateFilePath = "./templates/template.md"
	}
	templateGen, err := template.New("tmp").ParseFiles(templateFilePath)
	if err != nil {
		return err
	}
	f, err := os.OpenFile(g.FileName, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return err
	}

	return templateGen.ExecuteTemplate(f, "template.md", map[string]interface{}{
		"title":      g.Title,
		"author":     g.Author,
		"date":       time.Now().Format("2006-01-02T15:04:05+08:00"),
		"lastmod":    time.Now().Format("2006-01-02T15:04:05+08:00"),
		"tags":       genStringSlice(g.Tags),
		"categories": genStringSlice(g.Categories),
	})
}

// genStringSlice 将切片数组转换为字符串，类似["a", "b"]的形式
func genStringSlice(value []string) string {
	str := "["
	l := len(value)
	for i, v := range value {
		str += "\"" + v + "\""
		if i < l-1 {
			str += ", "
		}
	}
	str += "]"
	return str
}

// EnterMsg 获取标准输入中的值
func EnterMsg(v *[]string, tip string) (err error) {
	fmt.Println(tip)
	tips := make([]interface{}, 100)
	for i, _ := range tips {
		a := ""
		tips[i] = &a
	}
	_, _ = fmt.Scanln(tips...)
	for _, tip := range tips {
		if str, ok := tip.(*string); ok && *str != "" {
			*v = append(*v, *str)
		}
	}
	return
}
