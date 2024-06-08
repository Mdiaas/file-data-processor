package firestorage

import (
	"time"

	"github.com/mdiaas/processor/internal/core/entity"
)

func StudentToStudentDb(student entity.Student) StudentDb {
	return StudentDb{
		Id:             student.Id,
		FullName:       student.FullName,
		DocumentNumber: student.DocumentNumber,
		BirthDate:      student.BirthDate,
		MotherName:     student.MotherName,
		FatherName:     student.FatherName,
		Errors:         student.Errors,
		CreatedAt:      time.Now(),
	}
}
