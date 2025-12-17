package ds

import (
	"cds/bid/ent"
	"context"
	"fmt"
	"log"
	"net/url"
	"sync"

	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
)

// Dialect names for external usage.
const (
	MySQL     = "mysql"
	SQLServer = "mssql"
	SQLite    = "sqlite3"
	Postgres  = "postgres"
	Gremlin   = "gremlin"
)

type DataSource struct {
	typ      string
	username string
	password string
	host     string
	port     string
	database string
}

func (ds DataSource) DSN() string {
	var dsn string
	switch ds.typ {
	case MySQL:
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s", ds.username, ds.password,
			ds.host, ds.port, ds.database, "charset=utf8mb4&parseTime=True&loc=Asia%2FShanghai")
	case SQLServer:
		dsn = fmt.Sprintf(
			"sqlserver://%s:%s@%s:%s?database=%s&encrypt=disable&trustservercertificate=true",
			ds.username, url.QueryEscape(ds.password), ds.host, ds.port, ds.database,
		)
	default:
		panic("unsupported database type")
	}

	return dsn
}

type DS struct {
	mysql     DataSource
	sqlserver DataSource
}

var (
	dsDev = DS{
		mysql: DataSource{
			typ:      MySQL,
			username: "root",
			password: "root!@@&",
			host:     "localhost",
			port:     "3306",
			database: "cds_infra",
		},
		sqlserver: DataSource{
			typ:      SQLServer,
			username: "sa",
			password: "sa123$%^",
			host:     "192.168.45.150",
			port:     "1433",
			database: "JQMSDB_main_220726",
		},
	}
	dsTest = DS{
		mysql: DataSource{
			typ:      MySQL,
			username: "teamwork",
			password: "jqkj5350**)",
			host:     "10.1.83.26",
			port:     "3306",
			database: "teamwork",
		},
		sqlserver: DataSource{
			typ:      SQLServer,
			username: "sa",
			password: "sa123$%^",
			host:     "192.168.45.150",
			port:     "1433",
			database: "JQMSDB_main_220726",
		},
	}
	dsProd = DS{
		mysql: DataSource{
			username: "teamwork",
			password: "teamwork_jqkj5350**)123",
			host:     "192.168.44.82",
			port:     "3306",
			database: "teamwork",
		},
		sqlserver: DataSource{
			username: "buzhidaomingzi",
			password: "154sjk08123$%^",
			host:     "192.168.44.154",
			port:     "1433",
			database: "JQMSDB_main",
		},
	}
)

// 全局缓存：Ent Client
var (
	clientCache = make(map[string]*ent.Client)
	cacheMu     sync.RWMutex // 并发安全锁
)

func init() {
	err := initClients()
	if err != nil {
		panic(err)
	}
}

// 初始化所有已知数据源的 Client
func initClients() error {
	cacheMu.Lock()
	defer cacheMu.Unlock()

	mysql := dsTest.mysql
	// 创建 Ent Client
	client, err := ent.Open(mysql.typ, mysql.DSN())
	if err != nil {
		return err
	}
	// 缓存 Client
	clientCache[MySQL] = client

	//sqlserver := dsDev.sqlserver
	//client, err = ent.Open(sqlserver.typ, sqlserver.DSN())
	//if err != nil {
	//	return err
	//}
	//clientCache[SQLServer] = client

	return nil
}

func Mysql() (*ent.Client, bool) {
	return getClient(MySQL)
}

func CloseMysql(client *ent.Client) {
	_ = client.Close()
}

func SqlServer() (*ent.Client, bool) {
	return getClient(SQLServer)
}

func getClient(driver string) (*ent.Client, bool) {
	cacheMu.RLock()
	defer cacheMu.RUnlock()
	client, ok := clientCache[driver]
	return client, ok
}

func WithEntTx(ctx context.Context, client *ent.Client, fn func(tx *ent.Tx) error) error {
	tx, err := client.Tx(ctx)
	if err != nil {
		return err
	}
	defer func() {
		if r := recover(); r != nil {
			log.Print(r)
			err := tx.Rollback()
			if err != nil {
				return
			}
		}
	}()

	if err = fn(tx); err != nil {
		if rErr := tx.Rollback(); rErr != nil {
			err = fmt.Errorf("%w: rolling back transaction: %v", err, rErr)
		}
		return err
	}
	if err = tx.Commit(); err != nil {
		return fmt.Errorf("committing transaction: %w", err)
	}
	return nil
}
