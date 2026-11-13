package intermediate

import (
	"fmt"
	"golang.org/x/time/rate"
	"time"
)

func RateLimiter() {

	limiter := rate.NewLimiter(2, 4)

	for i := 0; i < 10; i++ {
		if limiter.Allow() {
			fmt.Println("Request allowed", i)
		} else {
			fmt.Println("Rate limited", i)
		}
		time.Sleep(200 * time.Millisecond)
	}

}
