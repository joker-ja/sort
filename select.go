package sort

import "fmt"

/**
 * GoLand
 * @author AnnnJ
 * @date 2022/3/31 18:19
 */

/*
SelectSort
简单选择排序
最好时间：O(n^2)
平均时间：O(n^2)
最坏时间：O(n^2)
辅助存储：O(1)
稳定性：不稳定
备注：n小时较好

原理：
对给定的一组记录，
1.经过第一轮比较后得到最小记录，然后将该记录与第一个记录的位置进行交换；
2.接着对不包括第一个记录以外的其他记录进行第二轮比较，得到最小记录并与第二个记录进行位置交换；
3.重复该过程，直到进行比较的记录只有一个时为止
*/

func SelectSort(data []int) {
	len := len(data)

	var compareCount = 0 //比较元素大小的总次数
	var moveCount = 0    //移动元素的总次数

	for i := 0; i < len; i++ {
		min := i
		for j := i + 1; j < len; j++ {
			compareCount++
			if data[j] < data[min] {
				min = j
			}
		}
		if min != i {
			data[i], data[min] = data[min], data[i]
			moveCount++
		}
	}

	fmt.Printf("SelectSort 总共比较了 %d 次，移动了 %d 次\n", compareCount, moveCount)
}
