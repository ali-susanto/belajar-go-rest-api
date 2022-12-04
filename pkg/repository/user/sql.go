package user

import (
	"context"

	"github.com/google/uuid"

	"gorm.io/gorm"
)

// userSQL struct
type userSQL struct {
	db *gorm.DB
}

func newSQL(db *gorm.DB) *userSQL {
	repo := &userSQL{
		db: db,
	}
	return repo
}

// All func
func (u userSQL) All(c context.Context) (users []*Table, err error) {
	users = make([]*Table, 0)
	err = u.db.WithContext(c).Find(&users).Error
	return users, err
}

// Create func
func (u userSQL) Create(c context.Context, user *Table) (*Table, error) {
	err := u.db.WithContext(c).Create(&user).Error
	return user, err
}

// FindByEmail func
func (u userSQL) FindByEmail(c context.Context, email string) (*Table, error) {
	user := &Table{}
	err := u.db.WithContext(c).Where(&Table{
		Email: email,
	}).First(&user).Error
	return user, err
}

// FindByID func
func (u userSQL) FindByID(c context.Context, id uuid.UUID) (*Table, error) {
	user := &Table{}
	err := u.db.WithContext(c).Where("id = ?", id).First(&user).Error
	return user, err
}

// ChangePassword func
func (u userSQL) ChangePassword(c context.Context, id uuid.UUID, password string) (*Table, error) {
	user, err := u.FindByID(c, id)
	if err != nil {
		return nil, err
	}

	user.Password = password
	err = u.db.WithContext(c).Save(user).Error

	return user, err
}
