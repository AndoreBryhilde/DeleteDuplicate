package main

import (
    "testing"
)

// 測試資料只產生一次供所有 benchmark 使用，避免重複產生造成干擾
var testData = generateRandomData(100_000)

// Benchmark 方法格式固定：BenchmarkXxx(b *testing.B)
func BenchmarkDedupSliceLinear(b *testing.B) {
    for i := 0; i < b.N; i++ {
        _ = dedupSliceLinear(testData)
    }
}

func BenchmarkDedupMap(b *testing.B) {
    for i := 0; i < b.N; i++ {
        _ = dedupMap(testData)
    }
}

func BenchmarkDedupSorted(b *testing.B) {
    for i := 0; i < b.N; i++ {
        // 因為排序會直接改變輸入資料順序，這邊先做個 copy
        inputCopy := make([]TEST1, len(testData))
        copy(inputCopy, testData)
        _ = dedupSorted(inputCopy)
    }
}

func BenchmarkDedupGeneric(b *testing.B) {
    for i := 0; i < b.N; i++ {
        _ = dedupGeneric(testData)
    }
}
