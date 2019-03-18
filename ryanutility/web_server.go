package ryanutility

import (
	"fmt"
	"html"
	"io"
	"log"
	"net/http"
)

//WebServerStart  启动一个监听端口
func WebServerStart() {
	fmt.Println("启动一个web服务端口")
	HelloWorld := func(w http.ResponseWriter, r *http.Request) {
		//w.Header().Set("Content-type","text/html")
		useragent := r.UserAgent()
		metho := r.Method
		io.WriteString(w, "hello world \n")
		io.WriteString(w, useragent+" \n")
		io.WriteString(w, metho+" \n")
		io.WriteString(w, r.RequestURI+"\n")
		r.ParseForm()
		for k, v := range r.Form {
			io.WriteString(w, k+"\n")
			io.WriteString(w, v[0]+"\n")
		}
		const body string = `this is a test <a href="http://www.baidu.com">baidu</a> `
		io.WriteString(w, html.EscapeString(body))
	}

	http.HandleFunc("/", HelloWorld)

	log.Fatal(http.ListenAndServe(":8089", nil))
}
