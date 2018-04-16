package main

import (
	"fmt"
)

// 线性查找
func liner(datalist []int, key int) bool {
		for _, data := range datalist{
			if data == key{
				return true
			}
		}
	return false
}

// 二进制查找：数组有序
func binary(needle int, datalist []int) bool {
	low := 0
	high := len(datalist) - 1

	for low <= high {
		medin := (low + high) / 2

		if datalist[medin] < needle {
			low = medin + 1
		}else {
			high = medin - 1
		}
	}

	if low == len(datalist) || datalist[low] != needle {
		return false
	}

	return true
}

// 插值查找
func interpolation(key int, datalist []int) int {
	min, max := datalist[0], datalist[len(datalist) - 1] //最小值，最大值
	low, high := 0, len(datalist) - 1 // 数组下标

	for {
		if key > max {
			return high + 1
		}// 大于最大值

		if key < min {
			return low
		}// 小于最小值

		var guess int

		if low == high {
			guess = high
		} else {
			size := high - low
			offset := int(float64(size - 1) * (float64(key - min) / float64(max - min)))//todo:公式意义？
			guess = offset + low
		}

		if datalist[guess] == key {
			for guess > 0 && datalist[guess - 1] == key {
				guess = guess - 1
			}
			return guess
		}//往回查询是否有重复

		if datalist[guess] > key {
			high = guess - 1
			max = datalist[high]
		}else {
			low = low + 1
			min = datalist[low]
		}//二进制查找
	}
}


func main(){
	item := []int{123, 123, 1234, 43, 53}
	fmt.Println(liner(item, 43))
	fmt.Println(liner(item, 49))

	item2 := []int{43, 53, 123, 1234}
	fmt.Println(binary(43, item2))
	fmt.Println(binary(49, item2))

	fmt.Println(interpolation(43, item2))
}