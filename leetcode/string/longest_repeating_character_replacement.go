package string

func characterReplacement(s string, k int) int {
	var (
		numBySMap = make(map[string]int)
		result    int
		start     int
		max       int
	)

	for i := range s {
		num, ok := numBySMap[string(s[i])]
		if ok {
			num += 1
			numBySMap[string(s[i])] = num
		} else {
			num = 1
			numBySMap[string(s[i])] = 1
		}

		if num > max {
			max = num
		}

		if (i + 1 - start - max) > k {
			num2 := numBySMap[string(s[start])]
			num2 -= 1
			numBySMap[string(s[start])] = num2
			start++
		}

		if i+1-start > result {
			result = i + 1 - start
		}
	}

	return result
}
