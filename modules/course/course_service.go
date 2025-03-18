package course

import (
	"github.com/creasty/defaults"
	"github.com/irvanherz/gourze/modules/course/dto"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type CourseService interface {
	FindManyCourses(filter *dto.CourseFilterInput) ([]Course, int64, error)
	CreateCourse(input *dto.CourseCreateInput) (*Course, error)
	FindCourseByID(id uint) (*Course, error)
	UpdateCourseByID(id uint, input *dto.CourseUpdateInput) (*Course, error)
	DeleteCourseByID(id uint) (*Course, error)
}

type courseService struct {
	Db *gorm.DB
}

func NewCourseService(db *gorm.DB) CourseService {
	return &courseService{Db: db}
}

func (s *courseService) FindManyCourses(filter *dto.CourseFilterInput) ([]Course, int64, error) {
	var courses []Course
	var count int64

	if err := defaults.Set(filter); err != nil {
		return nil, 0, err
	}
	query := s.Db
	query = filter.ApplyFilter(query)

	if err := query.Model(&Course{}).Count(&count).Error; err != nil {
		return nil, 0, err
	}

	query = filter.ApplyPagination(query)

	if err := query.Preload("User").Preload("Category").Find(&courses).Error; err != nil {
		return nil, 0, err
	}
	return courses, count, nil
}

func (s *courseService) CreateCourse(input *dto.CourseCreateInput) (*Course, error) {
	var course Course
	copier.Copy(&course, &input)

	if err := s.Db.Preload("User").Preload("Category").Create(&course).Error; err != nil {
		return nil, err
	}
	return &course, nil
}

func (s *courseService) FindCourseByID(id uint) (*Course, error) {
	var course Course
	if err := s.Db.Preload("User").Preload("Category").First(&course, id).Error; err != nil {
		return nil, err
	}
	return &course, nil
}

func (s *courseService) UpdateCourseByID(id uint, input *dto.CourseUpdateInput) (*Course, error) {
	var course Course
	if err := s.Db.First(&course, id).Error; err != nil {
		return nil, err
	}
	copier.Copy(&course, &input)
	if err := s.Db.Preload("User").Preload("Category").Save(&course).Error; err != nil {
		return nil, err
	}
	return &course, nil
}

func (s *courseService) DeleteCourseByID(id uint) (*Course, error) {
	var course Course
	if err := s.Db.First(&course, id).Error; err != nil {
		return nil, err
	}
	if err := s.Db.Preload("User").Preload("Category").Delete(&Course{}, id).Error; err != nil {
		return nil, err
	}
	return &course, nil
}
