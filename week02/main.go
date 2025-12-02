package main

import("fmt"
"time"
"net/http"
"math/rand")

func main() {
	fmt.Println("Week 03 課題")
	http.HandleFunc("/hello", hellohandler)
	http.HandleFunc("/now", nowhandler)
	http.HandleFunc("/headers", headers)
	http.HandleFunc("/dice", dicehandler)

	http.ListenAndServe(":8080", nil)
	seed := time.Now().UnixNano()
	r := rand.New(rand.NewSource(seed))

	fmt.Println(r.Int31())    // 1401608483 (int32)
	fmt.Println(r.Uint32())   // 3362137694 (uint32)
	fmt.Println("-------")
	fmt.Println(r.Int63())    // 9200732467715261966 (int64)
	fmt.Println(r.Uint64())   // 17815840155156866386 (uint64)
	fmt.Println("-------")
	fmt.Println(r.Float32())  // 0.34179267 (float32)
	fmt.Println(r.Float64())  // 0.7233553795829966 (float64)
	fmt.Println("-------")
	fmt.Println(r.Int31n(10000))  // 1451 (int32)
	fmt.Println(r.Int63n(10000))  // 7504 (int64)
}

func hellohandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "こんにちは from Cocespace !")
}
func nowhandler(w http.ResponseWriter, r *http.Request) {
	jst, _ := time.LoadLocation("Asia/Tokyo")
	fmt.Fprintln(w, (time.Now().In(jst)).Format("2006年01月02日 15:04:05"))
}
func dicehandler(w http.ResponseWriter, r *http.Request) {
	seed := time.Now().UnixNano()
	d := rand.New(rand.NewSource(seed))
  fmt.Fprintln(w,d.Int31n(5)+1)
}
func headers(w http.ResponseWriter, r *http.Request){
h:= r.Header
fmt.Fprintln(w,h)
}
