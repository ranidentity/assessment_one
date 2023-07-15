package service

import (
	"singo/model"
	"singo/serializer"
)

type RetrieveStudentService struct {
	FormTeachers   []string `form:"teacher"  binding:"required"`
	Teacher        model.Teacher
	StudentTeacher model.StudentTeacher
	Student        model.Student
}

/*
alternative: sql query joins
*/
func (service *RetrieveStudentService) RetrieveStudent() serializer.Response {
	teachers, err := service.Teacher.SelectMultiple(service.FormTeachers)
	if err != nil {
		return serializer.Response{
			Code: 40001,
			Msg:  "Teacher not found",
		}
	}
	var teacher_ids []int
	for _, i := range teachers {
		teacher_ids = append(teacher_ids, i.Id)
	}
	students, err := service.Student.SelectFromRelations(teacher_ids)
	if err != nil {
		return serializer.Response{
			Code: 40001,
			Msg:  "Record not found",
		}
	}
	var common []string
	for _, i := range students {
		if len(i.StudentTeacher) == len(teacher_ids) {
			common = append(common, i.Email)
		}
	}

	return serializer.BuildCommonStudentListResponse(common)
}
