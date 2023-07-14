package serializer

import "singo/model"

func BuildStudentTeacherResponse(input []model.StudentTeacher, teacher_email string, students []string) Response {
	return Response{
		Data: input,
	}
}
