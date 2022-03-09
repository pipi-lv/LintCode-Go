package main

import "math"

//有1000个桶，有且仅有一个桶里面装了毒药，其他的都装了水。这些桶从外面看上去完全相同。如果一只猪喝了毒药，它将在15分钟内死去。在一个小时内，至少需要多少只猪才能判断出哪一个桶里装的是毒药呢？
//假如一共有 n 个桶，只有一个桶装了毒药。一只猪将在喝完毒药 m 分钟后死去。你需要多少只猪才能在 p 分钟内找出那个装毒药的桶呢？

/**
 * @param buckets: an integer
 * @param minutesToDie: an integer
 * @param minutesToTest: an integer
 * @return: how many pigs you need to figure out the "poison" bucket within p minutes
 */
func poorPigs(buckets int, minutesToDie int, minutesToTest int) int {
	return int(math.Ceil(math.Log(float64(buckets)) / math.Log(float64(minutesToTest/minutesToDie+1))))
}
