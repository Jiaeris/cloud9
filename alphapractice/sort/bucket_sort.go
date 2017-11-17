package sort

import (
	"fmt"
	"math"
	"sort"
)

//桶装排序
type bucket struct {
	water []int
}

func NewBuket() *bucket {
	var a []int
	return &bucket{water: a}
}

func bucketSort(arr []int) {
	max := math.MinInt64
	min := math.MaxInt64

	for i := 0; i < len(arr); i++ {
		max = int(math.Max(float64(max), float64(arr[i])))
		min = int(math.Min(float64(min), float64(arr[i])))
	}
	fmt.Println(max, min)

	bucketNum := (max-min)/len(arr) + 1
	var buckets []*bucket
	for i := 0; i < bucketNum; i++ {
		buckets = append(buckets, NewBuket())
	}
	fmt.Println("桶个数：", len(buckets), bucketNum)

	//将每个元素放入桶
	for i := 0; i < len(arr); i++ {
		num := (arr[i] - min) / len(arr)
		buckets[num].water = append(buckets[num].water, arr[i])
	}
	//fmt.Println(buckets[0].water)

	//对每个桶进行排序
	for i := 0; i < len(buckets); i++ {
		sort.Ints(buckets[i].water)
	}

	count := 0
	for i := 0; i < len(buckets); i++ {
		//if !(len(bukets[i].water) > 0) {
		//	continue
		//}
		for j := 0; j < len(buckets[i].water); j++ {
			fmt.Println(buckets[i].water[len(buckets[i].water)-1-j])
			count++
			if count == 100 {
				break
			}
		}
	}
}
