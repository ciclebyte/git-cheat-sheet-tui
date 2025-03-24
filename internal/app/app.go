package app

import (
	"database/sql"
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/mattn/go-colorable"
	_ "github.com/mattn/go-sqlite3"
	"github.com/rivo/tview"
)

type SheetView struct {
	app             *tview.Application
	categoryList    *tview.List
	commanList      *tview.List
	usageView       *tview.TextView
	logView         *tview.TextView
	searchBox       *tview.InputField
	main            *tview.Flex
	modal           *tview.Flex
	pages           *tview.Pages
	focusable       []tview.Primitive
	showLog         bool
	focusIndex      int
	enableHighlight bool
	isSearching     bool // 是否正在搜索
	db              *sql.DB
	hightLightColor tcell.Color
}

func (c *SheetView) Run(searchKeyword string, showLog bool) {
	c.showLog = showLog
	// 设置高亮颜色
	c.hightLightColor = tcell.ColorRed
	// 启用Windows颜色支持
	defer colorable.EnableColorsStdout(nil)()
	// 启用高亮
	c.enableHighlight = true
	c.app = tview.NewApplication()
	// enable鼠标
	c.app.EnableMouse(false)

	// 初始化数据库
	if err := c.InitDb(); err != nil {
		fmt.Println("初始化缓存失败: " + err.Error())
		return
	}
	// 只有在数据库成功初始化后才设置defer
	defer func() {
		if c.db != nil {
			c.db.Close()
		}
	}()

	// 创建mainFlex
	c.main = c.CreateMainFlex()
	c.BindFunc()
	// 创建模态容器
	c.modal = c.CreateModal()
	// 使用 Pages 管理界面状态
	c.pages = tview.NewPages().
		AddPage(MAIN_PAGE, c.main, true, true).
		AddPage(MODAL_PAGE, c.modal, true, false)
	// 初始化可聚焦组件列表
	c.focusable = []tview.Primitive{c.categoryList, c.commanList, c.usageView}
	c.focusIndex = 0
	// 如果搜索词不为空，直接加载搜索结果
	if searchKeyword != "" {
		c.searchBox.SetText(searchKeyword)
		c.Search(searchKeyword)
		c.focusIndex = 1
	} else {
		// 初始化mainFlex中categoryList数据
		c.RenderCategoryListView()
		// 默认选择第一个category
		c.categoryList.SetCurrentItem(0)
	}

	// 设置根节点
	c.app.SetRoot(c.pages, true)
	// 设置焦点
	c.Focus(c.focusable[c.focusIndex])
	// 设置初始焦点组件的边框颜色
	if focusable, ok := c.focusable[c.focusIndex].(interface{ SetBorderColor(tcell.Color) }); ok {
		focusable.SetBorderColor(tcell.ColorSkyblue)
	}
	if err := c.app.Run(); err != nil {
		fmt.Println("TUI启动失败: " + err.Error())
	}
}
