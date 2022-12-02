package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/rongpengju/delivery_problem/internal/biz"
	"github.com/rongpengju/delivery_problem/internal/pkg/etc"
)

var (
	// flagconf is the config flag.
	flagconf string
	// weight
	weight float64
	// user id
	userid string
)

func init() {
	flag.StringVar(&flagconf, "conf", "configs/config.yaml", "config path, eg: -conf config.yaml")
	flag.Float64Var(&weight, "weight", 1.0, "please input weight")
	flag.StringVar(&userid, "uer", "", "please input user id")
}

func main() {
	flag.Parse()

	entrypoint, err := initEntrypoint()
	if err != nil {
		log.Panic(err)
	}

	// 快递费计算
	courier := biz.NewCourier(entrypoint.Config)
	fmt.Println("*****************************")
	fmt.Printf("\t合计费用为: %v元\n", courier.Calculate(weight))
	fmt.Println("*****************************")
}

func newConfig() (*etc.Config, error) {
	config, err := etc.LoadFromFile(flagconf)
	if err != nil {
		return nil, err
	}
	return config, err
}
