package main

import (
	"fmt"
	"net/http"
	"runtime"
	"strconv"
	"strings"
)

func main() {
	// Week 08: ここに課題のコードを記述してください
	// 詳細な課題内容はLMSで確認してください

	fmt.Println("Week 08 課題")

	// 以下に実装してくださ
	fmt.Printf("Go version: %s\n", runtime.Version())
	http.Handle("/", http.FileServer(http.Dir("public/")))

	http.HandleFunc("/avg", avghandler)

	fmt.Println("Launch server...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Failed to launch server: %v", err)
	}

}
func avghandler(w http.ResponseWriter, r *http.Request) {
	var sum, tt int
	var avg float64
	var dist [11]int

	if err := r.ParseForm(); err != nil {
		fmt.Println("errorだよ")
	}
	tokuten := strings.Split(r.FormValue("dd"), ",")
	fmt.Println(tokuten)
	for i := range tokuten {
		tt, _ = strconv.Atoi(tokuten[i])
		sum += tt

		idx := tt / 10
		if tt == 100 {
			idx = 10
		}
		dist[idx]++
	}
	avg = float64(sum / len(tokuten))
	fmt.Fprintln(w, avg)
	fmt.Println(avg)

	for i := 0; i < 10; i++ {
		fmt.Fprintf(w, "%2d~%2d点: %2d人\n", i*10, i*10+9, dist[i])
	}
	fmt.Fprintf(w, "100点：%d人\n", dist[10])
}
