package excample

// 链接pgsql

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

// TestPgsqlConnet 演示如何使用 pgx 库连接到 PostgreSQL 数据库，
// 执行查询以从 pg_stat_activity 视图中检索信息，并处理结果。
//
// 此函数执行以下步骤：
//  1. 使用连接 URL 建立到 PostgreSQL 数据库的连接。
//  2. 执行查询以从 pg_stat_activity 视图中检索活动查询的进程 ID (pid)、查询文本和状态，
//     排除空闲查询。
//  3. 遍历结果集并打印检索到的信息。
//
// 注意：
// - 确保在项目中正确导入并配置 pgx 库。
// - 根据您的环境替换连接 URL 中的凭据和数据库详细信息。
// - 安全地处理数据库凭据等敏感信息。
// - 如果任何步骤失败，函数将打印错误信息并退出程序。
func TestPgsqlConnet() {
	urlExample := "postgres://postgres:postgres@10.31.2.54:5432/postgres"
	conn, err := pgx.Connect(context.Background(), urlExample)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background()) // 关闭连接

	var pid int64
	var query string
	var state string
	querySql := "SELECT pid, query, $1 state FROM pg_stat_activity WHERE query <> '<IDLE>' ;"
	rows, err := conn.Query(context.Background(), querySql, "idle")
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&pid, &query, &state)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Scan failed: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("pid: %d, query: %s, state: %s\n", pid, query, state)
	}
}
