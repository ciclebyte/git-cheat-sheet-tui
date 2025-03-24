package app

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func (c *SheetView) CreateMainFlex() *tview.Flex {
	mainFlex := tview.NewFlex()
	mainFlex.SetDirection(tview.FlexRow)
	// Header
	mainFlex.AddItem(c.CreateHeader(), 3, 0, false)
	// 搜索框
	mainFlex.AddItem(c.CreateSearchBox(), 3, 0, false)
	// CategoryList
	mainFlex.AddItem(c.CreateMain(), 0, 1, false)
	if c.showLog {
		mainFlex.AddItem(c.CreateLogs(), 8, 10, false)
	}
	return mainFlex
}

func (c *SheetView) CreateModal() *tview.Flex {
	//TODO 添加模态框,暂时未实现,方便后续扩展
	modal := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(tview.NewTextView().SetText("请选择一个操作:"), 1, 1, false)
	// 创建下拉列表
	list := tview.NewList().
		AddItem("收藏", "", 0, func() {

		})
	// 不显示第二行文本,否则会有空行
	list.ShowSecondaryText(false)
	// 创建关闭按钮
	btnClose := tview.NewButton("关闭").SetSelectedFunc(func() {

	})

	// 组合模态内容
	modal.
		AddItem(list, 0, 4, true).
		AddItem(btnClose, 1, 1, true)

	// 创建带背景遮罩的模态层
	modalWrapper := tview.NewFlex().
		AddItem(nil, 0, 1, false).
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(nil, 0, 1, false).
			AddItem(modal, 20, 1, true).             // 模态框高度
			AddItem(nil, 0, 1, false), 40, 1, true). // 模态框宽度
		AddItem(nil, 0, 1, false)
	return modalWrapper
}

func (c *SheetView) CreateMain() *tview.Flex {
	// 仓库树
	c.CreateCategoryList()
	// 文件浏览器
	c.CreateCommandList()
	// 仓库描述
	c.CreateUsageView()

	// 创建布局
	mainFlex := tview.NewFlex()
	mainFlex.AddItem(c.categoryList, 35, 0, true)
	mainFlex.AddItem(tview.NewFlex().
		AddItem(c.commanList, 0, 1, true).
		AddItem(c.usageView, 0, 4, false), 0, 1, false)
	return mainFlex
}

func (c *SheetView) CreateHeader() *tview.TextView {
	return tview.NewTextView().
		SetDynamicColors(true).
		SetTextAlign(tview.AlignCenter).
		SetText("[::b]Git Cheat Sheet[::-] [darkcyan] ← →:切换 ↑↓:导航 CTRL+F:搜索 /CTRL+R:重置 Q:退出")
}

func (c *SheetView) CreateCategoryList() *tview.List {
	c.categoryList = tview.NewList().ShowSecondaryText(false)
	c.categoryList.SetBorder(true).SetTitle(" 分类 ")
	c.categoryList.SetBorderColor(tcell.ColorWhite) // 添加默认边框颜色
	return c.categoryList
}

func (c *SheetView) CreateCommandList() *tview.List {
	c.commanList = tview.NewList().ShowSecondaryText(true)
	c.commanList.SetBorder(true).SetTitle(" 命令 ")
	c.commanList.SetBorderColor(tcell.ColorWhite) // 添加默认边框颜色
	return c.commanList
}

func (c *SheetView) CreateUsageView() *tview.TextView {
	c.usageView = tview.NewTextView().
		SetDynamicColors(true).
		SetScrollable(true).
		SetWrap(true)
	c.usageView.SetBorder(true).SetTitle(" 命令使用 ")
	c.usageView.SetBorderColor(tcell.ColorWhite) // 添加默认边框颜色
	return c.usageView
}

func (c *SheetView) CreateLogs() *tview.TextView {
	c.logView = tview.NewTextView().
		SetDynamicColors(true).
		SetScrollable(true).
		SetWrap(false)
	c.logView.SetBorder(true).SetTitle(" 日志 ")
	return c.logView
}

func (c *SheetView) GetCurrentFocusIndex() int {
	// 获取当前获得焦点的组件
	currentFocus := c.app.GetFocus()
	// 查找当前焦点组件在 focusable 列表中的索引
	for i, comp := range c.focusable {
		if comp == currentFocus {
			c.focusIndex = i
			break
		}
	}
	return c.focusIndex
}

func (c *SheetView) CreateSearchBox() *tview.InputField {
	c.searchBox = tview.NewInputField()
	c.searchBox.SetLabel(" 搜索: ")
	c.searchBox.SetFieldBackgroundColor(tcell.ColorDefault)
	c.searchBox.SetFieldTextColor(tcell.ColorWhite)
	c.searchBox.SetLabelColor(tcell.ColorYellow)
	c.searchBox.SetBorder(true)
	c.searchBox.SetBorderColor(tcell.ColorDarkCyan)
	c.searchBox.SetPlaceholder("输入命令关键词...")
	c.searchBox.SetPlaceholderTextColor(tcell.ColorDarkGray)
	return c.searchBox
}
