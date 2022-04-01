package sort

/**
 * GoLand
 * @author AnnnJ
 * @date 2022/3/31 18:57
 */

func InsertSort(data []int) {
	if data == nil {
		return
	}

	for i := 1; i < len(data); i++ {
		tmp, j := data[i], i
		if data[j-1] > tmp {
			for j >= 1 && data[j-1] > tmp {
				data[j] = data[j-1]
				j--
			}
		}
	}
}
