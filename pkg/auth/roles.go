package auth

type Role struct {
	ID          int
	Description string
}

// Roles
var (
	RoleViewChats = &Role{
		ID: 1,
		Description: "Allows user to view server chats",
	}
	RoleMutePlayers = &Role{
		ID: 2,
		Description: "Allows user to mute players",
	}
	RoleBanPlayers = &Role{
		ID: 3,
		Description: "Allows user to ban players",
	}
	RoleExpungeHistory = &Role{
		ID: 4,
		Description: "Allows user to expunge items from players' moderation history",
	}
	RoleManageUsers = &Role {
		ID: 5,
		Description: "Allows user to add and remove panel users",
	}
)