package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"mccomack/goback/httptest"
	"mccomack/goback/idkwhatthisiscalledmaybemodule"
)

func main() {
	fmt.Println("Hello, World!")

	var a string
	_, err := fmt.Scanln(&a)
	if err != nil {
		log.Panicln(err)
	}

	fmt.Println(a)

	idkwhatthisiscalledmaybemodule.Goback()

	srv := httptest.New()

	go func() {
		if err := srv.Run(); err != nil {
			log.Println("server stopped:", err)
		}
	}()

	// Ctrl-C나 기타 terminate 신호 오면 os 기본동작 대신 여기서 처리
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM)
	<-sc                                          // 대기
	signal.Reset(syscall.SIGINT, syscall.SIGTERM) // 두번 누르면 os 기본동작

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Stop(ctx); err != nil {
		log.Println("shutdown error:", err)
	}
}
