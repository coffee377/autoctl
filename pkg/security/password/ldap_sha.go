package password

// Deprecated
type ldapShaPasswordEncoder struct{}

func (e ldapShaPasswordEncoder) Encode(rawPassword string) string {
	//TODO implement me
	panic("implement me")
}

func (e ldapShaPasswordEncoder) Matches(rawPassword string, encodedPassword string) bool {
	//TODO implement me
	panic("implement me")
}

func (e ldapShaPasswordEncoder) UpgradeEncoding(encodedPassword string) bool {
	//TODO implement me
	panic("implement me")
}
