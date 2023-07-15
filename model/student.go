package model

import (
	"errors"
	"fmt"

	"gorm.io/gorm"
)

func (Student) TableName() string {
	return "student"
}

type Student struct {
	Id     int    `gorm:"primaryKey"`
	Email  string `gorm:"column:email" json:"email"`
	Status uint   `gorm:"column:status" json:"status"`
	// Teacher        []Teacher        `gorm:"many2many:student_teacher;foreignKey:Id;joinForeignKey:StudentId;References:Id;joinReferences:TeacherId"`
	StudentTeacher []StudentTeacher `gorm:"foreignKey:StudentId;references:id" json:"student_teacher"`
}

func (m *Student) SelectOne(email string, status int) (result Student, err error) {
	db := DB.Model(m)
	db.Where("Email = ?", email)
	db.Where("status = ?", status)
	err = db.First(&result).Error
	return
}

func (m *Student) SelectFromRelations(teacher_ids []int) (result []Student, err error) {
	// DB.Raw("select s.* from `student` s join `student_teacher` st on s.`id` = st.`student_id` where st.teacher_id  in (?) and s.Status = 1", ids).Scan(&result)
	db := DB.Model(m)
	db.Preload("StudentTeacher", func(db *gorm.DB) *gorm.DB {
		return db.Where("teacher_id in ?", teacher_ids)
	})
	err = db.Find(&result).Error
	return
}

func (m *Student) UpdateStatus(email string, new_status int) (result Student, err error) {
	tx := DB.Begin()
	ret := tx.Model(m).Where("email = ?", email).Where("status != ?", new_status).Update("status", SUSPENDED)
	fmt.Println(ret.RowsAffected)
	if ret.RowsAffected == 0 {
		err = errors.New("0 row affected")
	}
	if err == nil {
		tx.Commit()
		result, err = m.SelectOne(email, new_status)
	} else {
		tx.Rollback()
	}
	return
}

func (m *Student) SelectFromEmails(email []string) (result []Student, err error) {
	db := DB.Model(m)
	db.Where("email in ?", email)
	db.Where("status != ?", SUSPENDED)
	err = db.Find(&result).Error
	return
}

func (m *Student) GetNotificationTarget(email []string, under_teacher int) (result []string, err error) {
	DB.Raw("SELECT DISTINCT(s.email) FROM student s LEFT JOIN student_teacher st on s.id = st.student_id WHERE (st.teacher_id = ?) OR s.email in (?)", under_teacher, email).Scan(&result)
	return
}

func (m *Student) RetrieveEmails(students []Student) (emails []string) {
	for _, i := range students {
		emails = append(emails, i.Email)
	}
	return
}
