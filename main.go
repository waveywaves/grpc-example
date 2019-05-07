package main

import (
	"fmt"
	"os"

	server "github.com/waveywaves/grpc-example/server"
)

func main() {
	if err := server.RunServer(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
