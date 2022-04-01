package sort

import "fmt"

/**
 * GoLand
 * @author AnnnJ
 * @date 2022/3/31 18:13
 */

// BubbleSort1 两两比较，但是不是全部进行相邻的比较
func BubbleSort1(data []int) {
	len := len(data)

	var compareCount = 0 //比较元素大小的总次数
	var moveCount = 0    //移动元素的总次数

	for i := 0; i < len-1; i++ {
		for j := i + 1; j <= len-1; j++ {
			compareCount++
			if data[i] > data[j] {
				data[j-1], data[j] = data[j], data[j-1]
				moveCount++
			}
		}
	}

	fmt.Printf("BubbleSort1 总共比较了 %d 次，移动了 %d 次\n", compareCount, moveCount)
}

// BubbleSort2 真正的两两相邻的比较
func BubbleSort2(data []int) {
	len := len(data)

	var compareCount = 0 //比较元素大小的总次数
	var moveCount = 0    //移动元素的总次数

	for i := 0; i < len-1; i++ {
		for j := len - 1; j > i; j-- {
			compareCount++
			if data[j-1] > data[j] {
				data[j-1], data[j] = data[j], data[j-1]
				moveCount++
			}
		}
	}

	fmt.Printf("BubbleSort2 总共比较了 %d 次，移动了 %d 次\n", compareCount, moveCount)
}

// BubbleSort3 添加一个是否进行了比较的标识，去控制极端情况，即数组原本就已经是有序的了，没有比较交换位置，所以当移动标识为假时，说明数组有序了
func BubbleSort3(data []int) {
	len := len(data)

	var compareCount = 0 //比较元素大小的总次数
	var moveCount = 0    //移动元素的总次数
	var moveFlag = true  //默认进行了比较操作

	for i := 0; i < len-1 && moveFlag; i++ { //而且要满足移动标识为true，才说明上一轮仍然不是完全顺序的了，所以需要移动，moveFlag需要为true
		for j := len - 1; j > i; j-- {
			compareCount++
			moveFlag = false //假设数据时有序的，就没有必要进行比较了，所以先假设移动标志moveFlag为false
			if data[j-1] > data[j] {
				moveCount++
				data[j-1], data[j] = data[j], data[j-1]
				moveFlag = true
			}
		}
	}

	fmt.Printf("BubbleSort3 总共比较了 %d 次，移动了 %d 次\n", compareCount, moveCount)
	fmt.Printf("移动标志：%v\n", moveFlag)
}
