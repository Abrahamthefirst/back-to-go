package repository

import (
	"context"
	"errors"

	"github.com/Abrahamthefirst/back-to-go/internal/entities"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

type UserModel struct {
	gorm.Model
	Email    string `gorm:"uniqueIndex"`
	Password string
	Username string
}

func (*UserModel) TableName() string {
	return "users"
}

func (m *UserModel) toDomain() *entities.User {
	return &entities.User{
		ID:        m.ID,
		Email:     m.Email,
		Username:  m.Username,
		Password:  m.Password,
		CreatedAt: m.CreatedAt,
	}
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db,
	}
}

func (repo *UserRepository) FindByEmail(ctx context.Context, email string) (*entities.User, error) {
	dbUser := UserModel{}

	if err := repo.db.WithContext(ctx).Where("email = ?", email).First(&dbUser).Error; err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {

			return nil, entities.RecordNotFound
		}
		return nil, err
	}
	return dbUser.toDomain(), nil
}

func (r *UserRepository) Create(user entities.User) (*entities.User, error) {
	dbUser := UserModel{
		Email:    user.Email,
		Password: user.Password,
		Username: user.Username,
	}

	if err := r.db.Create(&dbUser).Error; err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return nil, entities.ErrConflict
		}
		return nil, err
	}
	return dbUser.toDomain(), nil

}

// func (repo *UserRepository) FindByID(id uint) (*User, error) {
// 	var user User
// 	if err := repo.db.First(&user, id).Error; err != nil {
// 		return nil, err
// 	}
// 	return &user, nil
// }

// func (repo *UserRepository) FindAll() ([]User, error) {
// 	var users []User
// 	result := repo.db.Find(&users)
// 	return users, result.Error
// }
// func (repo *UserRepository) Update(user User) (*User, error) {

// 	result := repo.db.Model(&user).Updates(map[string]interface{}{
// 		"active": false,
// 		"age":    0,
// 	})

// 	if result.Error != nil {
// 		return nil, result.Error
// 	}

// 	return &user, nil
// }
// func (repo *UserRepository) SoftDelete(id uint) (uint, error) {
// 	// Soft delete a user by ID
// 	if err := repo.db.Delete(&User{}, id).Error; err != nil {
// 		return id, err
// 	}

// 	return id, nil
// }

// func (repo *UserRepository) HardDelete(id uint) (uint, error) {
// 	// Permanently delete a user record
// 	if err := repo.db.Unscoped().Delete(&User{}, id).Error; err != nil {
// 		return id, err
// 	}
// 	return id, nil
// }
