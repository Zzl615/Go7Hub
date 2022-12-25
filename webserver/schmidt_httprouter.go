/**
 * @Author: Noaghzil
 * @Date:   2022-12-25 22:43:37
 * @Last Modified by:   Noaghzil
 * @Last Modified time: 2022-12-26 07:45:05
 */
package webserver

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func UserInfo(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}

func ToolIndex(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "This is tool, %s!\n", ps.ByName("name"))
}

//一个新类型，用于存储域名对应的路由
type HostSwitch map[string]http.Handler

//实现http.Handler接口，进行不同域名的路由分发
func (hs HostSwitch) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	//根据域名获取对应的Handler路由，然后调用处理（分发机制）
	if handler := hs[r.Host]; handler != nil {
		handler.ServeHTTP(w, r)
	} else {
		http.Error(w, "Forbidden", 403)
	}
}

func Httprouter() {
	//声明两个路由
	playRouter := httprouter.New()
	playRouter.GET("/user/:name", UserInfo)

	toolRouter := httprouter.New()
	toolRouter.GET("/user/:name", ToolIndex)

	//分别用于处理不同的二级域名
	hostSwitch := make(HostSwitch)
	hostSwitch["play.flysnow.org:12345"] = playRouter
	hostSwitch["tool.flysnow.org:12345"] = toolRouter
	//根据域名获取对应的Handler路由，然后调用处理（分发机制）
	log.Fatal(http.ListenAndServe(":8080", hostSwitch))
}
