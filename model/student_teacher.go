package model

import (
	"fmt"
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
	Student   Student   `gorm:"foreignKey:StudentId;references:id"`
	Teacher   Teacher   `gorm:"foreignKey:TeacherId;references:id"`
}

func (m *StudentTeacher) TeacherRegisterStudent(input []StudentTeacher) error {
	// db := DB.Model(m)
	// tx := db.Begin()
	err := DB.Transaction(func(tx *gorm.DB) error {
		tx_result := tx.Create(input)
		if tx_result.Error != nil {
			return tx_result.Error
		}
		fmt.Println(tx_result.RowsAffected)
		// return nil will commit the whole transaction
		return nil
	})
	return err
}
