/**
 * @Author: Noaghzil
 * @Date:   2022-12-23 00:20:18
 * @Last Modified by:   Noaghzil
 * @Last Modified time: 2023-01-18 11:23:20
 */
package webserver

import (
	"fmt"
	"net/http"
)

type handle1 struct{}

func (h1 *handle1) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "handle one")
}

type handle2 struct{}

func (h2 *handle2) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "handle two")
}

func NetHttpServer() {
	handle1 := handle1{}
	handle2 := handle2{}
	//Handler:nil表明服务器使用默认的多路复用器DefaultServeMux
	s := &http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: nil,
	}

	//Handle函数调用的是多路复用器DefaultServeMux.Handle方法
	http.Handle("/handle1", &handle1)
	http.Handle("/handle2", &handle2)
	s.ListenAndServe()
}
