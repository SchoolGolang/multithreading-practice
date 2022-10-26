package main

import (
	"context"
	"github.com/SchoolGolang/multithreading-practice/application"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctxRoot := context.Background()
	ctx, cancelFunc := signal.NotifyContext(ctxRoot, syscall.SIGTERM, syscall.SIGINT, os.Interrupt)
	defer cancelFunc()

	application.Run(ctx)

}
