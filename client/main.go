package main

import (
  "log"
  "google.golang.org/grpc"
  "google.golang.org/grpc/credentials/insecure"
  pb "github.com/amruth-s05/grpc/proto"
)

const (
  PORT = ":8080"
)

func main() {
  conn, err := grpc.Dial("localhost"+PORT, grpc.WithTransportCredentials(
    insecure.NewCredentials() ))
  if err != nil {
    log.Fatal(err)
  }
  defer conn.Close()
  client := pb.NewGreetServiceClient(conn)
}
