package serializer

type TeacherRegisterStudentResponse struct {
	TeacherEmail string   `json:"teacher"`
	StudentEmail []string `json:"students"`
}

func BuildStudentTeacherResponse(teacher string, students []string) Response {
	return Response{
		Data: TeacherRegisterStudentResponse{
			TeacherEmail: teacher,
			StudentEmail: students,
		},
	}
}
