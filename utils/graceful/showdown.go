package graceful

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

//当收到一个os.Interrupt 或者 syscall.SIGTERM信号时，停止server
func ShowdownGin(instance *http.Server, timeout time.Duration) {
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("关闭 Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	err := instance.Shutdown(ctx)
	if err != nil {
		log.Fatal("Server 关闭：", err)
	}

	select {
	case <-ctx.Done():
		log.Println("超时5秒.")
	}
	log.Println("Server 退出")
}
