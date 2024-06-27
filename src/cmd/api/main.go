package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/familybook-project/familybook-api-gin/src/internal/app/middlewares"
	"github.com/familybook-project/familybook-api-gin/src/internal/app/users/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(middlewares.RecordUaAndTime)
	middlewares.InitDB()

	v1 := r.Group("/v1")
	{
		// 死活監視用
		v1.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})
		v1.GET("/users", handlers.GetUsersEndpoint)
		v1.GET("/users/:id", handlers.GetUserByIdEndpoint)
	}
	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	// サーバーをゴルーチンで起動し、グレースフルシャットダウンを妨げないようにする
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// シグナルを受け取るためのチャネルを作成
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-quit
	log.Println("Shutting down server...")

	// ヘルスチェックとロードバランサ通知を行う
	notifyLoadBalancer()

	// 現在処理中のリクエストを完了するためにタイムアウトを5秒で設定
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Println("Server forced to shutdown: ", err)
	}

	log.Println("Server exiting")
}

// ヘルスチェックとロードバランサ通知の擬似的な関数
func notifyLoadBalancer() {
	// ここにロードバランサに通知するコードを追加
	// 例: ヘルスチェックエンドポイントにリクエストを送信して、インスタンスを非アクティブにする
	log.Println("Notifying load balancer to stop sending traffic...")
	time.Sleep(2 * time.Second) // 擬似的な遅延
	log.Println("Load balancer notification complete.")
}
