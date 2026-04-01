package arrays

func Sum(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

func SumAll(nums ...[]int) (sums []int) {
	for _, num := range nums {
		sums = append(sums, Sum(num))
	}
	return sums
}

func TotalSum(nums ...[]int) (total int) {
	sums := SumAll(nums...)

	for _, sum := range sums {
		total += sum
	}

	return total

}

func SumTail(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	return Sum(nums[1:])
}

func SumAllTails(nums ...[]int) (sums []int) {
	for _, num := range nums {
		if len(num) == 0 {
			sums = append(sums, 0)
		} else {
			sums = append(sums, SumTail(num))
		}
	}
	return sums
}
