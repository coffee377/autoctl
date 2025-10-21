package main

import (
	"log"

	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"github.com/coffee377/entcc"
	"go.uber.org/zap"
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
	}, opts...); err != nil {
		log.Fatal("running ent codegen:", err)
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
