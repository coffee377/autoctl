package git

import (
	"bytes"
	"reflect"
	"strconv"
)

var (
	FormatSep    = ";"
	LogSep       = "\n!!LF!!\n"
	RecordFormat = "%H;%at;%ad;%an;%ae;%B!!LF!!" // 格式化顺序与 CommitRecord 结构体字段顺序一致，便于反射处理
)

// CommitRecord contains the commit information
type CommitRecord struct {
	Commit     string `json:"commit,omitempty"` // %H: commit hash
	Time       string `json:"-"`                // %at: author date, UNIX timestamp
	Date       string `json:"date,omitempty"`   // %ad: author date (format respects --date=option)
	Author     string `json:"author,omitempty"` // %an: author name
	Email      string `json:"email,omitempty"`  // %ae: author email
	RawMessage string `json:"-"`                // %B: raw commit message

	Timestamp int            `json:"timestamp,omitempty"` // %at: author date, UNIX timestamp
	Message   *CommitMessage `json:"message,omitempty"`   // 根据 RawMessage 进行转换
}

type CommitRecordOptions struct {
	RecordFormat string // git log 格式化模板
	FormatSep    string // 模板字符分割字符串
	FormattedLog string // 格式化的日志字符串
	//Verify       MessageVerify
}

func NewCommitRecord(formattedLog []byte) *CommitRecord {
	r := new(CommitRecord)
	r.From(formattedLog)
	return r
}

type Record interface {
	From(formattedLog []byte) *CommitRecord // 格式化日志
	//MessageVerify(message string) bool      // 验证格式是否正确
}

type record interface {
	//convert(formattedLog string) *CommitBase     // 从格式化 log 回去基本提交信息
	fillAutoField() *CommitRecord // 填充自动计算字段
}

func (record *CommitRecord) setProps(i int, v string) *CommitRecord {
	reflect.ValueOf(record).Elem().Field(i).Set(reflect.ValueOf(v))
	return record
}

func (record *CommitRecord) From(formattedLog []byte) *CommitRecord {

	// 2. commit message 格式校验
	//if !r.Match(message) {
	//	log.Fatal("commit %s 提交格式不符合规范", r.Commit)
	//}
	for i, v := range bytes.Split(formattedLog, []byte(FormatSep)) {
		record.setProps(i, string(v))
	}
	// 3. 时间戳类型转换
	record.fillAutoField()
	return record
}

func (record *CommitRecord) fillAutoField() *CommitRecord {
	record.Timestamp, _ = strconv.Atoi(record.Time)
	record.Message = NewCommitMessage(&record.RawMessage)
	return record
}

func (record *CommitRecord) Match(message string) bool {
	return false
}
