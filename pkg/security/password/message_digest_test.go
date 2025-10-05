package password

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testMD5Vectors = []vector{
	{
		"password",
		"{HFGfuhV0LFe2lx9SXNFbtfN5t/tVSnjYDZKeKPNZRYs=}db4da5ebd4a1aa6e9c760bc4df0aabc6",
		true,
	},
	{
		"password",
		"{/e3ykgMSrGERJu4tMje4zPARHEHTCnljyoHCzCjXtI4=}6d426c88d2fc5699d8b21f564a8f8ccf",
		true,
	},
	{
		"password",
		"{CZwERNvXfD1c1le3YETcReOxWYArpEtEUq11kbKVsbk=}2e2b4cdf2e15c4bd825d6ce69a0f8b229fb5a9b9",
		false,
	},
	{
		"password",
		"{CZwERNvXfD1c1le3YETcReOxWYArpEtEUq11kbKVsbk=}2e2b4cdf2e15c4bd825d6ce69a0f8b229fb5a9b9",
		false,
	},
	{
		"password",
		"{ss2fCxARIPfc5FqU/jOsZLuigeccxf5VRY32GOEXFvU=}239c93f284413764abf5f07bea837bfe728cf99d6c60a10e0f749a68aed441a0",
		false,
	},
	{
		"password",
		"{FhbpZ+LB2i0B94txAAhxDznMqlRDfltYLFU9bTkSk/A=}6e2366f26a174698367c94a9e7bb4024971f42983f55a96bd0bdf782fbd19828065d014653ecb5922fd6f58a87b5edbaef931d35831ed7835488019b930525ff",
		false,
	},
}

var testSHA1Vectors = []vector{
	{
		"password",
		"{CZwERNvXfD1c1le3YETcReOxWYArpEtEUq11kbKVsbk=}2e2b4cdf2e15c4bd825d6ce69a0f8b229fb5a9b9",
		true,
	},
	{
		"password",
		"{BoO1hnCnRDAj7TJg9awzvc9ZWKW9pK1O7HdkqfS0yHw=}08752f84845573b3c6358b06bb298b5dc16309bf",
		true,
	},
	{
		"password",
		"{HFGfuhV0LFe2lx9SXNFbtfN5t/tVSnjYDZKeKPNZRYs=}db4da5ebd4a1aa6e9c760bc4df0aabc6",
		false,
	},
	{
		"password",
		"{ss2fCxARIPfc5FqU/jOsZLuigeccxf5VRY32GOEXFvU=}239c93f284413764abf5f07bea837bfe728cf99d6c60a10e0f749a68aed441a0",
		false,
	},
	{
		"password",
		"{FhbpZ+LB2i0B94txAAhxDznMqlRDfltYLFU9bTkSk/A=}6e2366f26a174698367c94a9e7bb4024971f42983f55a96bd0bdf782fbd19828065d014653ecb5922fd6f58a87b5edbaef931d35831ed7835488019b930525ff",
		false,
	},
}

var testSHA256Vectors = []vector{
	{
		"password",
		"{ss2fCxARIPfc5FqU/jOsZLuigeccxf5VRY32GOEXFvU=}239c93f284413764abf5f07bea837bfe728cf99d6c60a10e0f749a68aed441a0",
		true,
	},
	{
		"password",
		"{OPcr7LVGfS+AyW6CT7Iwz82OnkGxg65QyoPeDWscsQ0=}b922310f8a1bffccdc4ba6f96555d6f38f43c433f3ca436f929af37d16384445",
		true,
	},
	{
		"password",
		"{HFGfuhV0LFe2lx9SXNFbtfN5t/tVSnjYDZKeKPNZRYs=}db4da5ebd4a1aa6e9c760bc4df0aabc6",
		false,
	},
	{
		"password",
		"{CZwERNvXfD1c1le3YETcReOxWYArpEtEUq11kbKVsbk=}2e2b4cdf2e15c4bd825d6ce69a0f8b229fb5a9b9",
		false,
	},
	{
		"password",
		"{FhbpZ+LB2i0B94txAAhxDznMqlRDfltYLFU9bTkSk/A=}6e2366f26a174698367c94a9e7bb4024971f42983f55a96bd0bdf782fbd19828065d014653ecb5922fd6f58a87b5edbaef931d35831ed7835488019b930525ff",
		false,
	},
}

var testSHA512Vectors = []vector{
	{
		"password",
		"{FhbpZ+LB2i0B94txAAhxDznMqlRDfltYLFU9bTkSk/A=}6e2366f26a174698367c94a9e7bb4024971f42983f55a96bd0bdf782fbd19828065d014653ecb5922fd6f58a87b5edbaef931d35831ed7835488019b930525ff",
		true,
	},
	{
		"password",
		"{QSaE9Ve5W6PtPE7yosJGAfaTCCAuptwgXpo8JHbw4UI=}2c02a43d80a870d7014d60998c0af2c61c8d4915a4c2901d9cd9abf6f99cf2053a9ad939d07f66727c8812e5ae9f9009347b3da131264bd04b5f8a3f5b0cd081",
		true,
	},
	{
		"password",
		"{HFGfuhV0LFe2lx9SXNFbtfN5t/tVSnjYDZKeKPNZRYs=}db4da5ebd4a1aa6e9c760bc4df0aabc6",
		false,
	},
	{
		"password",
		"{CZwERNvXfD1c1le3YETcReOxWYArpEtEUq11kbKVsbk=}2e2b4cdf2e15c4bd825d6ce69a0f8b229fb5a9b9",
		false,
	},
	{
		"password",
		"{ss2fCxARIPfc5FqU/jOsZLuigeccxf5VRY32GOEXFvU=}239c93f284413764abf5f07bea837bfe728cf99d6c60a10e0f749a68aed441a0",
		false,
	},
}

func testMessageDigestPasswordEncoder(t *testing.T, algorithm string, vectors []vector) {
	for _, v := range vectors {
		encoder := MessageDigestPasswordEncoder(algorithm)
		out := encoder.Encode(v.password)
		match1 := encoder.Matches(v.password, out)
		assert.Equal(t, match1, true)
		match2 := encoder.Matches(v.password, v.encodedPassword)
		assert.Equal(t, match2, v.match)
	}
}

func TestMD5(t *testing.T) {
	testMessageDigestPasswordEncoder(t, "MD5", testMD5Vectors)
}

func TestSHA1(t *testing.T) {
	testMessageDigestPasswordEncoder(t, "SHA-1", testSHA1Vectors)
}

func TestSHA256(t *testing.T) {
	testMessageDigestPasswordEncoder(t, "SHA-256", testSHA256Vectors)
}

func TestSHA512(t *testing.T) {
	testMessageDigestPasswordEncoder(t, "SHA-512", testSHA512Vectors)
}
