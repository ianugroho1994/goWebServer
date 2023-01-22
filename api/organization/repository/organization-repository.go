package repository

import (
	"errors"
	"fmt"
	"hardtmann/smartlab/api/organization/domain"

	"gorm.io/gorm"
)

type OrganizationRepository struct {
	database  *gorm.DB
	tableName string
}

func NewOrganizationRepository(db *gorm.DB) domain.IOrganizationRepository {
	return &OrganizationRepository{
		database: db,
	}
}

func (o *OrganizationRepository) Create(org *domain.Organization) error {
	err := o.database.Table(o.tableName).Create(org).Error
	if err != nil {
		return fmt.Errorf("error create organization : %w", err)
	}
	return nil
}

func (o *OrganizationRepository) Delete(id int64) error {
	if id == 1 {
		return errors.New("cannot delete top level organization")
	}
	return o.database.Table(o.tableName).Delete(&domain.Organization{}, id).Error
}

func (o *OrganizationRepository) Update(org *domain.Organization) error {
	q := o.database.Table(o.tableName).Save(org)
	if q.Error != nil {
		return fmt.Errorf("error updating organization: %w", q.Error)
	}
	if q.RowsAffected == 0 {
		return errors.New("no organization record updated")
	}
	return nil
}

func (o *OrganizationRepository) GetList(page, limit uint) ([]domain.Organization, int, error) {
	res := []domain.Organization{}
	var totalRecord int64
	err := o.database.Model(domain.Organization{}).Count(&totalRecord).Error
	if err != nil {
		return nil, 0, err
	}

	err = o.database.Offset(int((page - 1) * limit)).Limit(int(limit)).Find(&res).Error
	if err != nil {
		return nil, 0, err
	}

	return res, int(totalRecord), nil
}

func (o *OrganizationRepository) GetByID(id int64) (*domain.Organization, error) {
	res := &domain.Organization{}
	err := o.database.Table(o.tableName).Where("id = ?", id).First(res).Error
	if err != nil {
		return nil, fmt.Errorf("error get organization by ID: %w", err)
	}
	return res, nil
}

func (o *OrganizationRepository) GetAll() ([]domain.Organization, error) {
	res := []domain.Organization{}
	err := o.database.Table(o.tableName).Model(domain.Organization{}).Find(&res).Error
	if err != nil {
		return nil, fmt.Errorf("error get all organization: %w", err)
	}
	return res, nil
}
