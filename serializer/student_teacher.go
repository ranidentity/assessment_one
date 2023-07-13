package serializer

import "singo/model"

func BuildStudentTeacherResponse(input []model.StudentTeacher) Response {
	return Response{
		Data: input,
	}
}
