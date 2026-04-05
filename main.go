package main

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os/exec"
	"time"

	"lion-feeder/internal/api"
	"lion-feeder/internal/db"

	"github.com/go-chi/chi/v5"
)

//go:embed static
var staticFiles embed.FS

func main() {
	// 1. DB 초기화
	database, err := db.Init()
	if err != nil {
		log.Fatal("DB 초기화 실패:", err)
	}
	defer database.Close()

	// 2. 라우터 설정
	r := chi.NewRouter()

	// API 라우트
	api.RegisterRoutes(r, database)

	// 정적 파일 서빙 (static/ → /)
	subFS, _ := fs.Sub(staticFiles, "static")
	r.Handle("/*", http.FileServer(http.FS(subFS)))

	// 3. 브라우저 자동 오픈 (1초 딜레이로 서버 먼저 시작)
	go func() {
		time.Sleep(1 * time.Second)
		exec.Command("cmd", "/c", "start", "http://127.0.0.1:8080").Start()
	}()

	// 4. 서버 시작
	fmt.Println("서버 시작: http://127.0.0.1:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}