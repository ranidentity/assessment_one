package model

import (
	"time"

	"gorm.io/gorm"
)

func (StudentTeacher) TableName() string {
	return "student_teacher"
}

type StudentTeacher struct {
	Id        int       `gorm:"primaryKey"`
	StudentId int       `gorm:"column:student_id" json:"student_id"`
	TeacherId int       `gorm:"column:teacher_id" json:"teacher_id"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt time.Time `gorm:"column:deleted_at" json:"deleted_at"`
}

type StudentTeacherWithEmail struct {
	StudentId int
	TeacherId int
	Email     string
}

func (m *StudentTeacher) TeacherRegisterStudent(input []StudentTeacherWithEmail) (success_student []string, err error) {
	err = DB.Transaction(func(tx *gorm.DB) error {
		for _, i := range input {
			var count int64
			tx.Model(m).Where("student_id = ?", i.StudentId).Where("teacher_id = ?", i.TeacherId).Count(&count)
			if count == 0 {
				new := StudentTeacher{StudentId: i.StudentId, TeacherId: i.TeacherId}
				err = tx.Create(&new).Error
				if err != nil {
					return err
				} else {
					success_student = append(success_student, i.Email)
				}
			}
		}
		return nil
	})
	return
}

func (m *StudentTeacher) List(teacher_ids []int, relative string) (list []StudentTeacher, err error) {
	db := DB.Model(m)
	db.Where("teacher_id in ?", teacher_ids)
	switch relative {
	case "student":
		db.Preload("Student")
	case "teacher":
		db.Preload("Teacher")
	}
	err = db.First(&list).Error
	return
}
