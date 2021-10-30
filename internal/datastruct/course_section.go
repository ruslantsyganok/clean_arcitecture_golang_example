package datastruct

const CourseSectionTableName = "course_section"

type CourseSection struct {
	ID          int64  `db:"id"`
	CourseID    int64  `db:"course_id"`
	Title       string `db:"title"`
	Description string `db:"description"`
	FilePath    string `db:"file_name"`
}
