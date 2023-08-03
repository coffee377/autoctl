package password

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type vector struct {
	password        string
	encodedPassword string
	match           bool
}

var testVectors = []vector{
	{
		"password",
		"{BCRYPT}$2a$10$dXJ3SW6G7P50lGmMkkmwe.20cQQubK3.HZWzG3YB1tlRy.fqvM/BG",
		true,
	},
	{
		"password",
		"{noop}password",
		true,
	},
	{
		"password",
		"{pbkdf2}5d923b44a6d129f3ddf3e3c8d29412723dcbde72445e8ef6bf3b508fbf17fa4ed4d6b99ca763d8dc",
		true,
	},
	//{
	//	"password",
	//	"{scrypt}$e0801$8bWJaSu2IKSn9Z9kM+TPXfOc/9bdYSrN1oD9qfVThWEwdRTnO7re7Ei+fUZRJ68k9lTyuTeUp4of4g24hHnazw==$OAOec05+bXxvuu/1qZ6NUR+xQYvYv7BeL1QxwRpY5Pc=",
	//	true,
	//},
	//{
	//	"password",
	//	"{sha256}97cde38028ad898ebc02e690819fa220e88c62e0699403e94fff291cfffaf8410849f27605abcbc0",
	//	true,
	//},
}

func test(t *testing.T, vectors []vector) {
	for _, v := range vectors {
		encoder := CreateDelegatingPasswordEncoder()
		//out := encoder.Encode(v.password)
		match := encoder.Matches(v.password, v.encodedPassword)
		assert.Equal(t, match, v.match)
	}
}

func TestDelegatingPasswordEncoder(t *testing.T) {
	test(t, testVectors)
}
