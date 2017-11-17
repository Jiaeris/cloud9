package sort

import (
	"log"
)

//默认从小到大排序
//http://blog.csdn.net/hguisu/article/details/7776068/

//插入排序
//将数组视为两段，前段部分为有序数组，后半部分为无序数组。
//开始排序时，前部分只有一个元素
//依次将无序部分元素插入到合适位置。
func insertSort(a []int) []int {
	n := len(a)
	for i := 1; i < n; i++ {
		if a[i] < a[i-1] {
			j := i - 1
			tmp := a[i]
			a[i] = a[i-1]
			for !(j < 0) && tmp < a[j] {
				a[j+1] = a[j]
				j--
			}
			a[j+1] = tmp
		}
		log.Printf("%d", a)
	}
	return a
}

//冒泡排序
//每次冒泡将最大值放入最后,只需冒泡n-1次
//每次参与冒泡个数将少一个n-1-i（减少个数）
func bubbleSort(a []int) []int {
	n := len(a)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1-i; j++ {
			if a[j] > a[j+1] {
				var tmp int
				tmp = a[j+1]
				a[j+1] = a[j]
				a[j] = tmp
			}
		}
		log.Printf("%d", a)
	}
	return a
}

//选择排序
//选择最小（最大）的数放在最前面
//第一次选择放第一个，第二次选择放在第二个
func selectSort(a []int) []int {
	for i := 0; i < len(a); i++ {
		minKey := selectMinKey(a, i)
		if minKey != i {
			var tmp int
			tmp = a[i]
			a[i] = a[minKey]
			a[minKey] = tmp
		}
	}
	log.Printf("%d", a)
	return a
}

func selectMinKey(a []int, j int) (key int) {
	key = j
	for i := j + 1; i < len(a); i++ {
		if a[i] < a[key] {
			key = i
		}
	}
	return
}
