package model

func (Teacher) TableName() string {
	return "teacher"
}

type Teacher struct {
	Id      int        `gorm:"primaryKey"`
	Email   string     `gorm:"column:email" json:"email"`
	Student []*Student `gorm:"many2many:student_teacher;foreignKey:Id;joinForeignKey:TeacherId;References:Id;joinReferences:StudentId"`
}

func (m *Teacher) SelectOne(email string) (result Teacher, err error) {
	db := DB.Model(m)
	db.Where("Email = ?", email)
	err = db.First(&result).Error
	return
}

func (m *Teacher) SelectMultiple(email []string) (result []Teacher, err error) {
	db := DB.Model(m)
	db.Where("Email in ?", email)
	// db.Preload("Student")
	err = db.Find(&result).Error
	return
}
