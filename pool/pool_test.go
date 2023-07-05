package pool

import (
	"testing"
)

func BenchmarkPool(b *testing.B) {
	pool := NewPool(100, 100, 10)
	for i := 0; i < b.N; i++ {
		pool.Schedule(func() {
			// 模拟一个耗时的任务
			for j := 0; j < 1000000; j++ {
				_ = j * j
			}
			//runtime.NumCPU()
		})
	}
}
