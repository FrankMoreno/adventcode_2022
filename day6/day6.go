package day6

func Part1(input string) int {
	left, right := 0, 0
	found := make(map[byte]int)

	for ; right < len(input); right++ {
		foundIndex, ok := found[input[right]]
		if !ok {
			if right-left == 13 {
				return right + 1
			}
		} else {
			for i := left; i <= foundIndex; i++ {
				delete(found, input[i])
			}
			left = foundIndex + 1
		}
		found[input[right]] = right
	}

	return -1
}
