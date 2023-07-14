package service

import (
	"singo/model"
	"singo/serializer"
)

type RetrieveStudentService struct {
	FormTeachers   []string `form:"teachers[]"  binding:"required"`
	Teacher        model.Teacher
	StudentTeacher model.StudentTeacher
	Student        model.Student
}

/*
alternative: sql query joins
REDO - when 2 or more teachers
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
	// result, err := service.StudentTeacher.List(teacher_ids, "student")
	result, err := service.Student.SelectFromRelations(teacher_ids)
	if err != nil {
		return serializer.Response{
			Code: 40001,
			Msg:  "Record not found",
		}
	}
	return serializer.BuildStudentListResponse(result)
}
