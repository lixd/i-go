package property

type UserType string

const (
	UserTypeSystem UserType = "system"
	UserTypeUser   UserType = "register"
)

func (r UserType) String() string {
	return string(r)
}

// Values CredentialType list valid values for Enum.
func (UserType) Values() (kinds []string) {
	for _, s := range []UserType{UserTypeSystem, UserTypeUser} {
		kinds = append(kinds, string(s))
	}
	return
}
