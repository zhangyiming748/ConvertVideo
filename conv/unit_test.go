package conv

import (
	"sort"
	"testing"
)

/*
一亿数据全部为整型数字 1g内存 如何挑k个最大值 golang实现
我的思路是拆分成每个小文件
找到小文件中最大的值
在这些小文件中再排序找到最大的k个值
我查了几个AI也都是这样说的
但是明显存在的问题就是
一个小文件里最小的值有可能就是全部数字里第二大的值
比如
拆分之后一共就两个小文件
排序之后内容分别是
[1,2,3]
[5,6,7]
这种情况下如果k=2
得到的结果是[3,7]
而实际上应该是[6,7]
有大佬有思路吗?
*/
func TestSelect(t *testing.T) {
	file1 := []int{1, 2, 3}
	file2 := []int{6, 7, 8}
	file3 := []int{3, 2, 1}
	max1 := getMax(file1)
	max2 := getMax(file2)
	max3 := getMax(file3)
	k := 2
	maxs := []int{max1, max2, max3}
	sort.Slice(maxs, func(i, j int) bool {
		return maxs[i] > maxs[j]
	})
	t.Log(maxs[0:k])
	//函数返回 8,3 实际上期望是 7,8
}
func getMax(i []int) int {
	maxNum := i[0]
	for _, v := range i {
		if v > maxNum {
			maxNum = v
		}
	}
	return maxNum
}
