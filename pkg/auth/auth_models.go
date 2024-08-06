package auth

type Role struct {
	ID          int64
	Description string
}

var (
	RoleViewChats = &Role{
		ID: 1,
		Description: "Allows user to view server chats",
	}
)

type User struct {
	Id int64 
	Username string
}

