package repository

import (
	"finpro/model"

	"gorm.io/gorm"
)

type AdminRepository interface {
	FindByEmail(email string) (*model.Admin, error)
}

type adminRepository struct {
	db *gorm.DB
}

func NewAdminRepository(db *gorm.DB) AdminRepository {
	return &adminRepository{db}
}

func (r *adminRepository) FindByEmail(email string) (*model.Admin, error) {
	var admin model.Admin
	err := r.db.Where("email = ?", email).First(&admin).Error
	return &admin, err
}
