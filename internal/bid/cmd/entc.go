package main

import (
	"log"

	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"github.com/coffee377/entcc"
)

func main() {
	extension, err := entcc.NewExtension()
	if err != nil {
		log.Fatal(err)
	}
	opts := []entc.Option{
		//entc.Dependency(
		//	entc.DependencyType(&zap.Logger{}),
		//),
		entc.Extensions(extension),
	}
	if err := entc.Generate("./ent/schema", &gen.Config{}, opts...); err != nil {
		log.Fatal("running ent codegen:", err)
	}
}
