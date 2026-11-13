package intermediate

import (
	"fmt"
	"time"
)

func TimerFUn() {
	timer1 := time.NewTimer(2 * time.Second)

	<-timer1.C
	fmt.Println("Timer1 fired")

	timer2 := time.NewTimer(time.Second)

	go func() {
		<-timer2.C
		fmt.Println("timer2 is fired")
	}()
	stop := timer2.Stop()
	if stop {
		fmt.Println("timer2 stopped")
	}

}
