package serializer

type StudentEmailWrapper struct {
	Students []string `json:"students"`
}

func BuildCommonStudentListResponse(input []string) Response {
	result := StudentEmailWrapper{
		Students: input,
	}
	return Response{
		Data: result,
	}
}
