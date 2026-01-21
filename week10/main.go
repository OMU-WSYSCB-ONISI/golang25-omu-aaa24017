package main

import (
	"fmt"
	"html"
	"net/http"
	"os"
	"runtime"
	"time"
)

const saveFile = "public/memo.txt" // データファイルの保存先

func main() {
	fmt.Printf("Go version: %s\n", runtime.Version())

	http.HandleFunc("/hello", hellohandler)
	http.HandleFunc("/memo", memo)
	http.HandleFunc("/mwrite", mwrite)

	fmt.Println("Launch server: http://localhost:8080/memo")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Failed to launch server: %v", err)
	}
}

func hellohandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "こんにちは from Codespace !")
}

func memo(w http.ResponseWriter, r *http.Request) {
	// データファイルを開く
	text, err := os.ReadFile(saveFile)
	if err != nil {
		text = []byte("まだメモはありません。")
	}

	// HTMLを構築（CSSを追加してデザインを改善）
	s := `<html>
	<head>
		<meta charset="utf-8">
		<title>拡張メモ帳</title>
		<style>
			body { font-family: sans-serif; max-width: 600px; margin: 20px auto; padding: 20px; background: #f4f4f9; }
			textarea { width: 100%; height: 100px; padding: 10px; border-radius: 5px; border: 1px solid #ccc; margin-bottom: 10px; }
			input[type='submit'] { background: #007bff; color: white; border: none; padding: 10px 20px; border-radius: 5px; cursor: pointer; }
			input[type='submit']:hover { background: #0056b3; }
			.history { background: white; padding: 15px; border-radius: 5px; border: 1px solid #ddd; white-space: pre-wrap; margin-top: 20px; }
			h2 { color: #333; font-size: 1.2rem; }
		</style>
	</head>
	<body>
		<h2>📝 新規メモを追記</h2>
		<form method='post' action='/mwrite'>
			<textarea name='text' placeholder='ここにメッセージを入力...'></textarea>
			<input type='submit' value='履歴に保存' />
		</form>
		<h2>📜 履歴</h2>
		<div class='history'>` + html.EscapeString(string(text)) + `</div>
	</body>
	</html>`
	w.Write([]byte(s))
}

func mwrite(w http.ResponseWriter, r *http.Request) {
	// 投稿内容の解析
	r.ParseForm()
	if len(r.Form["text"]) == 0 || r.Form["text"][0] == "" {
		http.Redirect(w, r, "/memo", 302)
		return
	}

	// 【機能追加】現在時刻を取得してフォーマット
	now := time.Now().Format("2006-01-02 15:04:05")
	newEntry := fmt.Sprintf("--- %s ---\n%s\n\n", now, r.Form["text"][0])

	// 【機能追加】既存のファイルに「追記」する (os.O_APPEND)
	f, err := os.OpenFile(saveFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("File open error:", err)
		return
	}
	defer f.Close()

	if _, err := f.WriteString(newEntry); err != nil {
		fmt.Println("Write error:", err)
	}

	fmt.Println("Saved at " + now)

	// メモ画面へリダイレクト
	http.Redirect(w, r, "/memo", 303)
}
