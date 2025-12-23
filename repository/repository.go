package repository

import "gorm.io/gorm"

type Repository struct {
	db *gorm.DB
}

func (r *Repository) MigrateRepositories() {
	r.db.AutoMigrate()
}
