package bootstrap

import "gorm.io/gorm"

type Application struct {
	Env        *Env
	PostgresDB gorm.DB
}

func App() Application {
	app := &Application{}
	app.Env = NewEnv()
	app.PostgresDB = NewPostgres(app.Env)
	return *app
}
