package services

import (
	"errors"

	"example.com/student-api/models"
	"example.com/student-api/repositories"
)

type StudentService struct {
	Repo *repositories.StudentRepository
}

func validateStudent(s models.Student, checkID bool) error {

	if checkID && s.ID == "" {
		return errors.New("id must not be empty")
	}

	if s.Name == "" {
		return errors.New("name must not be empty")
	}

	if s.GPA < 0 || s.GPA > 4 {
		return errors.New("gpa must be between 0.00 and 4.00")
	}

	return nil
}

func (s *StudentService) GetStudents() ([]models.Student, error) {
	return s.Repo.GetAll()
}

func (s *StudentService) GetStudentByID(id string) (*models.Student, error) {
	return s.Repo.GetByID(id)
}

func (s *StudentService) CreateStudent(st models.Student) error {

	if err := validateStudent(st, true); err != nil {
		return err
	}

	return s.Repo.Create(st)
}

func (s *StudentService) UpdateStudent(id string, st models.Student) error {

	if err := validateStudent(st, false); err != nil {
		return err
	}

	return s.Repo.Update(id, st)
}

func (s *StudentService) DeleteStudent(id string) error {
	return s.Repo.Delete(id)
}
