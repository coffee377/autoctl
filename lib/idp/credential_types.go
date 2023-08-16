package idp

import (
	"github.com/ory/x/sqlxx"
	"time"
)

type Key = uint64

// CredentialType 身份凭证类型
type CredentialType struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// Credentials 表示特定的凭证类型
type Credentials struct {
	ID                       uint64 `json:"-" db:"id"`
	IdentityCredentialTypeID uint64 `json:"-" db:"identity_credential_type_id"`
	IdentityID               uint64 `json:"-" faker:"-" db:"identity_id"`

	// Type discriminates between different types of credentials.
	Type CredentialType `json:"type" db:"-"`

	// Identifiers represents a list of unique identifiers this credential type matches.
	Identifiers []string `json:"identifiers" db:"-"`

	// Config contains the concrete credential payload. This might contain the bcrypt-hashed password, the email
	// for passwordless authentication or access_token and refresh tokens from OpenID Connect flows.
	Config sqlxx.JSONRawMessage `json:"config,omitempty" db:"config"`

	// Version refers to the version of the credential. Useful when changing the config schema.
	Version int `json:"version" db:"version"`

	CreatedAt time.Time `json:"created_at" db:"created_at"`

	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	//NID       uuid.UUID `json:"-"  faker:"-" db:"nid"`
}

// CredentialIdentifier 身份凭证标识符
type CredentialIdentifier struct {
	ID         Key    `db:"id"`
	Identifier string `db:"identifier"`
	// IdentityCredentialsID is a helper struct field for gobuffalo.pop.
	IdentityCredentialsID Key `json:"-" db:"identity_credential_id"`
	// IdentityCredentialsTypeID is a helper struct field for gobuffalo.pop.
	IdentityCredentialsTypeID Key `json:"-" db:"identity_credential_type_id"`
	// CreatedAt is a helper struct field for gobuffalo.pop.
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	// UpdatedAt is a helper struct field for gobuffalo.pop.
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	//NID       uuid.UUID `json:"-"  faker:"-" db:"nid"`
}
