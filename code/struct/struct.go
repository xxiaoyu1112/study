package main

func main() {
	nums1 := []int{1, 2, 2}
	//nums2 := []int{2, 4, 5}
	m := make(map[int]struct{})
	for _, v := range nums1 {
		m[v] = struct{}{}
	}
}
