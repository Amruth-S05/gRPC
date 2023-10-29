package main

import (
  "context"
  "log"
  "time"

  pb "github.com/amruth-s05/grpc/proto"
)

func callSayHelloClientStream(client pb.GreetServiceClient, names *pb.NameList) {
  log.Println("Client streaming started")
  stream, err := client.SayHelloClientStreaming(context.Background())
  if err != nil {
    log.Fatal(err)
  }
  for _, name := range names.Names {
    req := &pb.HelloRequest {
      Name: name,
    }
    if err := stream.Send(req); err != nil {
      log.Fatal(err)
    }
    log.Printf("sent the request with name: %s", name)
    time.Sleep(time.Second * 2)
  }
  res, err := stream.CloseAndRecv()
  log.Printf("client streaming finished\n")
  if err != nil {
    log.Fatal(err)
  }
  log.Printf("%v", res.Messages)
}
