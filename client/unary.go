package main

import (
  "context"
  "time"
  "log"

  pb "github.com/amruth-s05/grpc/proto"
)

func callSayHello(client pb.GreetServiceClient) {
  ctx, cancel := context.WithTimeout(context.Background(), time.Second)
  defer cancel()
  res, err := client.SayHello(ctx, &pb.NoParam{})
  if err != nil {
    log.Fatal(err)
  }
  log.Printf("%s", res.Message)
}
