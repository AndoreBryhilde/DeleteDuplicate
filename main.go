package main

import (
    "fmt"
    "math/rand"
    "runtime"
    "sort"
    "time"
)

type TEST1 struct {
    ID         uint
    Test2IDCSV string
}

func generateRandomData(n int) []TEST1 {
    rand.Seed(time.Now().UnixNano())
    data := make([]TEST1, n)
    for i := 0; i < n; i++ {
        id := uint(rand.Intn(n / 2)) // 製造重複
        data[i] = TEST1{
            ID:         id,
            Test2IDCSV: fmt.Sprintf("csv-%d", rand.Intn(10000)),
        }
    }
    return data
}

// 方法 1：切片線性檢查
func dedupSliceLinear(data []TEST1) []TEST1 {
    var result []TEST1
    for _, item := range data {
        found := false
        for _, r := range result {
            if r.ID == item.ID {
                found = true
                break
            }
        }
        if !found {
            result = append(result, item)
        }
    }
    return result
}

// 方法 2：Map 去重
func dedupMap(data []TEST1) []TEST1 {
    m := make(map[uint]TEST1)
    for _, item := range data {
        m[item.ID] = item
    }
    result := make([]TEST1, 0, len(m))
    for _, v := range m {
        result = append(result, v)
    }
    return result
}

// 方法 3：排序 + 線性掃描
func dedupSorted(data []TEST1) []TEST1 {
    sort.Slice(data, func(i, j int) bool {
        return data[i].ID < data[j].ID
    })
    var result []TEST1
    for i, item := range data {
        if i == 0 || item.ID != data[i-1].ID {
            result = append(result, item)
        }
    }
    return result
}

// 方法 4：泛型保序 Map 去重
func dedupGeneric(data []TEST1) []TEST1 {
    seen := make(map[uint]struct{})
    var result []TEST1
    for _, item := range data {
        if _, ok := seen[item.ID]; !ok {
            seen[item.ID] = struct{}{}
            result = append(result, item)
        }
    }
    return result
}

// 通用效能測量
func measure(name string, fn func([]TEST1) []TEST1, input []TEST1) {
    var m1, m2 runtime.MemStats
    runtime.ReadMemStats(&m1)
    start := time.Now()

    result := fn(input)

    elapsed := time.Since(start)
    runtime.ReadMemStats(&m2)

    usedMB := float64(m2.Alloc-m1.Alloc) / 1024.0 / 1024.0
    fmt.Printf("%-20s ➜ Count: %-6d | Time: %-9v | MemUsed: %.2f MB\n", name, len(result), elapsed, usedMB)
}

func main() {
    fmt.Println("Generating 100,000 records...")
    data := generateRandomData(100_000)
    fmt.Printf("Original record count: %d\n\n", len(data))

    fmt.Println("====== Deduplication Benchmark ======")
    measure("Slice Linear", dedupSliceLinear, data)
    measure("Map", dedupMap, data)
    measure("Sorted + Scan", dedupSorted, data)
    measure("Generic Map", dedupGeneric, data)
}
