package password

import (
	"fmt"
	"strings"
)

const defaultIdPrefix = "{"
const defaultIdSuffix = "}"

type delegatingPasswordEncoder struct {
	idPrefix                         string
	idSuffix                         string
	idForEncode                      string
	idCaseSensitive                  bool
	passwordEncoderForEncode         Encoder
	idToPasswordEncoder              map[string]Encoder
	defaultPasswordEncoderForMatches Encoder
}

func (e delegatingPasswordEncoder) Encode(rawPassword string) string {
	encode := e.passwordEncoderForEncode.Encode(rawPassword)
	return fmt.Sprintf("%s%s%s%s", e.idPrefix, e.idForEncode, e.idSuffix, encode)
}

func (e delegatingPasswordEncoder) getDelegateEncoder(id string) Encoder {
	if e.idCaseSensitive {
		return e.idToPasswordEncoder[id]
	}
	return e.idToPasswordEncoder[strings.ToLower(id)]
}

func (e delegatingPasswordEncoder) Matches(rawPassword string, prefixEncodedPassword string) bool {
	if rawPassword == "" && prefixEncodedPassword == "" {
		return true
	} else {
		id := extractId(prefixEncodedPassword, e.idPrefix, e.idSuffix)
		delegate := e.getDelegateEncoder(id)
		if delegate != nil {
			encodedPassword := extractEncodedPassword(prefixEncodedPassword, e.idPrefix, e.idSuffix)
			return delegate.Matches(rawPassword, encodedPassword)
		}
		return e.defaultPasswordEncoderForMatches.Matches(rawPassword, prefixEncodedPassword)
	}
}

func (e delegatingPasswordEncoder) UpgradeEncoding(prefixEncodedPassword string) bool {
	id := extractId(prefixEncodedPassword, e.idPrefix, e.idSuffix)
	if !strings.EqualFold(e.idForEncode, id) {
		return true
	}
	encodedPassword := extractEncodedPassword(prefixEncodedPassword, e.idPrefix, e.idSuffix)
	delegateEncoder := e.getDelegateEncoder(id)
	if delegateEncoder != nil {
		return delegateEncoder.UpgradeEncoding(encodedPassword)
	}
	return false
}

func extractId(prefixEncodedPassword, idPrefix, idSuffix string) string {
	if prefixEncodedPassword != "" {
		start := strings.Index(prefixEncodedPassword, idPrefix)
		end := strings.Index(prefixEncodedPassword, idSuffix)
		if start == 0 && end > 0 {
			return prefixEncodedPassword[len(idPrefix):end]
		}
	}
	return ""
}

func extractEncodedPassword(prefixEncodedPassword string, idPrefix, idSuffix string) string {
	id := extractId(prefixEncodedPassword, idPrefix, idSuffix)
	prefix := fmt.Sprintf("%s%s%s", idPrefix, id, idSuffix)
	return strings.TrimPrefix(prefixEncodedPassword, prefix)
	//end := strings.Index(prefixEncodedPassword, idSuffix)
	//return prefixEncodedPassword[end+len(idSuffix):]
}

type unmappedIdPasswordEncoder struct {
	idPrefix string
	idSuffix string
}

func (e unmappedIdPasswordEncoder) Encode(rawPassword string) string {
	panic("encode is not supported")
}

func (e unmappedIdPasswordEncoder) Matches(rawPassword string, prefixEncodedPassword string) bool {
	id := extractId(prefixEncodedPassword, e.idPrefix, e.idSuffix)
	panic(fmt.Sprintf("There is no Encoder mapped for the id \"%s\"", id))
}

func (e unmappedIdPasswordEncoder) UpgradeEncoding(encodedPassword string) bool {
	return false
}

type DelegatingPasswordEncoderOption func(d *delegatingPasswordEncoder)

func DelegatingPasswordEncoder(idForEncode string, idToPasswordEncoder map[string]Encoder, opts ...DelegatingPasswordEncoderOption) Encoder {
	encoder := &delegatingPasswordEncoder{
		idPrefix:                         defaultIdPrefix,
		idSuffix:                         defaultIdSuffix,
		idForEncode:                      idForEncode,
		idCaseSensitive:                  true,
		passwordEncoderForEncode:         idToPasswordEncoder[idForEncode],
		idToPasswordEncoder:              idToPasswordEncoder,
		defaultPasswordEncoderForMatches: unmappedIdPasswordEncoder{idPrefix: defaultIdPrefix, idSuffix: defaultIdSuffix},
	}
	for _, opt := range opts {
		opt(encoder)
	}
	if !encoder.idCaseSensitive {
		caseInsensitivePasswordEncoder := make(map[string]Encoder)
		for k, v := range encoder.idToPasswordEncoder {
			caseInsensitivePasswordEncoder[strings.ToLower(k)] = v
		}
		encoder.idToPasswordEncoder = caseInsensitivePasswordEncoder
	}
	return encoder
}

func DelegatingWithIdPrefix(prefix string) DelegatingPasswordEncoderOption {
	return func(d *delegatingPasswordEncoder) {
		if prefix != "" {
			d.idPrefix = prefix
			unmapped := (d.defaultPasswordEncoderForMatches).(unmappedIdPasswordEncoder)
			unmapped.idPrefix = prefix
		}
	}
}

func DelegatingWithIdSuffix(suffix string) DelegatingPasswordEncoderOption {
	return func(d *delegatingPasswordEncoder) {
		if suffix != "" {
			d.idSuffix = suffix
			unmapped := (d.defaultPasswordEncoderForMatches).(unmappedIdPasswordEncoder)
			unmapped.idSuffix = suffix
		}
	}
}

func DelegatingWithId(id string) DelegatingPasswordEncoderOption {
	return func(d *delegatingPasswordEncoder) {
		if id != "" {
			d.idForEncode = id
			d.passwordEncoderForEncode = d.idToPasswordEncoder[id]
		}
	}
}

func DelegatingWithIdCaseInsensitive() DelegatingPasswordEncoderOption {
	return func(d *delegatingPasswordEncoder) {
		d.idCaseSensitive = false
	}
}

func DelegatingWithEncoders(idToPasswordEncoder map[string]Encoder) DelegatingPasswordEncoderOption {
	return func(d *delegatingPasswordEncoder) {
		if idToPasswordEncoder != nil {
			d.idToPasswordEncoder = idToPasswordEncoder
			d.passwordEncoderForEncode = d.idToPasswordEncoder[d.idForEncode]
		}
	}
}
