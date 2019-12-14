package main

import (
	"bytes"
	"context"
	"flag"
	"github.com/byyam/dolphin/common"
	hello "github.com/byyam/dolphin/proto"
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/jsonpb"
	"google.golang.org/grpc"
	"io/ioutil"
	"log"
	"net/http"
	"runtime"
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
	runtime.GOMAXPROCS(runtime.NumCPU())

	router := gin.Default()
	router.POST("/greeting", GreetingGRPC)
	_ = router.Run(common.Conf.Service.ListenAddr) // Do not use the default port: 8080
}

func readRaw(req *http.Request) (body []byte, err error) {
	body, err = ioutil.ReadAll(req.Body)
	if err != nil {
		return
	}
	// https://stackoverflow.com/questions/46948050/how-to-read-request-body-twice-in-golang-middleware?noredirect=1&lq=1
	_ = req.Body.Close()
	req.Body = ioutil.NopCloser(bytes.NewBuffer(body))
	return
}

func GreetingGRPC(c *gin.Context) {
	// read body
	x, err := readRaw(c.Request)
	if err != nil {
		log.Printf("read row body error:%s", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//log.Printf("receive:%s", string(x))

	// json 2 pb
	req := &hello.GreetingRequest{}
	if err := jsonpb.UnmarshalString(string(x), req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//log.Printf("req:%+v", req)

	// gRPC connection
	conn, err := grpc.Dial(common.Conf.Service.TCPServerAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		_ = conn.Close()
	}()
	client := hello.NewHelloClient(conn)

	// call rpc
	rsp, err := client.Greeting(context.Background(), req)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	//log.Printf("rsp:%v", rsp)

	mar := &jsonpb.Marshaler{}
	data, err := mar.MarshalToString(rsp)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	//log.Printf("data:%s", data)

	c.JSON(http.StatusOK, data)
}
