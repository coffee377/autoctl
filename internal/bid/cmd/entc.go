package main

import (
	"cds/bid/ent/migrate"
	"context"
	"log"

	atlas "ariga.io/atlas/sql/migrate"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"github.com/coffee377/entcc"
	"go.uber.org/zap"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	development, err := zap.NewDevelopment()
	if err != nil {
		log.Fatal(err)
	}

	extension, err := entcc.NewExtension(entcc.WithZapLogger(development))
	if err != nil {
		log.Fatal(err)
	}
	opts := []entc.Option{
		//entc.Dependency(
		//	entc.DependencyType(&zap.Logger{}),
		//),
		entc.Extensions(extension),
	}
	if err := entc.Generate("./ent/schema", &gen.Config{
		Hooks: []gen.Hook{
			testHook(development.Sugar()),
		},
		Features: []gen.Feature{gen.FeatureVersionedMigration},
	}, opts...); err != nil {
		log.Fatal("running ent codegen:", err)
	}
	//genMigrate()
}

func genMigrate() {
	ctx := context.Background()
	// Create a local migration directory able to understand Atlas migration file format for replay.
	dir, err := atlas.NewLocalDir("migrations")
	if err != nil {
		log.Fatalf("failed creating atlas migration directory: %v", err)
	}
	// Migrate diff options.
	opts := []schema.MigrateOption{
		schema.WithDir(dir), // provide migration directory
		//schema.WithMigrationMode(schema.ModeReplay), // provide migration mode
		schema.WithDialect(dialect.MySQL), // Ent dialect to use
		schema.WithFormatter(atlas.DefaultFormatter),
	}
	//if len(os.Args) != 2 {
	//	log.Fatalln("migration name is required. Use: 'go run -mod=mod ent/migrate/main.go <name>'")
	//}
	// Generate migrations using Atlas support for MySQL (note the Ent dialect option passed above).
	err = migrate.NamedDiff(ctx, "mysql://dev:dev@localhost:3306/ccl_base", "test", opts...)
	if err != nil {
		log.Fatalf("failed generating migration file: %v", err)
	}
}
func testHook(logger *zap.SugaredLogger) gen.Hook {
	return func(next gen.Generator) gen.Generator {
		return gen.GenerateFunc(func(g *gen.Graph) error {
			for _, node := range g.Nodes {
				logger.Debugf("=> %s", node.Name)
			}
			return next.Generate(g)
		})
	}
}
