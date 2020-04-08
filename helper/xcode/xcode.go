package xcode

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
)

//生成单号
//06123xxxxx
//sum 最少10位,sum 表示全部单号位数
func MakeYearDaysRand(sum int) string {
	//年
	strs := time.Now().Format("06")
	//一年中的第几天
	days := strconv.Itoa(GetDaysInYearByThisYear())
	count := len(days)
	if count < 3 {
		//重复字符0
		days = strings.Repeat("0", 3-count) + days
	}
	//组合
	strs += days
	//剩余随机数
	sum = sum - 5
	if sum < 1 {
		sum = 5
	}
	//0~9999999的随机数
	//ran := GetRand()
	//pow := math.Pow(10, float64(sum)) - 1
	//fmt.Println("sum=>", sum)
	//fmt.Println("pow=>", pow)
	result := strconv.Itoa(GetRand())
	count = len(result)
	//fmt.Println("result=>", result)
	if count < sum {
		//重复字符0
		result = strings.Repeat("0", sum-count) + result
	}
	//组合
	strs += result
	return strs
}

func GetRand() int {
	rand.Seed(time.Now().UnixNano())
	//rand.Seed(pow)
	num := rand.Intn(30000)
	return num
}

//年中的第几天
func GetDaysInYearByThisYear() int {
	now := time.Now()
	total := 0
	arr := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	y, month, d := now.Date()
	m := int(month)
	for i := 0; i < m-1; i++ {
		total = total + arr[i]
	}
	if (y%400 == 0 || (y%4 == 0 && y%100 != 0)) && m > 2 {
		total = total + d + 1

	} else {
		total = total + d
	}
	return total
}
