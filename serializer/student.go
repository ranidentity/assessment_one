package serializer

import "singo/model"

func BuildStudentListResponse(input []model.Student) Response {
	return Response{
		Data: input,
	}
}
