package repository

import (
	"belajar-go-rest-api/entity"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

// TokenRepository struct
type TokenRepository struct {
	db *gorm.DB
}

// NewTokenRepository func
func NewTokenRepository(database *gorm.DB) entity.TokenRepository {
	return &TokenRepository{
		db: database,
	}
}

// Create func
func (u TokenRepository) Create(token *entity.Token) (*entity.Token, error) {
	err := u.db.Create(&token).Error
	return token, err
}

// FindByUserID func
func (u TokenRepository) FindByUserID(id uuid.UUID) (*entity.Token, error) {
	token := &entity.Token{}
	err := u.db.Where("id = ?", id).First(&token).Error
	return token, err
}

// FindByToken func
func (u TokenRepository) FindByToken(token string) (*entity.Token, error) {
	tokenDB := &entity.Token{}
	err := u.db.Where("token = ?", token).First(&tokenDB).Error
	return tokenDB, err
}
