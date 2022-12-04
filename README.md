# Delivery Problem

```go
// go mod tidy

// go build main.go

// 快递费计算
// main -w 1.8 

// 用户uid查询
// main -u 74634624 
```

```text
/*
生成订单思路逻辑：
	1）题目限制最大重量为 100kg，所以我按照 0-100 划分了十个等级([0-10], [10-20]...[90-100]).
	2）因为要创建 100000 条订单记录，则按照随机生成的分布，对应第一步的百分比数据生成订单.
	3）因为要创建 100 个用户，所以用户也可以对应第一步的百分比来生成各自的订单.
	4）用户随机获取重量区间数据，然后从订单号切片中生成 Order 结构体数据
	5）最终构建 100000条，然后100个协程去写入Sqlite
*/
```