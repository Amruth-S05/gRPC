package main

import (
  "net"
  "log"
  "google.golang.org/grpc"
  pb "github.com/amruth-s05/grpc/proto"
)

const (
  PORT = ":8080"
)

type helloServer struct {
  pb.GreetServiceServer
}

func main() {
  ls, err := net.Listen("tcp", PORT)
  if err != nil {
    log.Fatal(err)
  }
  grpcServer := grpc.NewServer()
  pb.RegisterGreetServiceServer(grpcServer, &helloServer{})
  log.Printf("Server started at %v", ls.Addr())
  if err := grpcServer.Serve(ls); err != nil {
    log.Fatal(err)
  }
}
