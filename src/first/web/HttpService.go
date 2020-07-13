package web

import "net/http"

func CusService() {
	//文件服务器将当前目录作为根目录
	http.Handle("/", http.FileServer(http.Dir(".")))
	// 默认的 HTTP 服务侦听在本机 8080 端口
	http.ListenAndServe(":8055", nil)
}
