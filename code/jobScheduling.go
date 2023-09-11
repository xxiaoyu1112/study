package main

import (
	"code/code/utils"
	"sort"
)

func jobScheduling(startTime, endTime, profit []int) int {
	n := len(startTime)
	type job struct {
		start, end, profit int
	}
	jobs := make([]job, n)
	for i, start := range startTime {
		jobs[i] = job{start, endTime[i], profit[i]}
	}
	sort.Slice(jobs, func(i, j int) bool { return jobs[i].end < jobs[j].end }) // 按照结束时间排序

	f := make([]int, n+1)
	for i, job := range jobs {
		// j 表示满足结束时间小于等于第 i 份工作开始时间的兼职工作数量
		j := sort.Search(i, func(j int) bool { return jobs[j].end > job.start })
		// 为什么是 j 不是 j+1：上面算的是 > start，-1 后得到 <= start，但由于还要 +1，抵消了
		f[i+1] = utils.Max(f[i], f[j]+job.profit)
	}
	return f[n]
}

func main() {
	startTime, endTime, profit := []int{1, 2, 3, 3}, []int{3, 4, 5, 6}, []int{50, 10, 40, 70}
	jobScheduling(startTime, endTime, profit)
}
