package sort

import (
	"fmt"
	"math/rand"
	"time"
)

//堆排序
const (
	thousand    = 10000
	tenThousand = 10 * thousand
)

var compareTimes int

func heapSortTest() {
	arr := generateRandomArr(10000)
	//arr := generateArr()
	fmt.Println("初始数据 : ", len(arr))
	heapSort(arr)
}

//获取一个指定长度每位值为随机数的数组
//返回数组
func generateRandomArr(generateLength int) []int {
	arr := make([]int, generateLength)
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for i := 0; i < generateLength; i++ {
		arr[i] = r.Intn(tenThousand)
	}
	return arr
}

//获取一个指定长度每位值为随机数的数组
//返回数组
func generateArr() []int {
	arr := []int{5, 6, 20, 9, 5, 1, 0, 100}
	return arr
}

//完全二叉树规则：
//假设共有n个数
//最后一个非叶节点位置n/2，注:此处讨论的位置全都从0开始计数
//左叶子位置为2*i+1,右边叶子位置为2*i+2

//第一步:建立堆结构即（创建最大堆顶或最小堆顶的完全二叉树），此处创建大堆顶
//第二步:将最大值移动到数组末尾，前(size-1)形成一个新的数
//第三步:对新长度为除掉最后一位的二叉树，从堆顶开始调整
//第四步:重复2步骤，直到剩下一个数为止，即进行size-1次调整
func heapSort(a []int) {
	size := len(a)
	//第一步
	for i := size / 2; i >= 0; i-- {
		adjustBinaryTress(a, i, size)
	}
	//第二步...
	var temp int
	for i := size - 1; i >= 0; i-- {
		//将最大堆顶移至末尾
		temp = a[i]
		a[i] = a[0]
		a[0] = temp
		//再调整
		adjustBinaryTress(a, 0, i)
		//fmt.Printf("第%v次 : %v\n", size-i-1, a)
	}
	fmt.Println("堆排序 : ", a)
	fmt.Printf("比较次数 ：%v", compareTimes)
}

//参数:数组，起始调整位置，数组长度
func adjustBinaryTress(a []int, i int, length int) {

	child := 0
	lChild := 2*i + 1
	rChild := lChild + 1

	//若有叶子进入遍历
	for lChild < length {
		//判断 左叶子为止是否在数组倒数第二个数(即判断有无两个叶子)
		//判断哪个叶子的数更大
		//谁大就区谁
		child = lChild
		if lChild < length-1 {
			if a[lChild] < a[rChild] {
				child = rChild
			} else {
				child = lChild
			}
			compareTimes++
		}
		//如果叶子值大于等于父节点值，交换叶子与父节点值
		if a[child] >= a[i] {
			var temp int
			temp = a[i]
			a[i] = a[child]
			a[child] = temp
			compareTimes++
		} else {
			break
		}
		//继续调整 以交换叶子为节点的叶子 直到计算到叶子节点位子超过数组长度为止即2*i+1>length-1
		i = child
		lChild = 2*i + 1
		rChild = lChild + 1
	}

}
