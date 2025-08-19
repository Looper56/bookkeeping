package util

import (
	"database/sql/driver"
	"time"
)

const TimeFormat = "2006-01-02 15:04:05"

// LocalTime 时间处理类
type LocalTime time.Time

// UnmarshalJSON 自定义时间 unmarshal
func (t *LocalTime) UnmarshalJSON(data []byte) error {
	// 空值不进行解析
	if len(data) == 2 {
		*t = LocalTime(time.Time{})
		return nil
	}

	// 指定解析的格式
	now, err := time.Parse(`"`+TimeFormat+`"`, string(data))
	if err != nil {
		return err
	}
	*t = LocalTime(now)
	return nil
}

// MarshalJSON 自定义时间 marshal
func (t *LocalTime) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(TimeFormat)+2)
	b = append(b, '"')
	b = time.Time(*t).AppendFormat(b, TimeFormat)
	b = append(b, '"')
	return b, nil
}

// ToTime 转换成 time.Time
func (t *LocalTime) ToTime() time.Time {
	return time.Time(*t)
}

// Value mysql 时间转换
func (t *LocalTime) Value() (driver.Value, error) {
	return time.Time(*t).Format(TimeFormat), nil
}
