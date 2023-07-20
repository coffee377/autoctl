package semver

import (
	"regexp"
	"strconv"
	"strings"
)

type Identifier struct {
	Raw       string // 原始字符串
	IsNumeric bool   // 是否数字类型
	Num       uint64 // 数字类型时的具体值
}

func (identifier Identifier) Compare(other Identifier) int {
	if identifier.IsNumeric && !other.IsNumeric {
		return -1
	} else if !identifier.IsNumeric && other.IsNumeric {
		return 1
	} else if identifier.IsNumeric && other.IsNumeric {
		if identifier.Num < other.Num {
			return -1
		} else if identifier.Num > other.Num {
			return 1
		} else {
			return 0
		}
	} else {
		return strings.Compare(identifier.Raw, other.Raw)
	}
}

func NewIdentifier(identifier string) Identifier {
	res := Identifier{
		Raw: identifier,
	}
	result, _ := regexp.MatchString(NumberIdentifierReg, identifier)
	res.IsNumeric = result
	if result {
		res.Num, _ = strconv.ParseUint(identifier, 10, 64)
	}
	return res
}

func parseIdentifiers(str string) []Identifier {
	splits := strings.Split(str, ".")
	identifiers := make([]Identifier, 0, len(splits))
	for _, split := range splits {
		identifiers = append(identifiers, NewIdentifier(split))
	}
	return identifiers
}
