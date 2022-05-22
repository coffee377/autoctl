package git

import (
	"strings"
)

var (
	FormatSep       = ";"
	GitRecordFormat = "%H;%h;%an;%ae;%at;%ad;%B"
)

// CommitRecord contains the commit information
type CommitRecord struct {
	Commit      string        `json:"hash"`        // %H: commit hash
	ShortCommit string        `json:"shortCommit"` // %h: abbreviated commit hash
	Author      string        `json:"author"`      // %an: author name
	Email       string        `json:"email"`       // %ae: author email
	Timestamp   int8          `json:"timestamp"`   // %at: author date, UNIX timestamp
	Date        string        `json:"date"`        // %ad: author date (format respects --date=option)
	message     string        ``                   // %B: raw body (unwrapped subject and body)
	Message     CommitMessage `json:"message"`     // wrapper message
}

func NewCommitRecord(formatLog string) *CommitRecord {
	r := new(CommitRecord)
	//types := reflect.TypeOf(r)
	//fieldName := types.Elem().Field(0).Name
	//name := reflect.ValueOf(r).Elem().FieldByName(fieldName)
	//field.Typ
	for i, v := range strings.Split(formatLog, FormatSep) {
		r.setProps(i, v)
	}

	return r
}

func (record *CommitRecord) setProps(i int, v string) *CommitRecord {
	return record
}
