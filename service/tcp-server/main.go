package main

import (
	"flag"
	"github.com/byyam/dolphin/common"
	hello "github.com/byyam/dolphin/proto"
	"github.com/byyam/dolphin/service/tcp-server/processor"
	"golang.org/x/net/trace"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
)

const (
	defaultConfig = "conf/conf.toml"
)

func main() {
	var recipe string
	flag.StringVar(&recipe, "recipe", defaultConfig, "The configuration file")
	flag.Parse()
	log.Println(recipe)

	if err := common.InitConfig(recipe); err != nil {
		log.Fatalf("Initial configure file failed: %v", err)
	}

	runService()
}

func runService() {
	// listen port
	ln, err := net.Listen("tcp", common.Conf.Service.TCPServerAddr)
	if err != nil {
		log.Fatal(err)
	}

	// grpc service start
	gs := grpc.NewServer()

	// register service
	service := processor.NewHelloProcessor()
	hello.RegisterHelloServer(gs, service)

	// start service
	go gs.Serve(ln)

	trace.AuthRequest = func(req *http.Request) (any, sensitive bool) { return true, true }
	log.Println("Service started successfully.")
	log.Fatal(http.ListenAndServe(common.Conf.Service.TCPServerAddrDebug, nil))
}
