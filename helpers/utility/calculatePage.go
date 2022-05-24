package utility

import (
	"Hybrid-blog/constant"
	"math"
)

func CalculatePage(currentPage int, pageNumber int, countQuery int64, limit int) (int, int, []int, []int, int) {
	totalPage := int(math.Ceil(float64(countQuery) / float64(constant.LimitPages)))
	currentPage = pageNumber
	currentPageMinusI := []int{currentPage - 5, currentPage - 4, currentPage - 3, currentPage - 2, currentPage - 1}
	currentPageAddJ := []int{currentPage + 1, currentPage + 2, currentPage + 3, currentPage + 4, currentPage + 5}
	totalPageMinus5 := totalPage - 5
	return currentPage, totalPage, currentPageMinusI, currentPageAddJ, totalPageMinus5
}
