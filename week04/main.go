package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	// Week 04: ここに課題のコードを記述してください
	// 詳細な課題内容はLMSで確認してください

	fmt.Println("Week 04 課題")
	http.HandleFunc("/info", infohandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func infohandler(w http.ResponseWriter, r *http.Request) {
	jst, _ := time.LoadLocation("Asia/Tokyo")
	fmt.Fprintln(w, (time.Now().In(jst)).Format("2006年01月02日 15:04:05"))
	now := time.Now().In(jst)

	ua := r.Header.Get("User-Agent")
	fmt.Fprintf(w, "今の時刻は%sで、利用しているブラウザは%sですね。", now.Format("2006年01月02日 15:04:05"), ua)
}
