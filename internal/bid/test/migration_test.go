package test

import (
	"cds/bid/ent"
	"context"
	"fmt"
	"log"
	"testing"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql/schema"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
)

func Test_Migration(t *testing.T) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s", "root", "root!@@&", "localhost", "3306", "cds_infra",
		"charset=utf8mb4&parseTime=True&loc=Asia%2FShanghai")
	client, err := ent.Open(dialect.MySQL, dsn)
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}
	defer func(client *ent.Client) {
		err := client.Close()
		if err != nil {
			log.Fatalf("failed closing connection to mysql: %v", err)
		}
	}(client)

	ctx := context.Background()

	// 运行自动迁移工具来创建所有Schema资源
	if err := client.Debug().Schema.Create(ctx,
		schema.WithDropIndex(true),
		schema.WithDropColumn(true),
	); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

}
