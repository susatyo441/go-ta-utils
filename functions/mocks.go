package functions

import "github.com/susatyo441/go-ta-utils/dto"

// REVIEW: Make an independent mocks package?

// Mock paginated aggregation result of T or return empty paginated aggregation.
// Intended for use during Mock.On(...).Run(...) to mock the result of a paginated aggregation
//
// EXAMPLE:
//
//		aggregateResult := MakePaginationResult[mydto.MyDTO]()
//	    // OR
//		aggregateResult := MakePaginationResult([]mydto.MyDTO{{...}})
//		testCases := map[string]struct {
//		    aggregateResult []dto.PaginationResult[mydto.MyDTO]
//		    expectedResult  dto.PaginationResult[mydto.MyDTO]
//		}{
//		    "should success": {
//		        aggregateResult: aggregateResult,
//		        expectedResult:  aggregateResult[0],
//		    },
//		}
func MockPaginationResult[T any](
	data ...[]T,
) []dto.PaginationResult[T] {
	if len(data) == 0 {
		return []dto.PaginationResult[T]{{
			TotalRecords: 0,
			Data:         []T{},
		}}
	} else {
		return []dto.PaginationResult[T]{{
			TotalRecords: len(data[0]),
			Data:         data[0],
		}}
	}
}
