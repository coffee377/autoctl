package oa

import (
	"errors"
	"fmt"
	"time"
)

// TimeSegment 表示一个时间分段
type TimeSegment struct {
	Start time.Time
	End   time.Time
}

// ProcessTimeRange 处理时间范围并进行分段
// startTimeStr 和 endTimeStr 应为 "yyyy-MM-dd" 格式
// 返回按规则划分的时间分段列表
func ProcessTimeRange(startTimeStr, endTimeStr string) ([]TimeSegment, error) {
	// 定义时间解析格式
	layout := time.DateOnly

	var (
		startTime time.Time
		endTime   time.Time
		err       error
	)

	// 解析开始时间
	startTime, err = time.ParseInLocation(layout, startTimeStr, time.Local)
	if err != nil {
		return nil, fmt.Errorf("解析开始时间失败: %w", err)
	}

	// 解析结束时间
	if endTimeStr != "" {
		endTime, err = time.ParseInLocation(layout, endTimeStr, time.Local)
		if err != nil {
			return nil, fmt.Errorf("解析结束时间失败: %w", err)
		}
	} else {
		endTime = time.Now()
	}

	// 确保开始时间不晚于结束时间
	if startTime.After(endTime) {
		return nil, errors.New("开始时间不能晚于结束时间")
	}

	// 调整结束时间为所在时区当天的最后一刻 (23:59:59.999999999)
	endTime = time.Date(endTime.Year(), endTime.Month(), endTime.Day(), 23, 59, 59, 999999999, time.Local)

	// 计算两个时间之间的天数差
	diffDays := int(endTime.Sub(startTime).Hours() / 24)

	// 如果天数差不超过120天，直接返回一个分段
	if diffDays <= 120 {
		return []TimeSegment{
			{
				Start: startTime,
				End:   endTime,
			},
		}, nil
	}

	// 否则按每3个月进行分段
	var segments []TimeSegment
	currentStart := startTime

	for {
		// 计算3个月后的日期
		currentEnd := addMonths(currentStart, 3, time.Local)

		// 如果3个月后超过了总结束时间，则使用总结束时间
		if currentEnd.After(endTime) {
			currentEnd = endTime
		}

		// 添加当前分段
		segments = append(segments, TimeSegment{
			Start: currentStart,
			End:   currentEnd,
		})

		// 如果当前分段的结束时间已经是总结束时间，则退出循环
		if currentEnd.Equal(endTime) {
			break
		}

		// 下一段的开始时间是当前段结束时间的第二天0点
		currentStart = currentEnd.AddDate(0, 0, 1)
		currentStart = time.Date(currentStart.Year(), currentStart.Month(), currentStart.Day(), 0, 0, 0, 0, time.Local)
	}

	return segments, nil
}

// addMonths 为指定时间添加指定月数，并返回该月最后一天的结束时间
func addMonths(t time.Time, months int, loc *time.Location) time.Time {
	// 计算新的月份和年份
	newMonth := int(t.Month()) + months - 1
	newYear := t.Year()

	// 处理月份进位
	for newMonth > 12 {
		newMonth -= 12
		newYear++
	}

	// 计算该月的最后一天
	lastDay := 31
	for {
		testDate := time.Date(newYear, time.Month(newMonth), lastDay, 0, 0, 0, 0, loc)
		if testDate.Month() == time.Month(newMonth) {
			break
		}
		lastDay--
	}

	// 返回该月最后一天的结束时间
	return time.Date(newYear, time.Month(newMonth), lastDay, 23, 59, 59, 999999999, loc)
}
