package rules

type Role int

const (
	RoleVillager Role = iota
	RoleWerewolf
	RoleSeer
	RoleRobber
	RoleTroublemaker
)

type Color struct {
	R int `json:"r"`
	G int `json:"g"`
	B int `json:"b"`
}

type Player struct {
	Name    string `json:"name"`
	Color   Color  `json:"color"`
	Role    Role   `json:"role"`
	Creator bool   `json:"isCreator"`
}
