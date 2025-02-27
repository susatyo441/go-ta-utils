package functions

import (
	"github.com/susatyo441/go-ta-utils/dto"
)

// Format aggregation result for a pagination response.
//
// EXAMPLE:
//
//	data := []dto.PaginationResult[mydto.MyDTO]{}
//	aggrErr := uc.MyService.Aggregate(&data, ...)
//	return functions.FormatPaginationResultPtr(data)
func FormatPaginationResult[T any](
	data []dto.PaginationResult[T],
) dto.PaginationResult[T] {
	if len(data) == 0 {
		return dto.PaginationResult[T]{
			TotalRecords: 0,
			Data:         []T{},
		}
	} else {
		return data[0]
	}
}

// Format aggregation result for a pagination response.
// Pointer version of FormatPaginationResult.
//
// EXAMPLE:
//
//	data := []dto.PaginationResult[mydto.MyDTO]{}
//	aggrErr := uc.MyService.Aggregate(&data, ...)
//	return functions.FormatPaginationResultPtr(data)
func FormatPaginationResultPtr[T any](
	data []dto.PaginationResult[T],
) *dto.PaginationResult[T] {
	if len(data) == 0 {
		return &dto.PaginationResult[T]{
			TotalRecords: 0,
			Data:         []T{},
		}
	} else {
		return &data[0]
	}
}

// Format aggregation result for a list response
func FormatListResult[T any](
	data []T,
) []T {
	if len(data) == 0 {
		return []T{}
	} else {
		return data
	}
}
