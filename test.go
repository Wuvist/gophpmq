package main

import (
	"fmt"
	"test/msg"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/spiral/roadrunner"
)

func main() {
	srv := roadrunner.NewServer(
		&roadrunner.ServerConfig{
			Command: "php phpclient/client.php echo pipes",
			Relay:   "pipes",
			Pool: &roadrunner.Config{
				NumWorkers:      1,
				AllocateTimeout: time.Second,
				DestroyTimeout:  time.Second,
			},
		})
	defer srv.Stop()

	srv.Start()

	m := &msg.Msg{}
	m.Msg = "hello"
	m.Version = 2019

	data, err := proto.Marshal(m)
	if err != nil {
		panic(err)
	}

	res, err := srv.Exec(&roadrunner.Payload{Body: data})
	if err != nil {
		panic(err)
	}
	fmt.Println(string(res.Body))
}
