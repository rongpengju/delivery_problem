package gen

import (
	"log"
	"sync"
	"time"

	"github.com/sony/sonyflake"
)

var (
	sonyFlake     *sonyflake.Sonyflake
	sonyMachineID uint16
)

func GenerateID() []uint64 {
	// 初始化雪花生成算法配置
	if err := initFormat("2022-12-01", 1); err != nil {
		log.Printf("Init failed, err:%v\n", err)
		return nil
	}

	// 并发生成 100000个
	var res []uint64
	var wg sync.WaitGroup
	for i := 0; i < defaultOrderNumber; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			id, err := generateID()
			if err != nil {
				log.Printf("generate id failed, err:%v\n", err)
			}
			res = append(res, id)
		}()
		wg.Wait()
	}

	return res
}

func initFormat(startTime string, machineID uint16) error {
	sonyMachineID = machineID
	var st time.Time
	st, err := time.Parse("2006-01-02", startTime)
	if err != nil {
		return err
	}
	settings := sonyflake.Settings{
		StartTime: st,
		MachineID: getMachineID,
	}
	sonyFlake = sonyflake.NewSonyflake(settings)
	return nil
}

func getMachineID() (uint16, error) {
	return sonyMachineID, nil
}

func generateID() (uint64, error) {
	id, err := sonyFlake.NextID()
	if err != nil {
		log.Printf("sonyflake init failed, err:%v\n", err)
		return 0, nil
	}
	return id, nil
}
