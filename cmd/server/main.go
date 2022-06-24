package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/1920853199/passwd/service"
	"github.com/1920853199/passwd/util"
	"github.com/1920853199/passwd/cmd/server/api"
	"github.com/dgrijalva/jwt-go"
)

func init() {

}

func main() {
	var (
		host string
		port int
		path string
	)
	flag.StringVar(&path, "path", util.GetConfigPath(), "Directory storing data.")
	flag.StringVar(&host, "host", "0.0.0.0", "The ip on which to serve")
	flag.IntVar(&port, "port", 22622, "The port on which to serve")
	flag.Parse()

	// 初始化数据库
	db, err := service.NewDB(path)
	if err != nil {
		panic(err)
	}
	service.Init(db)

	// 初始化TOKEN
	token, err := util.Jwt.SetToken(jwt.StandardClaims{
		Id:      util.GetMacAddrs(),
		Subject: "passwd",
	})
	if err != nil {
		panic("Token generation failed.")
	}

	// 启动API服务
	mux := http.NewServeMux()
	mux.HandleFunc("/v1/execute", api.Execute)

	util.Println([]string{fmt.Sprintf("Listening and serving HTTP on %s:%d", host, port), token})

	err = http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), mux)
	if err != nil {
		panic(err)
	}
}
