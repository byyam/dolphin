package main

import (
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
	router := gin.Default()
	router.POST("/greeting", GreetingGRPC)
	router.Run(common.Conf.Service.ListenAddr) // Do not use the default port: 8080
}

func GreetingGRPC(c *gin.Context) {
	// read body
	body := c.Request.Body
	x, err := ioutil.ReadAll(body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Printf("receive:%s", string(x))

	// json 2 pb
	req := &hello.GreetingRequest{}
	if err := jsonpb.UnmarshalString(string(x), req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Printf("req:%+v", req)

	// grpc connection
	conn, err := grpc.Dial(common.Conf.Service.TCPServerAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	client := hello.NewHelloClient(conn)

	// call rpc
	rsp, err := client.Greeting(context.Background(), req)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	log.Printf("rsp:%v", rsp)

	marshaler := &jsonpb.Marshaler{}
	data, err := marshaler.MarshalToString(rsp)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	log.Printf("data:%s", data)

	c.JSON(http.StatusOK, data)
}
