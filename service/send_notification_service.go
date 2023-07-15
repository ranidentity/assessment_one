package service

import (
	"singo/model"
	"singo/serializer"
)

type SendNotificationService struct {
	FormTeacher    string `form:"teacher"  binding:"required"`
	FormMessage    string `form:"notification" binding:"required"`
	Teacher        model.Teacher
	StudentTeacher model.StudentTeacher
	Student        model.Student
}

func (service *SendNotificationService) CheckNotificationTarget() serializer.Response {
	teacher, err := service.Teacher.SelectOne(service.FormTeacher)
	if err != nil {
		return serializer.Response{
			Code: 40001,
			Msg:  "Teacher not found",
		}
	}
	emails := service.CaptureMention()

	// Check mention
	mentioned, err := service.Student.SelectFromEmails(emails)

	// TODO verify email format
	if err != nil {
		return serializer.Response{
			Code: 40001,
			Msg:  "Invalid student(s) mentioned",
		}
	}
	valid_emails := service.Student.RetrieveEmails(mentioned)

	// Take all student under teacher
	students, _ := service.Student.GetNotificationTarget(valid_emails, teacher.Id)
	return serializer.BuildRetrieveForNotificationResponse(students)
}

func (service *SendNotificationService) CaptureMention() []string {
	var result []string
	flag := false
	var mention string
	for _i, i := range service.FormMessage {
		this_char := string(i)
		switch flag {
		case true:
			if _i == len(service.FormMessage)-1 {
				mention += this_char
				result = append(result, mention)
				flag = false
			} else if this_char == " " {
				result = append(result, mention)
				flag = false
			} else {
				mention += this_char
			}
		case false:
			mention = ""
			if this_char == "@" {
				flag = true
			}
		}
	}
	return result
}
