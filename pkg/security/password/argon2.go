package password

import (
	"github.com/coffee377/autoctl/pkg/security/keygen"
	"golang.org/x/crypto/argon2"
)

const (
	defaultArgon2SaltLength  = 16
	defaultArgon2HashLength  = 32
	defaultArgon2Parallelism = 1
	defaultArgon2Memory      = 1 << 12
	defaultArgon2Iterations  = 3
)

type argon2PasswordEncoder struct {
	hashLength    uint32
	parallelism   int
	memory        uint32
	iterations    int
	saltGenerator keygen.BytesKeyGenerator
}

func (e argon2PasswordEncoder) Encode(rawPassword string) string {
	salt := e.saltGenerator.GenerateKey()
	argon2.IDKey([]byte(rawPassword), salt, 0, e.memory, 4, e.hashLength)

	//byte[] hash = new byte[this.hashLength];
	//// @formatter:off
	//Argon2Parameters params = new Argon2Parameters
	//.Builder(Argon2Parameters.ARGON2_id)
	//.withSalt(salt)
	//.withParallelism(this.parallelism)
	//.withMemoryAsKB(this.memory)
	//.withIterations(this.iterations)
	//.build();
	//// @formatter:on
	//Argon2BytesGenerator generator = new Argon2BytesGenerator();
	//generator.init(params);
	//generator.generateBytes(rawPassword.toString().toCharArray(), hash);
	//return Argon2EncodingUtils.encode(hash, params);
	return ""
}

func (e argon2PasswordEncoder) Matches(rawPassword string, encodedPassword string) bool {
	//if (encodedPassword == null) {
	//	this.logger.warn("password hash is null");
	//	return false;
	//}
	//Argon2EncodingUtils.Argon2Hash decoded;
	//try {
	//	decoded = Argon2EncodingUtils.decode(encodedPassword);
	//}
	//catch (IllegalArgumentException ex) {
	//	this.logger.warn("Malformed password hash", ex);
	//	return false;
	//}
	//byte[] hashBytes = new byte[decoded.getHash().length];
	//Argon2BytesGenerator generator = new Argon2BytesGenerator();
	//generator.init(decoded.getParameters());
	//generator.generateBytes(rawPassword.toString().toCharArray(), hashBytes);
	//return constantTimeArrayEquals(decoded.getHash(), hashBytes);
	return false
}

func (e argon2PasswordEncoder) UpgradeEncoding(encodedPassword string) bool {
	//if (encodedPassword == null || encodedPassword.length() == 0) {
	//	this.logger.warn("password hash is null");
	//	return false;
	//}
	//Argon2Parameters parameters = Argon2EncodingUtils.decode(encodedPassword).getParameters();
	//return parameters.getMemory() < this.memory || parameters.getIterations() < this.iterations;
	return false
}

type Argon2PasswordEncoderOption func(encoder *argon2PasswordEncoder)

func Argon2PasswordEncoder(opts ...Argon2PasswordEncoderOption) Encoder {
	encoder := &argon2PasswordEncoder{
		hashLength:    defaultArgon2HashLength,
		parallelism:   defaultArgon2Parallelism,
		memory:        defaultArgon2Memory,
		iterations:    defaultArgon2Iterations,
		saltGenerator: keygen.RandomBytesKeyGenerator(keygen.RandomWithKeyLength(defaultArgon2SaltLength)),
	}
	return encoder
}

func Argon2WithSaltLength(length int) Argon2PasswordEncoderOption {
	return func(encoder *argon2PasswordEncoder) {
		encoder.saltGenerator.SetKeyLength(length)
	}
}

func Argon2WithHashLength(length uint32) Argon2PasswordEncoderOption {
	return func(encoder *argon2PasswordEncoder) {
		encoder.hashLength = length
	}
}

func Argon2WithParallelism(parallelism int) Argon2PasswordEncoderOption {
	return func(encoder *argon2PasswordEncoder) {
		encoder.parallelism = parallelism
	}
}

func Argon2WithIterations(iterations int) Argon2PasswordEncoderOption {
	return func(encoder *argon2PasswordEncoder) {
		encoder.iterations = iterations
	}
}
