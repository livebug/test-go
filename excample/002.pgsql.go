package excample

// 链接pgsql

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

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
