package usecase

import (
	"log"

	"github.com/jinzhu/gorm"

	"Hybrid-blog/helpers/database"
	"Hybrid-blog/models"
	model "Hybrid-blog/models"
	repository "Hybrid-blog/repositories"
)

type UserCaseBranchDepart interface {
	GetDepartments() ([]*model.Department, error)
	GetPositions() ([]*model.Position, error)
	Getbranch() ([]*model.Branch, error)
}

type (
	userCaseBranchDepart struct {
		db     *gorm.DB
		bdRepo repository.BranchDepartment
	}
)

func NewUCaseBranchDepart() UserCaseBranchDepart {
	return &userCaseBranchDepart{
		db:     database.DB(),
		bdRepo: repository.NewBranchDepartmentRepo(),
	}
}

func (u *userCaseBranchDepart) GetDepartments() ([]*models.Department, error) {
	departments, err := u.bdRepo.GetDepartments(u.db)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return departments, nil
}

func (u *userCaseBranchDepart) GetPositions() ([]*model.Position, error) {
	positions, err := u.bdRepo.GetPositions(u.db)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return positions, nil
}

func (u *userCaseBranchDepart) Getbranch() ([]*model.Branch, error) {
	branches, err := u.bdRepo.Getbranch(u.db)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return branches, nil
}
