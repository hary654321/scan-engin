package utils

import (
	"math/rand"
	"time"
)

func Max(nums ...int) int {
	var maxNum = 0
	for _, num := range nums {
		if num > maxNum {
			maxNum = num
		}
	}
	return maxNum
}

func RanNum(max int) int {

	//slog.Println(slog.INFO, max)
	rand.Seed(time.Now().UnixNano())

	// 表示生成 [0,50)之间的随机数
	res := rand.Intn(max)

	//slog.Println(slog.INFO, res)
	return res
}
