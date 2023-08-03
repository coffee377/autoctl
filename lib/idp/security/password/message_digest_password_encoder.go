package password

// Deprecated
type messageDigestPasswordEncoder struct {
	algorithm string
}

func (e messageDigestPasswordEncoder) Encode(rawPassword string) string {
	//TODO implement me
	//md5 := md5.New()
	//sha1 := sha1.New()
	//h := sha256.New()
	//h := sha512.New()
	panic("implement me")
}

func (e messageDigestPasswordEncoder) Matches(rawPassword string, encodedPassword string) bool {
	//TODO implement me
	panic("implement me")
}

func (e messageDigestPasswordEncoder) UpgradeEncoding(encodedPassword string) bool {
	//TODO implement me
	panic("implement me")
}
