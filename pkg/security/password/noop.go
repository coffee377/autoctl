package password

type noopPasswordEncoder struct{}

// NoopPasswordEncoder Deprecated
func NoopPasswordEncoder() Encoder {
	return &noopPasswordEncoder{}
}

func (e noopPasswordEncoder) Encode(rawPassword string) string {
	return rawPassword
}

func (e noopPasswordEncoder) Matches(rawPassword string, encodedPassword string) bool {
	return rawPassword == encodedPassword
}

func (e noopPasswordEncoder) UpgradeEncoding(encodedPassword string) bool {
	return false
}
