package main

import (
  "log"
  "context"
  "time"
  "io"

  pb "github.com/amruth-s05/grpc/proto"
)

func callBidirectionalStream(client pb.GreetServiceClient, names *pb.NameList) {
  log.Printf("Bidirectional streaming started")
  stream, err := client.SayHelloBidirectionalStreaming(context.Background())
  if err != nil {
    log.Fatal(err)
  }
  waitc := make(chan struct{})

  go func() {
    for {
      message, err := stream.Recv()
      if err == io.EOF { break }
      if err != nil { log.Fatal(err) }
      log.Println(message)
    }
    close(waitc)
  }()

  for _, name := range names.Names {
    req := &pb.HelloRequest {
      Name: name,
    }
    if err := stream.Send(req); err != nil {
      log.Fatal(err)
    }
    time.Sleep(2 * time.Second)
  }
  stream.CloseSend()
  <-waitc
  log.Printf("Bidirectional streaming finished")
}
