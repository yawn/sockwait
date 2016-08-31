package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"sync"
	"syscall"
	"time"

	dd "github.com/yawn/doubledash"
)

const app = "sw"

var (
	build   = "undefined"
	version = "unreleased"
)

var (
	sleep   time.Duration
	timeout time.Duration
)

func init() {

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s %s (%s):\n", app, version, build)
		flag.PrintDefaults()
	}

	flag.DurationVar(&timeout, "timeout", 5000*time.Millisecond, "maximum waiting time")
	flag.DurationVar(&sleep, "sleep", 100*time.Millisecond, "time to sleep in between attempts")

	flag.Parse()

}

func main() {

	var (
		args, extra = dd.Split(flag.Args())
		urls        []string
		wait        sync.WaitGroup
	)

	for _, e := range args[0:] {
		urls = append(urls, e)
		wait.Add(1)
	}

	if len(urls) < 1 {
		log.Fatalf("no hosts passed to wait for")
	}

	for _, e := range urls {

		go func(host string) {

			var started = time.Now()

			defer wait.Done()

			for {

				conn, err := net.Dial("tcp", host)

				if err != nil {

					time.Sleep(sleep)

					if time.Now().Sub(started) > timeout {
						log.Fatalf("timeout (%d ms) reached waiting for (at least) %q, giving up",
							timeout/time.Millisecond,
							host)
					}

				} else {
					conn.Close()
					break
				}

			}

		}(e)

	}

	wait.Wait()

	if len(extra) > 0 {

		if err := syscall.Exec(extra[0], extra[0:], os.Environ()); err != nil {
			log.Fatalf("failed to execute %q: %s", extra, err)
		}

	}

}
