package password

// Deprecated
type standardPasswordEncoder struct{}

func (e standardPasswordEncoder) Encode(rawPassword string) string {
	//TODO implement me
	panic("implement me")
}

func (e standardPasswordEncoder) Matches(rawPassword string, encodedPassword string) bool {
	//TODO implement me
	panic("implement me")
}

func (e standardPasswordEncoder) UpgradeEncoding(encodedPassword string) bool {
	//TODO implement me
	panic("implement me")
}
