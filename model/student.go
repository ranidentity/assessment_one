package model

func (Student) TableName() string {
	return "student"
}

type Student struct {
	Id      int        `gorm:"primaryKey"`
	Email   string     `gorm:"column:email" json:"email"`
	Status  uint       `gorm:"column:status" json:"status"`
	Teacher []*Teacher `gorm:"many2many:student_teacher;foreignKey:Id;joinForeignKey:StudentId;References:Id;joinReferences:TeacherId"`
}

func (m *Student) SelectOne(email string) (result Student, err error) {
	db := DB.Model(m)
	db.Where("Email = ?", email)
	db.Where("status = 1")
	err = db.First(&result).Error
	return
}

func (m *Student) SelectFromRelations(ids []int) (result []Student, err error) {
	db := DB.Model(m)
	db.Joins("Teacher")
	db.Where("Teacher.Id in ?", ids)
	db.Where("Student.Status = ?", 1)
	err = db.Find(&result).Error
	return
}
