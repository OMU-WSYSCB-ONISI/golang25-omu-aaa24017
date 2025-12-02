package main

import (
	"fmt"
	"net/http"
	"runtime"
	"strconv"
)

func main() {
	// Week 07: ここに課題のコードを記述してください
	// 詳細な課題内容はLMSで確認してください

	fmt.Println("Week 07 課題")

	// 以下に実装してください
	fmt.Printf("Go version: %s\n", runtime.Version())

	http.HandleFunc("/fdump", fdump)
	http.HandleFunc("/cal02", calpmhandler)
	http.Handle("/", http.FileServer(http.Dir("public/")))

	fmt.Println("Launch server...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Failed to launch server: %v", err)
	}
}

func fdump(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Println("errorだよ")
	}
	// フォームはマップとして利用でき以下で内容を確認できる．
	for k, v := range r.Form {
		fmt.Printf("%v : %v\n", k, v)
	}
}

func calpmhandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Println("errorだよ")
	}
	x, _ := strconv.Atoi(r.FormValue("x"))
	y, _ := strconv.Atoi(r.FormValue("y"))

	switch r.FormValue("cal0") {
	case "+":
		fmt.Fprintln(w, x+y)
	case "-":
		fmt.Fprintln(w, x-y)
	case "*":
		fmt.Fprintln(w, x*y)
	case "/":
		fmt.Fprintln(w, x/y)
	}
}
