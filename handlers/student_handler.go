package handlers

import (
	"net/http"

	"example.com/student-api/models"
	"example.com/student-api/services"
	"github.com/gin-gonic/gin"
)

type StudentHandler struct {
	Service *services.StudentService
}

func errorResponse(c *gin.Context, code int, msg string) {
	c.JSON(code, gin.H{"error": msg})
}

func (h *StudentHandler) GetStudents(c *gin.Context) {

	data, err := h.Service.GetStudents()
	if err != nil {
		errorResponse(c, 500, "internal error")
		return
	}

	c.JSON(http.StatusOK, data)
}

func (h *StudentHandler) GetStudentByID(c *gin.Context) {

	id := c.Param("id")

	s, err := h.Service.GetStudentByID(id)
	if err != nil {
		errorResponse(c, 404, "Student not found")
		return
	}

	c.JSON(http.StatusOK, s)
}

func (h *StudentHandler) CreateStudent(c *gin.Context) {

	var s models.Student

	if err := c.ShouldBindJSON(&s); err != nil {
		errorResponse(c, 400, "invalid request body")
		return
	}

	if err := h.Service.CreateStudent(s); err != nil {
		errorResponse(c, 400, err.Error())
		return
	}

	c.JSON(http.StatusCreated, s)
}

func (h *StudentHandler) UpdateStudent(c *gin.Context) {

	id := c.Param("id")

	var s models.Student
	if err := c.ShouldBindJSON(&s); err != nil {
		errorResponse(c, 400, "invalid request body")
		return
	}

	err := h.Service.UpdateStudent(id, s)
	if err != nil {
		if err.Error() == "not found" {
			errorResponse(c, 404, "Student not found")
			return
		}
		errorResponse(c, 400, err.Error())
		return
	}

	updated, _ := h.Service.GetStudentByID(id)
	c.JSON(http.StatusOK, updated)
}

func (h *StudentHandler) DeleteStudent(c *gin.Context) {

	id := c.Param("id")

	err := h.Service.DeleteStudent(id)
	if err != nil {
		errorResponse(c, 404, "Student not found")
		return
	}

	c.Status(http.StatusNoContent)
}
