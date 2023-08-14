package password

type IPassword interface {
	GetEncodedPassword() string
	GetCryptoType() string
	GetCryptoSalt() string
}

type Password struct {
	Raw             string
	Types           string
	Salt            string
	EncodedPassword string
}
