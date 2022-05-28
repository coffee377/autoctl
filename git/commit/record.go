package git

import (
	"reflect"
	"strconv"
)

var (
	FormatSep    = ";"
	RecordFormat = "%H;%at;%ad;%an;%ae;%B" // 格式化顺序与 CommitRecord 结构体字段顺序一致，便于反射处理
)

type CommitBase struct {
	Commit string `json:"commit,omitempty"` // %H: commit hash
	Time   string `json:"-"`                // %at: author date, UNIX timestamp
	Date   string `json:"date,omitempty"`   // %ad: author date (format respects --date=option)
	Author string `json:"author,omitempty"` // %an: author name
	Email  string `json:"email,omitempty"`  // %ae: author email
	Raw    string `json:"-"`                // %B: 原始
}

// CommitRecord contains the commit information
type CommitRecord struct {
	CommitBase                // 提交记录基础字段
	Timestamp  int            `json:"timestamp,omitempty"` // 根据 Time 进行转换
	Message    *CommitMessage `json:"message,omitempty"`   // 根据 Raw 进行转换
}

type CommitRecordOptions struct {
	RecordFormat string // git log 格式化模板
	FormatSep    string // 模板字符分割字符串
	FormattedLog string // 格式化的日志字符串
	//Verify       MessageVerify
}

func NewCommitRecord(formattedLog string) *CommitRecord {
	r := new(CommitRecord)

	// 1. 模板字段处理
	//for i, v := range strings.Split(formattedLog, FormatSep) {
	//	r.setProps(i, v)
	//}

	message := r.Raw
	// 2. commit message 格式校验
	//if !r.Match(message) {
	//	log.Fatal("commit %s 提交格式不符合规范", r.Commit)
	//}

	// 3. 时间戳类型转换
	r.Timestamp, _ = strconv.Atoi(r.Date)

	// 4. 提交信息类型转换
	r.Message = NewCommitMessage([]byte(message))

	return r
}

type Record interface {
	From(formattedLog string) *CommitRecord
	MessageVerify(message string) bool // 验证格式是否正确
}

type record interface {
	convert(formattedLog string) *CommitBase     // 从格式化 log 回去基本提交信息
	fillAutoField(base CommitBase) *CommitRecord // 填充自动计算字段
}

func (record *CommitRecord) setProps(i int, v string) *CommitRecord {
	reflect.ValueOf(record).Elem().Field(i).Set(reflect.ValueOf(v))
	return record
}

func (record *CommitRecord) convert(formattedLog string) *CommitBase {
	return &CommitBase{}
}

func (record *CommitRecord) fillAutoField(base CommitBase) *CommitRecord {
	record.Timestamp, _ = strconv.Atoi(base.Date)
	record.Message = NewCommitMessage([]byte(base.Raw))
	return record
}

func (record *CommitRecord) Match(message string) bool {
	return false
}
