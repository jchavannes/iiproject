package db

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var _db *gorm.DB

var (
	dbInterfaces = []interface{}{
		Key{},
		User{},
		Session{},
		Profile{},
		Contact{},
		UserContact{},
	}
)

func getDb() (*gorm.DB, error) {
	if _db == nil {
		var err error
		_db, err = gorm.Open("sqlite3", "iiproject.db")
		if err != nil {
			return _db, errors.New("Failed to connect to database\n")
		}
		for _, iface := range dbInterfaces {
			result := _db.AutoMigrate(iface)
			if result.Error != nil {
				return result, result.Error
			}
		}
	}
	return _db, nil
}

func create(value interface{}) *gorm.DB {
	db, _ := getDb()
	result := db.Create(value)
	if result.Error != nil {
		fmt.Printf("Db error: %s\n", result.Error)
		return result
	}
	return result
}

func find(out interface{}, where ...interface{}) *gorm.DB {
	db, _ := getDb()
	result := db.Find(out, where...)
	if result.Error != nil && !result.RecordNotFound() {
		fmt.Printf("Db error: %s\n", result.Error)
		return result
	}
	return result
}

func findString(out interface{}, where string, args ...string) *gorm.DB {
	db, _ := getDb()
	result := db.Where(where, args).Find(out)
	if result.Error != nil && !result.RecordNotFound() {
		fmt.Printf("Db error: %s\n", result.Error)
		return result
	}
	return result
}

func save(value interface{}) *gorm.DB {
	db, _ := getDb()
	if db.Error != nil {
		fmt.Printf("Db error: %s\n", db.Error)
		return db
	}
	result := db.Save(value)
	if result.Error != nil {
		fmt.Printf("Db error: %s\n", result.Error)
		return result
	}
	return result
}
