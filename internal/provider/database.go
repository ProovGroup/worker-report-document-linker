package provider

import (
	"fmt"

	env "github.com/ProovGroup/lib-env"
	"github.com/ProovGroup/lib-env/database"
)

var didInit = false
var db database.Database

func initDB() {
	var dbExists bool
	db, dbExists = e.GetDB(env.NewDBSelector("main"))
	if !dbExists {
		fmt.Println("Database: 'main' not found!")
		panic("Database: 'main' not found!")
	}
	didInit = true
}

func GetDB() database.Database {
	if !didInit {
		initDB()
	}
	return db
}
