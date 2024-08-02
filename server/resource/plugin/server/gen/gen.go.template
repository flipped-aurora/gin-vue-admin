package main

import (
	"gorm.io/gen"
	"path/filepath"
)

//go:generate go mod tidy
//go:generate go mod download
//go:generate go run gen.go
func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath: filepath.Join("..", "..", "..", "{{ .Package }}", "blender", "model", "dao"),
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
	})
	g.ApplyBasic()
	g.Execute()
}
