package leetcode

import (
	"fmt"
	"sort"
	"strings"
)

func MyTest() {

}

//汉明距离
//异或求1的个数
func hammingDistance(x int, y int) int {
	a := x ^ y
	count := 0
	for i := 0; i < 32; i++ {
		count += (a >> uint(i)) & 0x01
	}
	return count
}

//判断是否为圆
func judgeCircle(moves string) bool {
	x := 0
	y := 0
	for _, m := range moves {
		if m == 'L' {
			x--
		} else if m == 'R' {
			x++
		} else if m == 'U' {
			y--
		} else {
			//D
			y++
		}
	}
	return x == 0 && y == 0
}

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//两树叠加
func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 == nil {
		return t2
	}
	if t2 == nil {
		return t1
	}
	t1.Val = t2.Val + t1.Val
	t1.Left = mergeTrees(t1.Left, t2.Left)
	t1.Right = mergeTrees(t1.Right, t2.Right)
	return t1
}

//2n元素数组，分为n组，每组中较小元素之和尽量最小
func arrayPairSum(nums []int) int {
	sort.Ints(nums)
	sum := 0
	for i := 0; i < len(nums); i += 2 {
		sum += nums[i]
	}
	return sum
}

//求一个数的补码
func findComplement(num int) int {
	bit := 0
	//从左往右查找第一个不为0的数的位置
	for i := 0; i < 32; i++ {
		if ((num >> uint(32-i-1)) & 1) != 0 {
			bit = i
			break
		}
	}

	fmt.Println(bit)

	return num ^ (0xffffffff >> uint(bit))
}

//匹配字符串的所有字符是否在键盘的同一行
func findWords(words []string) []string {
	w1 := "qwertyuiop"
	w2 := "asdfghjkl"
	w3 := "zxcvbnm"
	var r []string

	for _, v := range words {
		if isOnOneLine(w1, v) {
			r = append(r, v)
			continue
		}
		if isOnOneLine(w2, v) {
			r = append(r, v)
			continue
		}
		if isOnOneLine(w3, v) {
			r = append(r, v)
			continue
		}

	}
	return r
}

func isOnOneLine(w, v string) bool {
	v = strings.ToLower(v)
	for _, s := range v {
		if !strings.Contains(w, string(s)) {
			return false
		}

	}
	return true
}

//字符串中所有 单词逆序显示
func reverseWords(s string) string {
	strs := strings.Split(s, " ")
	var reversWord string
	for k, str := range strs {
		var nb []byte
		for i := 0; i < len(str); i++ {
			w := []byte(str)[len(str)-1-i]
			nb = append(nb, w)
		}
		if k == len(strs)-1 {
			reversWord += string(nb)
			break
		}
		reversWord += string(nb) + " "
	}
	return reversWord
}

//字符串倒序
func reverseString(s string) string {
	length := len(s)
	b := []byte(s)
	for i := 0; i < length/2; i++ {
		var tmp byte
		tmp = b[i]
		b[i] = b[length-1-i]
		b[length-1-i] = tmp
	}
	return string(b)
}

//
func calPoints(ops []string) int {
	return 0
}
