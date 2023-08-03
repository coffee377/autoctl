package password

type scryptPasswordEncoder struct{}

func (e scryptPasswordEncoder) Encode(rawPassword string) string {
	//TODO implement me
	panic("implement me")
}

func (e scryptPasswordEncoder) Matches(rawPassword string, encodedPassword string) bool {
	//TODO implement me
	panic("implement me")
}

func (e scryptPasswordEncoder) UpgradeEncoding(encodedPassword string) bool {
	//TODO implement me
	panic("implement me")
}
