package arrays

func DeleteDuplicates(srtArr []int) []int {
	if len(srtArr) == 1 {
		return srtArr
	}
	prev := srtArr[0]
	deduped := []int{prev}
	for i := 1; i < len(srtArr); i++ {
		current := srtArr[i]
		if prev != current {
			deduped = append(deduped, current)
			prev = current
		}
	}
	return deduped
}

func DeleteDuplicatesInPlace(srtArr []int) []int {
	writeIdx := 1
	for i := 1; i < len(srtArr); i++ {
		if srtArr[writeIdx-1] != srtArr[i] {
			srtArr[writeIdx] = srtArr[i]
			writeIdx++
		}
	}
	return srtArr[0:writeIdx]
}
