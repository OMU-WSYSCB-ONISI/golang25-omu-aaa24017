package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
)

func main() {
	// Week 03: ここに課題のコードを記述してください
	// 詳細な課題内容はLMSで確認してください
	fmt.Println("Week 03 課題")
	http.HandleFunc("/webfortune", fortunehandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
func fortunehandler(w http.ResponseWriter, _ *http.Request) {
	fortunes := []string{"大吉", "中吉", "吉", "凶"}
	fortune := fortunes[rand.Intn(len(fortunes))]
	fmt.Fprintf(w, "今の運勢は%sです", fortune)
}
