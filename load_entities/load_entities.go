package loadentities

import (
	"gorm.io/gen"
	"gorm.io/gorm"
)

func MigrationExistData(db *gorm.DB) {

	g := gen.NewGenerator(gen.Config{
		OutPath: "./query",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
	})

	g.UseDB(db)
	g.ApplyBasic(
		g.GenerateAllTable()...,
	)
	g.Execute()
}
