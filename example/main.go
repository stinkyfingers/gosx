package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"syscall/js"

	"github.com/stinkyfingers/gosx/example/components"
)

//https://github.com/siongui/godom/tree/master/wasm
func main() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM|syscall.SIGQUIT|syscall.SIGINT)

	body := js.Global().Get("document").Get("body")

	ctx, cf := context.WithCancel(context.Background())
	components.YearSelect(body, ctx)
	<-ch
	cf()
}
