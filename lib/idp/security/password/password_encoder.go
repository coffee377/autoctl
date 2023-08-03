package password

type Encoder interface {
	// Encode 对原始密码进行编码
	// 一般来说，一个好的编码算法会使用SHA-1或更高级别的哈希函数
	// 结合一个随机生成的8字节或更长的盐值
	Encode(rawPassword string) string
	// Matches
	// Verify the encoded password obtained from storage matches the submitted raw
	// password after it too is encoded. Returns true if the passwords match, false if
	// they do not. The stored password itself is never decoded.
	// @param rawPassword the raw password to encode and match
	// @param encodedPassword the encoded password from storage to compare with
	// @return true if the raw password, after encoding, matches the encoded password from
	// storage
	Matches(rawPassword string, encodedPassword string) bool
	// UpgradeEncoding
	// Returns true if the encoded password should be encoded again for better security,
	// else false. The default implementation always returns false.
	// @param encodedPassword the encoded password to check
	// @return true if the encoded password should be encoded again for better security,
	// else false.
	//
	UpgradeEncoding(encodedPassword string) bool
}

// CreateDelegatingPasswordEncoder https://docs.oracle.com/en/java/javase/11/docs/specs/security/standard-names.html#messagedigest-algorithms
func CreateDelegatingPasswordEncoder() Encoder {
	encodingId := "bcrypt"
	encoders := make(map[string]Encoder)
	encoders[encodingId] = BCryptPasswordEncoder()
	encoders["ldap"] = ldapShaPasswordEncoder{}
	encoders["MD4"] = Md4PasswordEncoder()
	encoders["MD5"] = messageDigestPasswordEncoder{algorithm: "MD5"}
	encoders["noop"] = NoopPasswordEncoder()
	encoders["pbkdf2"] = Pbkdf2PasswordEncoder()
	encoders["scrypt"] = scryptPasswordEncoder{}
	encoders["SHA-1"] = messageDigestPasswordEncoder{algorithm: "SHA-1"}
	encoders["SHA-256"] = messageDigestPasswordEncoder{algorithm: "SHA-256"}
	encoders["sha256"] = standardPasswordEncoder{}
	encoders["argon2"] = argon2PasswordEncoder{}
	return DelegatingPasswordEncoder(encodingId, encoders, DelegatingWithIdCaseInsensitive())
}

type EncoderFactories struct {
}

func (EncoderFactories) CreateDelegatingPasswordEncoder2() Encoder {
	encodingId := "bcrypt"
	encoders := make(map[string]Encoder)
	encoders[encodingId] = BCryptPasswordEncoder()
	encoders["ldap"] = ldapShaPasswordEncoder{}
	encoders["MD4"] = Md4PasswordEncoder()
	encoders["MD5"] = messageDigestPasswordEncoder{algorithm: "MD5"}
	encoders["noop"] = NoopPasswordEncoder()
	encoders["pbkdf2"] = Pbkdf2PasswordEncoder()
	encoders["scrypt"] = scryptPasswordEncoder{}
	encoders["SHA-1"] = messageDigestPasswordEncoder{algorithm: "SHA-1"}
	encoders["SHA-256"] = messageDigestPasswordEncoder{algorithm: "SHA-256"}
	encoders["sha256"] = standardPasswordEncoder{}
	encoders["argon2"] = argon2PasswordEncoder{}
	return DelegatingPasswordEncoder(encodingId, encoders)
}
