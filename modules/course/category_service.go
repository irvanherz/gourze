package course

import (
	"github.com/creasty/defaults"
	"github.com/irvanherz/gourze/modules/course/dto"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type CategoryService interface {
	FindManyCategories(filter *dto.CategoryFilterInput) ([]Category, int64, error)
	CreateCategory(input *dto.CategoryCreateInput) (*Category, error)
	FindCategoryByID(id uint) (*Category, error)
	UpdateCategoryByID(id uint, input *dto.CategoryUpdateInput) (*Category, error)
	DeleteCategoryByID(id uint) (*Category, error)
}

type categoryService struct {
	Db *gorm.DB
}

func NewCategoryService(db *gorm.DB) CategoryService {
	return &categoryService{Db: db}
}

func (s *categoryService) FindManyCategories(filter *dto.CategoryFilterInput) ([]Category, int64, error) {
	var categories []Category
	var count int64

	if err := defaults.Set(filter); err != nil {
		return nil, 0, err
	}
	query := s.Db
	query = filter.ApplyFilter(query)

	if err := query.Model(&Category{}).Count(&count).Error; err != nil {
		return nil, 0, err
	}

	query = filter.ApplyPagination(query)

	if err := query.Find(&categories).Error; err != nil {
		return nil, 0, err
	}
	return categories, count, nil
}

func (s *categoryService) CreateCategory(input *dto.CategoryCreateInput) (*Category, error) {
	var category Category
	copier.Copy(&category, &input)

	if err := s.Db.Create(&category).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

func (s *categoryService) FindCategoryByID(id uint) (*Category, error) {
	var category Category
	if err := s.Db.First(&category, id).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

func (s *categoryService) UpdateCategoryByID(id uint, input *dto.CategoryUpdateInput) (*Category, error) {
	var category Category
	if err := s.Db.First(&category, id).Error; err != nil {
		return nil, err
	}
	copier.Copy(&category, &input)
	if err := s.Db.Save(&category).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

func (s *categoryService) DeleteCategoryByID(id uint) (*Category, error) {
	var category Category
	if err := s.Db.First(&category, id).Error; err != nil {
		return nil, err
	}
	if err := s.Db.Delete(&Category{}, id).Error; err != nil {
		return nil, err
	}
	return &category, nil
}
