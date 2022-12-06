package day6

const LENGTH = 13

func Part1(input string) int {
	left, right := 0, 0
	found := make(map[byte]int)

	for ; right < len(input); right++ {
		foundIndex, ok := found[input[right]]
		if !ok {
			if right-left == LENGTH {
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
