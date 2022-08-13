package app

import (
	gormDBS "gitlab.com/d6825/golang_template/platform/database"
	"gorm.io/gorm"
)

type DatabaseSql struct{
    Gorm *gorm.DB
}
type App struct {
	db *DatabaseSql
}

func Init() *App {
    gorms := gormDBS.GormMysqlConnection()
    return &App{
        db: &DatabaseSql{
            Gorm:gorms,
        },
    }
}

func (app *App) SetDbGorm(db *gorm.DB) {
    app.db.Gorm = db
}

func (app *App) GetDbGorm() *gorm.DB {
	return app.db.Gorm
}