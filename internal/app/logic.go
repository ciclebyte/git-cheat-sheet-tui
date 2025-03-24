package app

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

// 定义一个名为 SetToDefaultBorderColor 的方法，该方法属于 SheetView 类型
func (c *SheetView) SetColor(p tview.Primitive, color tcell.Color) {
	switch v := p.(type) {
	case *tview.List:
		v.SetBorderColor(color)
	case *tview.TextView:
		v.SetBorderColor(color)
	case *tview.InputField:
		v.SetBorderColor(color)
	}
}

func (c *SheetView) SetToDefaultBorderColor() {
	for _, comp := range c.focusable {
		c.SetColor(comp, tcell.ColorWhite)
	}
}

func (c *SheetView) Focus(p tview.Primitive) {
	c.app.SetFocus(p)
	c.SetColor(p, c.hightLightColor)
}

func (c *SheetView) UpdateDefaultBorderColor() {
	// Reset border color of current focused component
	if currentFocus := c.app.GetFocus(); currentFocus != nil {
		c.SetColor(currentFocus, tcell.ColorWhite)
	}
}

func (c *SheetView) LeftFocus() {
	c.UpdateDefaultBorderColor()
	c.focusIndex = (c.GetCurrentFocusIndex() - 1 + len(c.focusable)) % len(c.focusable)
	c.app.SetFocus(c.focusable[c.focusIndex])
	c.SetColor(c.focusable[c.focusIndex], c.hightLightColor)
}

func (c *SheetView) RightFocus() {
	c.UpdateDefaultBorderColor()
	c.focusIndex = (c.GetCurrentFocusIndex() + 1) % len(c.focusable)
	c.app.SetFocus(c.focusable[c.focusIndex])
	c.SetColor(c.focusable[c.focusIndex], c.hightLightColor)
}

func (c *SheetView) LogInfo(msg string) {
	if !c.showLog {
		return
	}
	fmt.Fprintf(c.logView, "[red][%s] INFO:[-] %s\n", time.Now().Format("15:04:05"), msg)
	c.logView.ScrollToEnd()
}

func (c *SheetView) LogError(msg string) {
	if !c.showLog {
		return
	}
	fmt.Fprintf(c.logView, "[red][%s] ERROR:[-] %s\n", time.Now().Format("15:04:05"), msg)
	c.logView.ScrollToEnd()
}

func (c *SheetView) BindFunc() {
	c.categoryList.SetChangedFunc(func(i int, s1, s2 string, r rune) {
		if !c.isSearching {
			c.RenderComandListView(s1)
		}
	})
	c.commanList.SetChangedFunc(func(i int, s1, s2 string, r rune) {
		c.RenderUsageView(s1, s2)
	})
	// 键盘事件处理
	c.categoryList.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyLeft, tcell.KeyRight:
			// 阻止 ListView 默认的展开/折叠行为
			return nil
		}
		return event
	})
	// 键盘事件处理
	c.commanList.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyLeft, tcell.KeyRight:
			// 阻止 ListView 默认的展开/折叠行为
			return nil
		}
		return event
	})
	// 添加键盘事件处理
	c.searchBox.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyLeft:
			c.focusIndex = 0
			c.app.SetFocus(c.categoryList)
			return nil
		case tcell.KeyRight:
			c.focusIndex = 0
			c.app.SetFocus(c.categoryList)
			return nil
		}
		return event
	})
	c.searchBox.SetDoneFunc(func(key tcell.Key) {
		if key == tcell.KeyEnter {
			c.Search(c.searchBox.GetText())
		} else if key == tcell.KeyF2 {
			c.ResetSearch()
			c.searchBox.SetText("")
		}
	})
	// 设置键盘输入
	c.main.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch {
		case event.Key() == tcell.KeyEsc || event.Rune() == 'q':
			c.app.Stop()
		case event.Key() == tcell.KeyLeft:
			c.LeftFocus()
		case event.Key() == tcell.KeyRight:
			c.RightFocus()
		case event.Key() == tcell.KeyCtrlR:
			c.ResetSearch()
		case event.Key() == tcell.KeyCtrlF:
			c.SetToDefaultBorderColor()
			c.app.SetFocus(c.searchBox)
		}
		return event
	})
}

func (c *SheetView) Search(key string) {
	c.isSearching = true
	c.categoryList.Clear()
	c.commanList.Clear()

	// 从数据库搜索
	rows, err := c.db.Query(`
        SELECT title, command 
        FROM commands 
        WHERE title LIKE ? 
           OR command LIKE ? 
           OR description LIKE ?`,
		"%"+key+"%", "%"+key+"%", "%"+key+"%")
	if err != nil {
		c.LogError("搜索失败: " + err.Error())
		return
	}
	defer rows.Close()

	count := 0
	for rows.Next() {
		var title, command string
		if err := rows.Scan(&title, &command); err != nil {
			c.LogError("查询数据失败: " + err.Error())
			continue
		}
		c.commanList.AddItem(title, command, 0, func() {})
		count++
	}

	c.LogInfo(fmt.Sprintf("找到 %d 条结果", count))
	c.categoryList.AddItem("搜索结果("+strconv.Itoa(count)+")条", "", 0, func() {})

	// 切换到命令列表
	c.Focus(c.commanList)
	// 选择第一个命令
	if c.commanList.GetItemCount() > 0 {
		c.commanList.SetCurrentItem(0)
	} else {
		c.LogInfo("未找到匹配的命令") // 添加日志记录空结果
	}
}

func (c *SheetView) ResetSearch() {
	// 重新渲染分类列表
	c.isSearching = false
	// 重置搜索状态
	c.searchBox.SetText("")
	// 清空搜索结果
	c.categoryList.Clear()
	c.commanList.Clear()
	// 设置为默认边框颜色
	c.SetToDefaultBorderColor()
	// 切换到分类列表
	c.Focus(c.categoryList)
	c.RenderCategoryListView()
	// 选择第一个分类
	c.categoryList.SetCurrentItem(0)
	index := c.categoryList.GetCurrentItem()
	text, _ := c.categoryList.GetItemText(index)
	c.RenderComandListView(text)
}
