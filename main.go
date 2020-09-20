package main

import (
	"flag"
	"fmt"
	"sync"
	"time"

	"github.com/go-ping/ping"

	"goforbroke1006/gfb-mping/internal"
)

var (
	addrRange string
)

func init() {
	flag.StringVar(&addrRange, "range", "192.168.1.1/24", "Specify addresses range")
	flag.Parse()
}

func main() {
	list, err := internal.GetAddressesList(addrRange)
	if err != nil {
		fmt.Printf("err: %v\n", err.Error())
		return
	}

	wg := sync.WaitGroup{}
	wg.Add(len(list))

	limiter := make(chan struct{}, 256)

	for _, addr := range list {
		limiter <- struct{}{}

		go func(addr string) {
			defer func() {
				wg.Done()
				<-limiter
			}()

			pinger, err := ping.NewPinger(addr)
			if err != nil {
				fmt.Printf("err: %v\n", err.Error())
				return
			}

			pinger.SetPrivileged(false)
			pinger.Timeout = 5 * time.Second
			pinger.Count = 3

			pinger.Run() // blocks until finished

			stats := pinger.Statistics()

			if stats.PacketLoss == 100 {
				return
			}

			fmt.Printf("%s sent: %d recv: %d loss: %.2f\n",
				stats.Addr,
				stats.PacketsSent, stats.PacketsRecv, stats.PacketLoss,
			)

		}(addr)
	}

	wg.Wait()
}
