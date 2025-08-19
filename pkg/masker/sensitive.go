package masker

import (
	"bookkeeping/pkg/pagination"
	"encoding/json"

	"github.com/pkg/errors"
)

// SensitiveData 敏感数据代理结构体
type SensitiveData[T any] struct {
	Data          T    // 指向原始数据
	MaskSensitive bool // 是否应用敏感数据遮盖
}

// DataSanitizer ...
type DataSanitizer interface {
	PrepareForSerialization()
}

// MarshalJSON ...
func (sd SensitiveData[T]) MarshalJSON() ([]byte, error) {
	if sd.MaskSensitive {
		// 检查 *T 是否实现了 DataSanitizer 接口
		if customizer, ok := interface{}(sd.Data).(DataSanitizer); ok {
			customizer.PrepareForSerialization() // 调用 PrepareForSerialization 方法
		}

		sensitiveData, err := Struct(sd.Data)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		return json.Marshal(sensitiveData)
	}

	return json.Marshal(sd.Data)
}

// WrapWithSensitiveData 单个结构体版本，返回代理结构体
func WrapWithSensitiveData[T any](data T, maskSensitive bool) SensitiveData[T] {
	// 直接创建并返回单个 SensitiveData 实例
	return SensitiveData[T]{Data: data, MaskSensitive: maskSensitive}
}

// WrapWithSensitiveSlice 泛型函数，将任何切片转换为其代理结构体的切片
func WrapWithSensitiveSlice[T any](data []T, maskSensitive bool) []SensitiveData[T] {
	result := make([]SensitiveData[T], len(data))
	for i, v := range data {
		result[i] = SensitiveData[T]{Data: v, MaskSensitive: maskSensitive}
	}
	return result
}

// WrapWithSimplePagination 包装简单分页中的数据为SensitiveData[T]类型
func WrapWithSimplePagination[T any](p *pagination.SimplePagination[T], maskSensitive bool) *pagination.SimplePagination[SensitiveData[T]] {
	newPagination := &pagination.SimplePagination[SensitiveData[T]]{
		AbstractPagination: &pagination.AbstractPagination[SensitiveData[T]]{
			Limit: p.Limit,
			Page:  p.Page,
			Sort:  p.Sort,
		},
		HasNextPage: p.HasNextPage,
	}
	newPagination.SetRows(WrapWithSensitiveSlice(p.Rows, maskSensitive))
	return newPagination
}

// WrapWithLengthAwarePaginate 包装分页中的数据为SensitiveData[T]类型
func WrapWithLengthAwarePaginate[T any](p *pagination.LengthAwarePagination[T], maskSensitive bool) *pagination.LengthAwarePagination[SensitiveData[T]] {
	newPagination := &pagination.LengthAwarePagination[SensitiveData[T]]{
		AbstractPagination: &pagination.AbstractPagination[SensitiveData[T]]{
			Limit: p.Limit,
			Page:  p.Page,
			Sort:  p.Sort,
		},
		TotalRows:  p.TotalRows,
		TotalPages: p.TotalPages,
	}
	newPagination.SetRows(WrapWithSensitiveSlice(p.Rows, maskSensitive))
	return newPagination
}
