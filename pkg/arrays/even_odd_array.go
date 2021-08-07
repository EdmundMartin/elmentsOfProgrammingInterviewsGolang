package arrays

func EvenOdd(arr []int) {
	nextEven, nextOdd := 0, len(arr)-1
	for nextEven < nextOdd {
		if arr[nextEven]%2 == 0 {
			nextEven++
		} else {
			arr[nextEven], arr[nextOdd] = arr[nextOdd], arr[nextEven]
			nextOdd--
		}
	}
}
