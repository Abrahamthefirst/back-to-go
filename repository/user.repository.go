package repository

import (
	"errors"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

type User struct {
	ID       int
	Email    string
	Password string
	Deleted  gorm.DeletedAt
	Name     string
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db,
	}
}

func (repo *UserRepository) FindByEmail(email string) (*User, error) {
	var user User

	if err := repo.db.Where("email = ?", email).First(&user).Error; err != nil {

		if errors.Is(RecordNotFound, err) {
			return nil, RecordNotFound
		}
		return nil, err
	}

	return &user, nil
}

func (repo *UserRepository) Create(user User) (*User, error) {
	if err := repo.db.Create(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo *UserRepository) FindByID(id uint) (*User, error) {
	var user User
	if err := repo.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo *UserRepository) FindAll() ([]User, error) {
	var users []User
	result := repo.db.Find(&users)
	return users, result.Error
}
func (repo *UserRepository) Update(user User) (*User, error) {

	result := repo.db.Model(&user).Updates(map[string]interface{}{
		"active": false,
		"age":    0,
	})

	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}
func (repo *UserRepository) SoftDelete(id uint) (uint, error) {
	// Soft delete a user by ID
	if err := repo.db.Delete(&User{}, id).Error; err != nil {
		return id, err
	}

	return id, nil
}

func (repo *UserRepository) HardDelete(id uint) (uint, error) {
	// Permanently delete a user record
	if err := repo.db.Unscoped().Delete(&User{}, id).Error; err != nil{
		return id, err
	}
	return id, nil
}
