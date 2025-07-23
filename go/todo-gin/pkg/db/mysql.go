package db

import (
	"log"
	"os/user"

	mysql_go "github.com/go-sql-driver/mysql"
	"github.com/hardikchoksi151/todo-gin/config"
	"github.com/hardikchoksi151/todo-gin/internal/todo"
	"github.com/k0kubun/pp/v3"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDB(cfg *config.Config) *gorm.DB {
	pp.Println(cfg)
	cnfg := mysql_go.Config{User: cfg.DBUser, Passwd: cfg.DBPassword, Addr: cfg.DBHost + cfg.DBPort, DBName: cfg.DBName}
	// dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
	// cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)
	dsn := cnfg.FormatDSN()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("could not connect to the database: %v", err)
	}
	return db
}

func RunMigrations(db *gorm.DB) {
	log.Println("Running migrations...")

	// Ensure the User table is created first
	if err := db.AutoMigrate(&user.User{}); err != nil {
		log.Fatalf("Migration failed for User: %v", err)
	}

	// Then migrate the Todo table
	if err := db.AutoMigrate(&todo.Todo{}); err != nil {
		log.Fatalf("Migration failed for Todo: %v", err)
	}

	log.Println("Migrations completed successfully")
}
