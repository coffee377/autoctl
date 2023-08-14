package password

import (
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
	"testing"
	"time"
)

func TestBCryptPasswordEncoder(t *testing.T) {

	testCases := []struct {
		name            string
		rawPassword     string
		encodedPassword string
		strength        int
		match           bool
		upgrade         bool
	}{
		{"BCrypt", "123456", "$2a$10$p4uNvJWWYV35ube8WNuBp.QiZQZLYOtYOXvCvBAawWsLwdzHK3nOq", 0, true, false},
		{"BCrypt", "1234567", "$2a$10$p4uNvJWWYV35ube8WNuBp.QiZQZLYOtYOXvCvBAawWsLwdzHK3nOq", 0, false, false},
		{"BCrypt", "123456", "$2a$11$0fIjgSPB4WAqSpCCyDlWuOs0A/0k19n7XffaDjoztwYvIQcd5ivQy", 11, true, false},
		{"BCrypt", "1234567", "$2a$11$0fIjgSPB4WAqSpCCyDlWuOs0A/0k19n7XffaDjoztwYvIQcd5ivQy", 11, false, false},
		{"BCrypt", "123456", "$2a$12$WpsJZhOPOVT3Fif9KT.DDOSxbRu45wAJ2DQZu2o16WnOFKdoQpArq", 12, true, false},
		{"BCrypt", "1234567", "$2a$12$WpsJZhOPOVT3Fif9KT.DDOSxbRu45wAJ2DQZu2o16WnOFKdoQpArq", 12, false, false},

		{"BCrypt", "123456", "$2a$10$p4uNvJWWYV35ube8WNuBp.QiZQZLYOtYOXvCvBAawWsLwdzHK3nOq", 14, true, true},
	}

	for _, tc := range testCases {
		encoder := BCryptPasswordEncoder(BCryptWithStrength(tc.strength))

		t.Run(tc.name, func(t *testing.T) {
			encoded := encoder.Encode(tc.rawPassword)

			// 1. 每次加密后内容都不一样
			assert.NotEqual(t, encoded, tc.encodedPassword, "PasswordEncoder => Encode(%s) = %s; expected %s", tc.rawPassword, encoded, tc.encodedPassword)

			// 2. 密码是否匹配
			matches := encoder.Matches(tc.rawPassword, tc.encodedPassword)
			assert.Equal(t, matches, tc.match, "PasswordEncoder => Matches(%s,%s) = %v; expected %v", tc.rawPassword, tc.encodedPassword, matches, tc.match)

			// 3. 密码是否需要升级
			upgradeEncoding := encoder.UpgradeEncoding(tc.encodedPassword)
			assert.Equal(t, upgradeEncoding, tc.upgrade, "PasswordEncoder => UpgradeEncoding(%s) = %v; expected %v", tc.encodedPassword, upgradeEncoding, tc.upgrade)

		})
	}
}

func TestBcryptPasswordEncoder_Encode(t *testing.T) {
	for i := 0; i < 16; i++ {
		encoder := BCryptPasswordEncoder(BCryptWithStrength(i))
		start := time.Now()
		encode := encoder.Encode("WuYujie890927")
		duration := time.Since(start)
		cost, _ := bcrypt.Cost([]byte(encode))
		println("设置密码强度：", i, "实际密码强度：", cost, "耗时：", duration.String())
		//println(fmt.Sprintf("设置密码强度：%v\t\t实际密码强度：%v\t\t耗时：%v", i, cost, duration.Nanoseconds()))
	}

}
