package password

import (
	"bytes"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"github.com/coffee377/autoctl/lib/idp/security/keygen"
	"golang.org/x/crypto/pbkdf2"
	"hash"
	"strings"
)

type pbkdf2PasswordEncoder struct {
	secret             string
	saltGenerator      keygen.BytesKeyGenerator
	iterations         int
	algorithm          string
	hashLength         int
	encodeHashAsBase64 bool
}

const defaultSaltLength = 8
const defaultHashWidth = 256
const defaultIterations = 185000

func (e pbkdf2PasswordEncoder) Encode(rawPassword string) string {
	salt := e.saltGenerator.GenerateKey()
	encoded := e.saltEncode(rawPassword, salt)
	if e.encodeHashAsBase64 {
		return base64.StdEncoding.EncodeToString(encoded)
	}
	return hex.EncodeToString(encoded)
}

func (e pbkdf2PasswordEncoder) Matches(rawPassword string, encodedPassword string) bool {
	digested := e.decode(encodedPassword)
	salt := subArray(digested, 0, e.saltGenerator.GetKeyLength())
	return bytes.Equal(digested, e.saltEncode(rawPassword, salt))
}

func (e pbkdf2PasswordEncoder) UpgradeEncoding(string) bool {
	return false
}

func (e pbkdf2PasswordEncoder) saltEncode(password string, salt []byte) []byte {
	concatenateSalt := concatenate(salt, []byte(e.secret))
	iter := e.iterations
	if e.iterations <= 0 {
		iter = defaultIterations
	}
	keyLen := e.hashLength
	if e.hashLength <= 0 {
		keyLen = defaultHashWidth / 8
	}
	var h func() hash.Hash
	switch strings.ToLower(e.algorithm) {
	case "sha1":
		h = sha1.New
		break
	case "sha256":
		h = sha256.New
		break
	case "sha512":
		h = sha512.New
		break
	default:
		h = sha1.New
	}
	key := pbkdf2.Key([]byte(password), concatenateSalt, iter, keyLen, h)
	return concatenate(salt, key)
}

func (e pbkdf2PasswordEncoder) encode(encoded []byte) string {
	if e.encodeHashAsBase64 {
		return base64.StdEncoding.EncodeToString(encoded)
	}
	return hex.EncodeToString(encoded)
}

func (e pbkdf2PasswordEncoder) decode(encodedBytes string) []byte {
	var (
		res []byte
		err error
	)
	if e.encodeHashAsBase64 {
		res, err = base64.StdEncoding.DecodeString(encodedBytes)
	}
	res, err = hex.DecodeString(encodedBytes)
	if err != nil {
		panic(err)
	}
	return res
}

func concatenate(arrays ...[]byte) []byte {
	res := make([]byte, 0)
	for _, array := range arrays {
		if len(array) > 0 {
			res = append(res, array...)
		}
	}
	return res

}

func subArray(array []byte, beginIndex, endIndex int) []byte {
	n := endIndex - beginIndex
	subarray := make([]byte, n)         // 创建目标数组，长度为n
	copy(subarray, array[beginIndex:n]) // 拷贝源数组的前n个元素到目标数组
	return subarray
}

func Pbkdf2PasswordEncoder(opts ...Pbkdf2Option) Encoder {
	encoder := &pbkdf2PasswordEncoder{
		saltGenerator:      keygen.RandomBytesKeyGenerator(keygen.RandomWithKeyLength(defaultSaltLength)),
		iterations:         defaultIterations,
		hashLength:         defaultHashWidth / 8,
		encodeHashAsBase64: false,
	}
	for _, opt := range opts {
		opt(encoder)
	}
	return encoder
}

type Pbkdf2Option func(pbkdf2 *pbkdf2PasswordEncoder)

func Pbkdf2WithSecret(secret string) Pbkdf2Option {
	return func(pbkdf2 *pbkdf2PasswordEncoder) {
		pbkdf2.secret = secret
	}
}

func Pbkdf2WithSaltGenerator(saltGenerator keygen.BytesKeyGenerator) Pbkdf2Option {
	return func(pbkdf2 *pbkdf2PasswordEncoder) {
		pbkdf2.saltGenerator = saltGenerator
	}
}

func Pbkdf2WithDefaultSaltGenerator(saltLength int) Pbkdf2Option {
	return func(pbkdf2 *pbkdf2PasswordEncoder) {
		pbkdf2.saltGenerator = keygen.RandomBytesKeyGenerator(keygen.RandomWithKeyLength(saltLength))
	}
}

func Pbkdf2WithIterations(iterations int) Pbkdf2Option {
	return func(pbkdf2 *pbkdf2PasswordEncoder) {
		pbkdf2.iterations = iterations
	}
}

func Pbkdf2WithHashLength(hashLength int) Pbkdf2Option {
	return func(pbkdf2 *pbkdf2PasswordEncoder) {
		pbkdf2.hashLength = hashLength
	}
}

func Pbkdf2WithAlgorithm(algorithm string) Pbkdf2Option {
	return func(pbkdf2 *pbkdf2PasswordEncoder) {
		pbkdf2.algorithm = algorithm
	}
}

func Pbkdf2WithEncodeHashAsBase64(encodeHashAsBase64 bool) Pbkdf2Option {
	return func(pbkdf2 *pbkdf2PasswordEncoder) {
		pbkdf2.encodeHashAsBase64 = encodeHashAsBase64
	}
}
