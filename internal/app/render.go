package app

import (
	"github.com/ciclebyte/git-cheat-sheet-tui/internal/utils"
	"github.com/rivo/tview"
)

var markdownRenderCache = make(map[string]string)

type Meta struct {
	Title       string `yaml:"title"`
	Command     string `yaml:"command"`
	Description string `yaml:"description"`
}

func (c *SheetView) RenderCategoryListView() {
	// 从数据库查询分类
	rows, err := c.db.Query("SELECT DISTINCT category FROM commands")
	if err != nil {
		c.LogError("查询分类失败: " + err.Error())
		return
	}
	defer rows.Close()

	for rows.Next() {
		var category string
		if err := rows.Scan(&category); err != nil {
			continue
		}
		c.categoryList.AddItem(category, "", 0, func() {
			c.categoryList.SetCurrentItem(c.categoryList.GetCurrentItem())
		})
	}

}

func (c *SheetView) RenderComandListView(categoryName string) {
	c.commanList.Clear()

	// 从数据库查询命令
	rows, err := c.db.Query("SELECT title, command FROM commands WHERE category = ?", categoryName)
	if err != nil {
		c.LogError("查询命令失败: " + err.Error())
		return
	}
	defer rows.Close()

	for rows.Next() {
		var title, command string
		if err := rows.Scan(&title, &command); err != nil {
			continue
		}
		c.commanList.AddItem(title, command, 0, func() {
			c.commanList.SetCurrentItem(c.categoryList.GetCurrentItem())
		})
	}

}

func (c *SheetView) RenderUsageView(commandName string, command string) {
	// 从数据库查询内容
	var content string
	c.LogInfo(" title:" + commandName)
	err := c.db.QueryRow("SELECT content FROM commands WHERE title = ? and command = ? ",
		commandName, command).Scan(&content)
	if err != nil {
		c.LogError("查询内容失败: " + err.Error())
		return
	}
	c.usageView.Clear()

	// 直接使用数据库中的内容
	renderedContent, err := utils.MarkdownInstance.MarkdownRender(content)
	if err != nil {
		c.usageView.SetText(content)
	} else {
		if cachedContent, ok := markdownRenderCache[commandName]; ok {
			c.LogInfo("cached render")
			c.usageView.SetText(cachedContent)
		} else {
			c.LogInfo("fisrt render")
			ansiContent := tview.TranslateANSI(renderedContent)
			// 保存
			markdownRenderCache[commandName] = ansiContent
			c.usageView.SetText(ansiContent)
		}
	}
}
