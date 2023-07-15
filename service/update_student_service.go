package service

import (
	"singo/model"
	"singo/serializer"
)

type UpdateStudentService struct {
	FormStudent string `form:"student"  binding:"required"`
	Student     model.Student
}

// TODO only can update own student
func (service *UpdateStudentService) UpdateStudent() serializer.Response {
	_, err := service.Student.UpdateStatus(service.FormStudent, model.SUSPENDED)
	if err != nil {
		return serializer.Response{
			Code:  40001,
			Msg:   "Fail to update student: " + service.FormStudent,
			Error: err.Error(),
		}
	}
	return serializer.BuildUpdateStudentResponse(service.FormStudent)
}
