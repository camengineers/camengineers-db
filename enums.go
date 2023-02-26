package models

//user role enum

var UserRole = newUserRoleRegistry()

func newUserRoleRegistry() *userRoleRegistry {
	return &userRoleRegistry{
		Administrator: 1,
		Manufacturer:  2,
		EndUser:       3,
	}
}

type userRoleRegistry struct {
	Administrator int16
	Manufacturer  int16
	EndUser       int16
}

// platform enum

var Platform = newPlatformRegistry()

func newPlatformRegistry() *platformRegistry {
	return &platformRegistry{
		Web:     1,
		iOS:     2,
		Android: 3,
	}
}

type platformRegistry struct {
	Web     int16
	iOS     int16
	Android int16
}
