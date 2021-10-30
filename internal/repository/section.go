package repository

import (
	"fmt"

	"github.com/Masterminds/squirrel"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"zen_api/internal/datastruct"
)

type SectionQuery interface {
	GetSectionByID(sectionID int64) (*datastruct.CourseSection, error)
	GetCourseSections(courseID, limit int64) ([]datastruct.CourseSection, error)
	CreateCourseSection(section datastruct.CourseSection) (*int64, error)
	UpdateCourseSection(section datastruct.CourseSection) (*datastruct.CourseSection, error)
	DeleteCourseSection(id int64) error
}

type sectionQuery struct{}

func (s *sectionQuery) GetSectionByID(sectionID int64) (*datastruct.CourseSection, error) {
	qb := pgQb().
		Select("id", "course_id", "title", "description", "file_name").
		From(datastruct.CourseSectionTableName).
		Where(squirrel.Eq{"id": sectionID})

	var section datastruct.CourseSection

	err := qb.QueryRow().Scan(&section.ID, &section.CourseID, &section.Title, &section.Description, &section.FilePath)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "cannot get section by id: %v", err)
	}

	return &section, nil
}

func (s *sectionQuery) GetCourseSections(courseID, limit int64) ([]datastruct.CourseSection, error) {
	qb := pgQb().
		Select("id, course_id, title, description, file_name").
		From(datastruct.CourseSectionTableName).
		Where(squirrel.Eq{"course_id": courseID}).
		OrderBy("id").
		Limit(uint64(limit))

	var sections []datastruct.CourseSection
	var section datastruct.CourseSection
	rows, err := qb.Query()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err = rows.Scan(&section.ID, &section.CourseID, &section.Title, &section.Description, &section.FilePath)
		if err != nil {
			return nil, err
		}
		sections = append(sections, section)
	}

	return sections, nil
}

func (s *sectionQuery) CreateCourseSection(section datastruct.CourseSection) (*int64, error) {
	qb := pgQb().
		Insert(datastruct.CourseSectionTableName).
		Columns("course_id", "title", "description", "file_name").
		Values(section.CourseID, section.Title, section.Description, section.FilePath).
		Suffix("RETURNING id")

	var id int64
	err := qb.QueryRow().Scan(&id)
	if err != nil {
		return nil, fmt.Errorf("cannot create course section")
	}
	return &id, nil
}

func (s *sectionQuery) UpdateCourseSection(section datastruct.CourseSection) (*datastruct.CourseSection, error) {
	qb := pgQb().Update(datastruct.CourseSectionTableName).SetMap(map[string]interface{}{
		"title":       section.Title,
		"description": section.Description,
		"file_name":   section.FilePath,
	}).Where(squirrel.Eq{"id": section.ID}).Suffix("RETURNING *")

	var updatedSection datastruct.CourseSection

	err := qb.QueryRow().Scan(&updatedSection.ID, &updatedSection.CourseID,
		&updatedSection.Title, &updatedSection.Description, &updatedSection.FilePath)
	if err != nil {
		return nil, err
	}
	return &updatedSection, nil
}

func (s *sectionQuery) DeleteCourseSection(id int64) error {
	qb := pgQb().
		Delete(datastruct.CourseSectionTableName).
		From(datastruct.CourseSectionTableName).
		Where(squirrel.Eq{"id": id})

	_, err := qb.Exec()
	if err != nil {
		return err
	}
	return nil
}
