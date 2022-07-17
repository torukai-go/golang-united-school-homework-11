package batch

import (
	"sync"
	"time"

	"golang.org/x/sync/errgroup"
)

type user struct {
	ID int64
}

func getOne(id int64) user {
	time.Sleep(time.Millisecond * 100)
	return user{ID: id}
}

func getBatch(n int64, pool int64) (res []user) {

	errg := new(errgroup.Group)
	errg.SetLimit(int(pool))
	var mutex sync.Mutex
	var i int64

	for i = 0; i < n; i++ {
		id := i
		errg.Go(func() error {
			tempUser := getOne(id)
			mutex.Lock()
			defer mutex.Unlock()
			res = append(res, tempUser)

			return nil
		})
	}

	errg.Wait()
	return res
}
