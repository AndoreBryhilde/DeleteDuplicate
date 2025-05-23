# Go 大量資料去重測試專案

本專案使用 Go 搭配 GORM 與 SQLite，產生 100,000 筆模擬資料，實作與比較四種不同的去重邏輯方法，分析其在**效能、記憶體使用與是否保留原順序**上的差異。

---

## 📌 專案目的

實務場景中，因 SQL 使用 `LIKE '%x%'` 查詢，導致資料查詢後有大量 ID 重複，需在程式中進行高效的去重。本專案透過模擬大量資料，觀察不同去重策略的效能與行為。

---

## 🧪 四種去重方法比較

| 方法名稱         | 時間複雜度 | 空間複雜度 | 保留原順序 | 適用情境描述                                   |
|------------------|-------------|-------------|--------------|------------------------------------------------|
| 切片線性檢查      | O(N²)        | O(N)        | ✅           | 資料量小、邏輯簡單、教學與示範用途為主       |
| Map 去重         | O(N)         | O(N)        | ❌           | 追求極致效能但不在乎順序                     |
| 排序 + 掃描去重  | O(N log N)   | O(N)        | ❌ (排序後)  | 順序可被打亂，節省記憶體，不想用 Map         |
| 泛型保序去重     | O(N)         | O(N)        | ✅           | 資料量大但需保順序，最推薦做法                |

---

## ⚙️ 執行方式

go test -bench=. -benchmem

### 1️⃣ 安裝依賴

```bash
go mod tidy
