package service

import (
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"zen_api/internal/datastruct"
	"zen_api/internal/dto"
	"zen_api/internal/repository"
)

type SectionService interface {
	GetCourseSections(userID int64, courseID int64) ([]dto.Section, error)
	CreateCourseSection(section dto.Section) (*int64, error)
	UpdateCourseSection(section dto.Section) (*datastruct.CourseSection, error)
	DeleteCourseSection(sectionID int64, userID int64) error
	GetSectionByID(id int64) (*datastruct.CourseSection, error)
}

type sectionService struct {
	dao repository.DAO
}

func NewSectionService(dao repository.DAO) SectionService {
	return &sectionService{dao: dao}
}

func (s *sectionService) GetSectionByID(id int64) (*datastruct.CourseSection, error) {
	section, err := s.dao.NewSectionQuery().GetSectionByID(id)
	if err != nil {
		return nil, err
	}
	return section, nil
}

func (s *sectionService) GetCourseSections(userID int64, courseID int64) ([]dto.Section, error) {
	course, err := s.dao.NewCourseQuery().GetCourse(courseID)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "course doesn't exist %v", err)
	}

	var user *datastruct.Person
	var transaction *datastruct.Transaction

	var allSections []dto.Section
	var sections []datastruct.CourseSection
	if &userID != nil || userID != 0 {
		user, err = s.dao.NewUserQuery().GetUser(userID)
		if err != nil {
			return nil, err
		}
		transaction, err = s.dao.NewTransactionQuery().GetTransactionByUserCourseIDs(userID, courseID)
		if err != nil {
			log.Printf("user didn't paid for the course: %v", err)
		}
		if transaction.Status == datastruct.PAID || user.Role == datastruct.ADMIN || course.Price == 0 || course.UserID == userID {
			sections, err = s.dao.NewSectionQuery().GetCourseSections(courseID, 100)
			if err != nil {
				return nil, err
			}
		} else {
			sections, err = s.dao.NewSectionQuery().GetCourseSections(courseID, 2)
			if err != nil {
				return nil, err
			}
		}
	} else {
		sections, err = s.dao.NewSectionQuery().GetCourseSections(courseID, 2)
		if err != nil {
			return nil, err
		}
	}

	for _, section := range sections {
		allSections = append(allSections, dto.Section{ID: section.ID, CourseID: section.CourseID,
			Title: section.Title, Description: section.Description, FilePath: section.FilePath})
	}

	return allSections, nil
}

func (s *sectionService) CreateCourseSection(section dto.Section) (*int64, error) {
	user, err := s.dao.NewUserQuery().GetUser(section.UserID)
	if err != nil {
		return nil, err
	}

	course, err := s.dao.NewCourseQuery().GetCourse(section.CourseID)
	if err != nil {
		return nil, err
	}

	if course.UserID == section.UserID || user.Role == datastruct.ADMIN {
		id, err := s.dao.NewSectionQuery().CreateCourseSection(datastruct.CourseSection{CourseID: section.CourseID,
			Title: section.Title, Description: section.Description, FilePath: section.FilePath})
		if err != nil {
			return nil, err
		}
		return id, nil
	}
	return nil, status.Errorf(codes.PermissionDenied, "you don't have access %v", err)
}

func (s *sectionService) UpdateCourseSection(section dto.Section) (*datastruct.CourseSection, error) {
	user, err := s.dao.NewUserQuery().GetUser(section.UserID)
	if err != nil {
		return nil, err
	}

	course, err := s.dao.NewCourseQuery().GetCourse(section.CourseID)
	if err != nil {
		return nil, err
	}

	if user.Role == datastruct.ADMIN || course.UserID == section.UserID {
		updatedSection, err := s.dao.NewSectionQuery().UpdateCourseSection(datastruct.CourseSection{ID: section.ID,
			Title: section.Title, Description: section.Description, FilePath: section.FilePath})
		if err != nil {
			return nil, err
		}
		return updatedSection, nil
	}
	return nil, status.Errorf(codes.PermissionDenied, "you can change only your course sections")
}

func (s *sectionService) DeleteCourseSection(sectionID int64, userID int64) error {
	user, err := s.dao.NewUserQuery().GetUser(userID)
	if err != nil {
		return err
	}

	section, err := s.dao.NewSectionQuery().GetSectionByID(sectionID)
	if err != nil {
		return err
	}

	course, err := s.dao.NewCourseQuery().GetCourse(section.CourseID)
	if err != nil {
		return err
	}

	if user.Role == datastruct.ADMIN || section.CourseID == course.ID && course.UserID == user.ID {
		err = s.dao.NewSectionQuery().DeleteCourseSection(sectionID)
		if err != nil {
			return err
		}
	}

	return nil
}
