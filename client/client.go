// Simple grpc client.
// Client makes grpc get, insert, get, delete requests and prints received data from service
package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"os/exec"
	"strconv"
	"sync"
	"time"
	"url_shortener/internal/pkg/shortener"
)

const RequestTimeout = 20

// runs grpc client
func main() {
	//cfg := config.NewConfig()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*RequestTimeout)
	defer cancel()

	conn, err := grpc.DialContext(ctx, "localhost"+":"+"50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("client start error: %s", err)
	}

	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Fatalf("client close error: %s", err)
		}
	}(conn)

	client := shortener.NewUrlShortenerServiceClient(conn)
	// run concurrent create requests
	var wg sync.WaitGroup
	for i := 0; i < 200; i++ {
		wg.Add(1)
		go func(i int) {
			created, err := client.Create(ctx, &shortener.CreateUrl{Url: "https://example.com/payload" + strconv.Itoa(i)})
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(created)
			wg.Done()
		}(i)
	}
	wg.Wait()

	//6LAze
	//6LAzh
	//6LAzm
	//get request

	urlResponse, err := client.Get(ctx, &shortener.GetUrl{Url: "6LAze"})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%v\n", urlResponse)

	deleted, err := client.Delete(ctx, &shortener.DeleteUrl{Url: "6LAzh"})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("removed status is " + deleted.GetStatus())

	//open
	redirect, err := client.Redirect(ctx, &shortener.RedirectUrl{Url: "6LAze"})
	if err != nil {
		fmt.Println(err)
	}
	err = exec.Command("open", redirect.GetUrl()).Start()
	if err != nil {
		fmt.Println(err)
	}
}
