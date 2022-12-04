package gen

import (
	"sync"

	"github.com/google/uuid"
)

func GenerateUserId() []uint32 {
	var res []uint32
	var wg sync.WaitGroup
	for i := 0; i < defaultUserNumber; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			userId, err := uuid.NewRandom()
			if err != nil {
				return
			}
			res = append(res, userId.ID())
		}()
		wg.Wait()
	}
	return res
}
