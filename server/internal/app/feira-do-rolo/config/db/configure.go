package db

import "gorm.io/gorm"

func ConfigureDatabase() (*gorm.DB, error) {
	db := &gorm.DB{}
	db, err := openConnectionToDatabase()
	if err != nil {
		return db, err
	}
	err = migrateAllModels(db)
	if err != nil {
		return db, err
	}

	return db, nil

}
