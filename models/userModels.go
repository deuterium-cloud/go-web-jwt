package models

type UserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type User struct {
	Username     string
	HashPassword string
	Roles        []string
}

var Users = []User{
	{
		Username:     "nuke",
		HashPassword: "$2a$10$XxU19Rn4n2bNHfEGWKCbzu8pV0JLVFd1GMpGx0ytZIK4F3rJvrWdS",
		Roles:        []string{"USER"},
	},
	{
		Username:     "deuterium",
		HashPassword: "$$2a$10$2ZlrUv4w/30wGAElPbpa7uttukn/Ux.nTJDTWW7DCSOXQenMxlf4i",
		Roles:        []string{"USER"},
	},
	{
		Username:     "milan",
		HashPassword: "$2a$10$6p7kW/hnVypG/lVgUxVXau9DbROyZU0b.8.scC66JGQWfJerEbNOS",
		Roles:        []string{"USER", "ADMIN"},
	},
}
