package password

type argon2PasswordEncoder struct{}

func (e argon2PasswordEncoder) Encode(rawPassword string) string {
	return ""
}

func (e argon2PasswordEncoder) Matches(rawPassword string, encodedPassword string) bool {
	return false
}

func (e argon2PasswordEncoder) UpgradeEncoding(encodedPassword string) bool {
	return false
}
