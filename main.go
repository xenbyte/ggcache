package main

import (
	"context"
	"flag"
	"fmt"
	"ggcache/cache"
	"ggcache/client"
	"log"
)

func main() {

	var (
		listenAddr = flag.String("listenaddr", ":3000", "listen address of the server")
		leaderAddr = flag.String("leaderaddr", "", "listen address of the leader")
	)

	flag.Parse()

	opts := ServerOpts{
		ListenAddr: *listenAddr,
		IsLeader:   len(*leaderAddr) == 0,
		LeaderAddr: *leaderAddr,
	}
	// go func() {
	// 	client, err := client.New(":3000", client.Options{})
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}

	// 	for i := 0; i < 10; i++ {
	// 		SendCommand(client)
	// 	}
	// 	client.Close()
	// 	time.Sleep(time.Second)
	// }()

	server := NewServer(opts, cache.New())
	if err := server.Start(); err != nil {
		fmt.Printf(err.Error())
	}
}

func SendCommand(c *client.Client) {
	_, err := c.Set(context.Background(), []byte("GG"), []byte("FWTF"), 0)
	if err != nil {
		log.Fatal(err)
	}
}
