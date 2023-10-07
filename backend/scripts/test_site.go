package scripts

import (
	"fmt"
	"net"
	"sync"
)

func TestOpenPorts(domain string) {

	wg := sync.WaitGroup{}

	for i := 0; i < 65535; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", domain, i))
			if err != nil {
				return
			}
			conn.Close()
			fmt.Printf("PORT %d OPEN\n", i)
		}(i)

	}
	wg.Wait()

}
