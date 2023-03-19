package models

//user role enum

var UserRole = newUserRoleRegistry()

func newUserRoleRegistry() *userRoleRegistry {
	return &userRoleRegistry{
		Administrator: 1,
		Manager:       2,
		EndUser:       3,
	}
}

type userRoleRegistry struct {
	Administrator int16
	Manager       int16
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

var OtpType = newOtpTypeRegistry()

func newOtpTypeRegistry() *otpTypeRegistry {
	return &otpTypeRegistry{
		SignupRequest: 1,
		LoginRequest:  2,
	}
}

type otpTypeRegistry struct {
	SignupRequest int32
	LoginRequest  int32
}

var UploadUrlType = newUploadUrlTypeRegistry()

func newUploadUrlTypeRegistry() *uploadUrlTypeRegistry {
	return &uploadUrlTypeRegistry{
		Default:    1,
		ProfilePic: 2,
	}
}

type uploadUrlTypeRegistry struct {
	Default    int64
	ProfilePic int64
}

var MimeType = newMimeTypeRegistry()

func newMimeTypeRegistry() *mimeTypeRegistry {
	return &mimeTypeRegistry{
		PNG: "image/png",
	}
}

type mimeTypeRegistry struct {
	PNG string
}
