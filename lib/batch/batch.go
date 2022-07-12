package batch

import (
	"sync"
	"time"
)

type user struct {
	ID int64
}

func getOne(id int64) user {
	time.Sleep(time.Millisecond * 100)
	return user{ID: id}
}

func getBatch(n int64, pool int64) (res []user) {

	res = make([]user, 0, n)
	ch := make(chan user, n)

	var wg sync.WaitGroup
	sem := make(chan struct{}, pool)

	for i := 0; i < int(n); i++ {
		wg.Add(1)
		sem <- struct{}{}

		go func(j int) {
			defer wg.Done()

			user := getOne(int64(j))
			ch <- user
			<- sem
		}(i)
	}

	wg.Wait()
	close(ch)

	for u := range ch {
		res = append(res, u)
	}

	return res
}
