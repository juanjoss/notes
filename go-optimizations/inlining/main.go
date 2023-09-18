package main

func avg(sumFunc func(...float64) float64, nums ...float64) float64 {
	return sumFunc(nums...) / float64(len(nums))
}

func sum(nums ...float64) float64 {
	var total float64

	for i := range nums {
		total += nums[i]
	}

	return total
}

//go:noinline
func sumWithoutInlining(nums ...float64) float64 {
	var total float64

	for i := range nums {
		total += nums[i]
	}

	return total
}

func sumWithHighInliningCost(nums ...float64) float64 {
	var total float64
	numsLen := len(nums) - 1

	total += nums[numsLen]
	total -= nums[numsLen]
	total += nums[numsLen]
	total -= nums[numsLen]
	total += nums[numsLen]
	numsLen--

	total += nums[numsLen]
	total -= nums[numsLen]
	total += nums[numsLen]
	total -= nums[numsLen]
	total += nums[numsLen]
	numsLen--

	total += nums[numsLen]
	total -= nums[numsLen]
	total += nums[numsLen]
	total -= nums[numsLen]
	total += nums[numsLen]
	numsLen--

	total += nums[numsLen]
	total -= nums[numsLen]
	total += nums[numsLen]
	total -= nums[numsLen]
	total += nums[numsLen]
	numsLen--

	total += nums[numsLen]
	total -= nums[numsLen]
	total += nums[numsLen]
	total -= nums[numsLen]
	total += nums[numsLen]
	numsLen--

	return total
}
