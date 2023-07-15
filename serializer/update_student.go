package serializer

func BuildUpdateStudentResponse(email string) Response {
	return Response{
		Msg: "Suspended student: " + email,
	}
}
