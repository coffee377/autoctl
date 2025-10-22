package test

import (
	"cds/bid/ds"
	"context"
	"log"
	"testing"

	"entgo.io/ent/dialect/sql/schema"

	"github.com/stretchr/testify/assert"
)

func Test_Migration(t *testing.T) {
	client, ok := ds.Mysql()
	defer ds.CloseMysql(client)
	assert.Equal(t, true, ok)

	ctx := context.Background()

	// 运行自动迁移工具来创建所有Schema资源
	if err := client.Debug().Schema.Create(ctx,
		schema.WithDropIndex(true),
		schema.WithDropColumn(true),
	); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

}
