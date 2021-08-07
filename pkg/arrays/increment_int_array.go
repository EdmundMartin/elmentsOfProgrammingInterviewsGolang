package arrays

/*
Given an array of integers, increment the value by 1
As such an array of {1, 2, 9} would become {1, 3, 0}
*/

func PlusOne(arr []int) []int {
	arr[len(arr)-1]++
	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i] != 10 {
			break
		}
		arr[i] = 0
		arr[i-1] += 1
	}
	if arr[0] == 10 {
		arr[0] = 1
		arr = append(arr, 0)
	}
	return arr
}
