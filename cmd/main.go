package main

import (
	"flag"
	"fmt"
	"log"
	"sync"

	"github.com/rongpengju/delivery_problem/internal/biz"
	"github.com/rongpengju/delivery_problem/internal/data"
	"github.com/rongpengju/delivery_problem/internal/pkg/etc"
	"github.com/rongpengju/delivery_problem/internal/pkg/gen"
)

var (
	conf   string  // conf is the config flag.
	weight float64 // weight
	userid int     // user id
)

func init() {
	flag.StringVar(&conf, "conf", "configs/config.yaml", "config path, eg: -conf config.yaml")
	flag.Float64Var(&weight, "w", 0, "please input weight")
	flag.IntVar(&userid, "u", 0, "please input user id")
}

func main() {
	flag.Parse()
	entrypoint, err := initEntrypoint()
	if err != nil {
		log.Panic(err)
	}

	// 快递费计算功能
	if weight > 0 {
		courier := biz.NewCourier(entrypoint.Config)
		fmt.Println("*****************************")
		fmt.Printf("\t合计费用为: %v元\n", courier.Calculate(weight))
		fmt.Println("*****************************")
	} else {
		fmt.Println("请您构建程序并按照README文档中指令使用快递计费功能")
		fmt.Println("*****************************")
	}

	// 生成订单信息
	fmt.Println("数据正在创建，请您等待...")
	orders := gen.GenerateAllOrder()
	db, _ := data.NewDatabase()
	var wg sync.WaitGroup
	for i, _ := range orders {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if err := db.Insert(orders[i]); err != nil {
				log.Println(err)
				return
			}
		}()
		wg.Wait()
	}
	fmt.Println("数据创建完成!")
	fmt.Println("*****************************")

	// 查询用户订单信息
	if userid > 0 {
		orderList, err := db.Get(userid)
		if err != nil {
			log.Println("query order failed: ", err)
			return
		}
		if len(orderList) == 0 {
			fmt.Printf("未找到用户%d的订单信息\n", userid)
			return
		}
		fmt.Printf("用户 %v 的订单信息如下: \n", userid)
		for index, order := range orderList {
			fmt.Printf("%d: \t", index)
			fmt.Printf("\t订单号: %v\t", order.OrderId)
			fmt.Printf("\t重量: %v\t", order.Weight)
			fmt.Printf("\t创建时间: %v\n", order.CreatedAt)
		}
	} else {
		fmt.Println("请您构建程序并按照README文档中指令使用用户订单查询功能")
		fmt.Println("*****************************")
	}

	// 关闭数据库链接
	if err := db.Close(); err != nil {
		log.Println("Close database error: ", err)
		return
	}
}

func newConfig() (*etc.Config, error) {
	config, err := etc.LoadFromFile(conf)
	if err != nil {
		return nil, err
	}
	return config, err
}
