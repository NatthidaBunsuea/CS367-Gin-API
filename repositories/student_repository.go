package repositories

import (
	"database/sql"
	"errors"

	"example.com/student-api/models"
)

type StudentRepository struct {
	DB *sql.DB
}

func (r *StudentRepository) GetAll() ([]models.Student, error) {

	rows, err := r.DB.Query("SELECT id,name,major,gpa FROM students")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var students []models.Student

	for rows.Next() {
		var s models.Student
		rows.Scan(&s.ID, &s.Name, &s.Major, &s.GPA)
		students = append(students, s)
	}

	return students, nil
}

func (r *StudentRepository) GetByID(id string) (*models.Student, error) {

	row := r.DB.QueryRow(
		"SELECT id,name,major,gpa FROM students WHERE id=?",
		id,
	)

	var s models.Student
	err := row.Scan(&s.ID, &s.Name, &s.Major, &s.GPA)

	if err == sql.ErrNoRows {
		return nil, errors.New("not found")
	}
	if err != nil {
		return nil, err
	}

	return &s, nil
}

func (r *StudentRepository) Create(s models.Student) error {

	_, err := r.DB.Exec(
		"INSERT INTO students(id,name,major,gpa) VALUES(?,?,?,?)",
		s.ID, s.Name, s.Major, s.GPA,
	)

	return err
}

func (r *StudentRepository) Update(id string, s models.Student) error {

	res, err := r.DB.Exec(
		"UPDATE students SET name=?, major=?, gpa=? WHERE id=?",
		s.Name, s.Major, s.GPA, id,
	)

	if err != nil {
		return err
	}

	affected, _ := res.RowsAffected()
	if affected == 0 {
		return errors.New("not found")
	}

	return nil
}

func (r *StudentRepository) Delete(id string) error {

	res, err := r.DB.Exec(
		"DELETE FROM students WHERE id=?",
		id,
	)

	if err != nil {
		return err
	}

	affected, _ := res.RowsAffected()
	if affected == 0 {
		return errors.New("not found")
	}

	return nil
}
