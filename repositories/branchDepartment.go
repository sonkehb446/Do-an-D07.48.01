package repository

import (
	"github.com/jinzhu/gorm"

	"Hybrid-blog/models"
	model "Hybrid-blog/models"
)

// class get thong tin position branch department
type BranchDepartment interface {
	GetDepartments(db *gorm.DB) ([]*model.Department, error)
	GetPositions(db *gorm.DB) ([]*model.Position, error)
	Getbranch(db *gorm.DB) ([]*model.Branch, error)
}

type branchDepartment struct {
}

func NewBranchDepartmentRepo() BranchDepartment {
	return &branchDepartment{}
}

func (r *branchDepartment) GetDepartments(db *gorm.DB) ([]*models.Department, error) {
	var list []*models.Department
	result := db.Table("public.departments").Select("*").Scan(&list)
	if result.Error != nil {
		return nil, result.Error
	}
	return list, nil
}

func (r *branchDepartment) GetPositions(db *gorm.DB) ([]*models.Position, error) {
	var list []*models.Position
	result := db.Table("public.positions").Select("*").Scan(&list)
	if result.Error != nil {
		return nil, result.Error
	}
	return list, nil
}

func (r *branchDepartment) Getbranch(db *gorm.DB) ([]*models.Branch, error) {
	var list []*models.Branch
	result := db.Table("public.branches").Select("*").Scan(&list)
	if result.Error != nil {
		return nil, result.Error
	}
	return list, nil
}
