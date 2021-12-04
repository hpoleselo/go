package schemas

type PayloadSchema struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int    `json:"age" validate:"required"`
}