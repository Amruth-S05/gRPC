package main

import (
  "log"
  "context"
  "io"

  pb "github.com/amruth-s05/grpc/proto"
)

func callSayHelloServerStream(client pb.GreetServiceClient, names *pb.NameList) {
  log.Printf("streaming started\n")
  stream, err := client.SayHelloServerStreaming(context.Background(), names)
  if err != nil {
    log.Fatalf("could not send names: %v", err)
  }
  for {
    message, err := stream.Recv()
    if err == io.EOF {
      break
    } else if err != nil {
      log.Fatal(err)
    }
    log.Println(message)
  }
  log.Println("streaming finished")
}
