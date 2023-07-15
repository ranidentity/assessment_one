package service

import (
	"singo/model"
	"singo/serializer"
)

type StudentRegistrationService struct {
	FormTeacher    string   `form:"teacher_email" binding:"required,email"`
	FormStudents   []string `form:"students"  binding:"required"`
	Student        model.Student
	Teacher        model.Teacher
	StudentTeacher model.StudentTeacher
}

func (service *StudentRegistrationService) Register() serializer.Response {
	var error_email []string
	teacher, err := service.Teacher.SelectOne(service.FormTeacher)
	if err != nil {
		return serializer.Response{
			Code: 40001,
			Msg:  "Teacher not found",
		}
	}
	var data []model.StudentTeacherWithEmail
	for _, i := range service.FormStudents {
		student, err := service.Student.SelectOne(i, model.ACTIVE)
		if err != nil {
			error_email = append(error_email, i)
		} else {
			data = append(data, model.StudentTeacherWithEmail{Email: i, StudentId: student.Id, TeacherId: teacher.Id})
		}
	}
	successful_registered_students, err := service.StudentTeacher.TeacherRegisterStudent(data)
	if err != nil {
		return serializer.ParamErr("Registration failed", err)
	}
	return serializer.BuildStudentTeacherResponse(service.FormTeacher, successful_registered_students)
}

// func (service *StudentRegistrationService) valid() *serializer.Response {
// 	var error_email []string
// 	for _, i := range service.FormStudents {
// 		_, err := service.Student.SelectOne(i, model.ACTIVE)s
// 		if err != nil {
// 			error_email = append(error_email, i)
// 		}
// 	}
// 	if len(error_email) > 0 {
// 		return &serializer.Response{
// 			Code: 40001,
// 			Msg:  "Student not found",
// 		}
// 	}
// 	return nil
// }
