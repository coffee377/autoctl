package dingtalk

import (
	"fmt"
	"strings"
	"time"
)

type overtime struct {
	time        time.Time
	startTime   time.Time
	endTime     time.Time
	daysOffHour float32 // 调休工时
	mealFee     float32 // 餐费
	trafficFee  float32 // 交通费
}

type Overtime interface {
	GetDate() string             // 日期
	GetWeek() int8               // 一个月中的第几周
	GetStartTime() string        // 开始时间
	GetEndTime() string          // 结束时间
	GetActualHour() float32      // 实际工时
	GetDaysOffHourHour() float32 // 调休工时
	GetMealFee() float32         // 餐费
	GetTrafficFee() float32      // 交通费
	GetRemark() string           // 备注
}

func (o overtime) GetDate() string {
	return o.time.Format("2006/01/02")
}

func (o overtime) GetWeek() int8 {
	// 获取这个日期在一个月中的天数
	dayOfMonth := o.time.Day()
	// 获取这个日期在这一周的天数
	dayOfWeek := int(o.time.Weekday())
	// 计算这个日期在一个月中的周数
	weekOfMonth := (dayOfMonth - 1 + dayOfWeek) / 7
	if (dayOfMonth-1+dayOfWeek)%7 != 0 {
		weekOfMonth++
	}
	return int8(weekOfMonth)
}

func (o overtime) GetStartTime() string {
	return o.time.Format("15:04:05")
}

func (o overtime) GetEndTime() string {
	return o.time.Format("15:04:05")
}

func (o overtime) GetActualHour() float32 {
	return 0
}

func (o overtime) GetDaysOffHourHour() float32 {
	return o.daysOffHour
}

func (o overtime) GetMealFee() float32 {
	return o.mealFee
}

func (o overtime) GetTrafficFee() float32 {
	return o.trafficFee
}

func (o overtime) GetRemark() string {
	s := make([]string, 0)
	if o.GetDate() != "" {
		s = append(s, fmt.Sprintf("%s 加班", o.GetDate()))
	}
	if o.mealFee > 0 {
		s = append(s, fmt.Sprintf("餐费：%f", o.mealFee))
	}
	if o.trafficFee > 0 {
		s = append(s, fmt.Sprintf("交通费：%f", o.trafficFee))
	}
	if len(s) > 1 {
		return strings.Join(s, " ")
	}
	return ""
}

type OvertimeOption = func(*overtime)

func NewOvertime(opts ...OvertimeOption) Overtime {
	o := &overtime{
		mealFee:    20,
		trafficFee: 35,
	}
	for _, opt := range opts {
		opt(o)
	}
	return o
}

func WithDate(date string) OvertimeOption {
	return func(o *overtime) {
		parse, _ := time.Parse("2006-01-02", date)
		o.time = parse
	}
}

func WithMealFee(mealFee float32) OvertimeOption {
	return func(o *overtime) {
		o.mealFee = mealFee
	}
}

func WithDaysOffHour(daysOffHour float32) OvertimeOption {
	return func(o *overtime) {
		o.daysOffHour = daysOffHour
	}
}

func WithTrafficFee(trafficFee float32) OvertimeOption {
	return func(o *overtime) {
		o.trafficFee = trafficFee
	}
}
