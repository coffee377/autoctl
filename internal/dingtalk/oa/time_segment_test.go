package oa

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProcessTimeRange_1(t *testing.T) {
	timeRange1, err1 := ProcessTimeRange("2025-01-01", "2025-01-12")
	assert.Nil(t, err1)
	assert.Equal(t, 1, len(timeRange1))
	assert.Equal(t, int64(1735660800000), timeRange1[0].Start.UnixMilli())
	assert.Equal(t, int64(1736697599999), timeRange1[0].End.UnixMilli())
}

func TestProcessTimeRange_2(t *testing.T) {
	timeRange1, err1 := ProcessTimeRange("2025-01-01", "2025-06-30")
	assert.Nil(t, err1)
	assert.Equal(t, 2, len(timeRange1))

	assert.Equal(t, int64(1735660800000), timeRange1[0].Start.UnixMilli())
	assert.Equal(t, int64(1743436799999), timeRange1[0].End.UnixMilli())

	assert.Equal(t, int64(1743436800000), timeRange1[1].Start.UnixMilli())
	assert.Equal(t, int64(1751299199999), timeRange1[1].End.UnixMilli())
}

func TestProcessTimeRange_3(t *testing.T) {
	timeRange1, err1 := ProcessTimeRange("2025-01-01", "2025-10-16")
	assert.Nil(t, err1)
	assert.Equal(t, 4, len(timeRange1))

	// 1季度 01-01 ~ 03-31
	assert.Equal(t, int64(1735660800000), timeRange1[0].Start.UnixMilli())
	assert.Equal(t, int64(1743436799999), timeRange1[0].End.UnixMilli())

	// 2季度 04-01 ~ 06-30
	assert.Equal(t, int64(1743436800000), timeRange1[1].Start.UnixMilli())
	assert.Equal(t, int64(1751299199999), timeRange1[1].End.UnixMilli())

	// 3季度 07-01 ~ 09-30
	assert.Equal(t, int64(1751299200000), timeRange1[2].Start.UnixMilli())
	assert.Equal(t, int64(1759247999999), timeRange1[2].End.UnixMilli())

	// 4季度 10-01 ~ 10-16
	assert.Equal(t, int64(1759248000000), timeRange1[3].Start.UnixMilli())
	assert.Equal(t, int64(1760630399999), timeRange1[3].End.UnixMilli())
	//assert.Equal(t, int64(1767196799999), timeRange1[3].End.UnixMilli())

}
