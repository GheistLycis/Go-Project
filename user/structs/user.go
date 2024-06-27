package user_structs

type Gender string

const (
	Male   Gender = "M"
	Female Gender = "F"
	Other  Gender = "O"
)

type User struct {
	CreateUser
	ID int
}
