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
