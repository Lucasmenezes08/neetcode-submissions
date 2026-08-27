func hasDuplicate(nums []int) bool {
	numbers := make(map[int]bool)

	for _ , num := range(nums){
		if numbers[num]{
			return true
		}
		numbers[num] = true
	}
	return false
}
