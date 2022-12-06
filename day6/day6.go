package day6

const LENGTH = 14

func Part1(input string) int {
	left, right := 0, 0
	found := make(map[byte]int)

	for ; right < len(input); right++ {
		foundIndex, ok := found[input[right]]

		if ok {
			for i := left; i <= foundIndex; i++ {
				delete(found, input[i])
			}
			left = foundIndex + 1
		} else if right-left+1 == LENGTH {
			return right + 1
		}

		found[input[right]] = right
	}

	return -1
}
