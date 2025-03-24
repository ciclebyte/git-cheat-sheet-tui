package app

import (
	"database/sql"
	"fmt"

	"github.com/ciclebyte/git-cheat-sheet-tui/assets"
)

func (c *SheetView) InitDb() error {
	// 从嵌入的资源文件系统中读取数据库文件
	dbFile, err := assets.Resources.ReadFile("resources/commands.sql")
	if err != nil {
		return fmt.Errorf("读取数据库文件失败: %w", err)
	}

	// 创建内存数据库连接
	db, err := sql.Open("sqlite3", "file::memory:?cache=shared")
	if err != nil {
		return fmt.Errorf("连接数据库失败: %w", err)
	}

	// 直接执行SQL语句加载数据
	_, err = db.Exec(string(dbFile))
	if err != nil {
		return fmt.Errorf("加载数据库数据失败: %w", err)
	}

	c.db = db
	return nil
}
