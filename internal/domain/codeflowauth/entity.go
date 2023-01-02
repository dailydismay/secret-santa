package codeflowauth

type UserProfileID string

func NewUserProfileID(s string) UserProfileID {
	return UserProfileID(s)
}

type UserProfile struct {
	id        UserProfileID
	firstName string
	lastName  string
	photo     string
}

func (up *UserProfile) GetID() UserProfileID {
	return up.id
}

func (up *UserProfile) GetFirstName() string {
	return up.firstName
}

func (up *UserProfile) GetLastName() string {
	return up.lastName
}

func (up *UserProfile) GetPhoto() string {
	return up.photo
}

func NewUserProfile(
	id UserProfileID,
	firstName string,
	lastName string,
	photo string,
) *UserProfile {
	return &UserProfile{
		id:        id,
		firstName: firstName,
		lastName:  lastName,
		photo:     photo,
	}
}
