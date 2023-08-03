package password

import (
	"golang.org/x/crypto/bcrypt"
)

type bcryptPasswordEncoder struct {
	strength int // 密码强度 [4,31],建议不要超过15,否则会很耗时
}

func (e bcryptPasswordEncoder) Encode(rawPassword string) string {
	if rawPassword == "" {
		panic("rawPassword cannot be null")
	} else {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(rawPassword), e.strength)
		if err != nil {
			panic(err.Error())
		}
		return string(hashedPassword)
	}
}

func (e bcryptPasswordEncoder) Matches(rawPassword string, encodedPassword string) bool {
	if rawPassword == "" {
		panic("rawPassword cannot be null")
	} else if encodedPassword != "" {
		if err := bcrypt.CompareHashAndPassword([]byte(encodedPassword), []byte(rawPassword)); err == nil {
			return true
		}
		return false
	}
	return false
}

func (e bcryptPasswordEncoder) UpgradeEncoding(encodedPassword string) bool {
	if encodedPassword != "" {
		cost, err := bcrypt.Cost([]byte(encodedPassword))
		if err != nil {
			panic(err)
		}
		return cost < e.strength
	}
	return false
}

type BCryptPasswordEncoderOption func(bcrypt *bcryptPasswordEncoder)

func BCryptPasswordEncoder(opts ...BCryptPasswordEncoderOption) Encoder {
	encoder := &bcryptPasswordEncoder{
		strength: 10,
	}
	for _, opt := range opts {
		opt(encoder)
	}
	return encoder
}

func BCryptWithStrength(strength int) BCryptPasswordEncoderOption {
	return func(bcrypt *bcryptPasswordEncoder) {
		bcrypt.strength = strength
		if strength < 4 {
			bcrypt.strength = 10
		}
		if strength > 31 {
			bcrypt.strength = 31
		}
	}
}
