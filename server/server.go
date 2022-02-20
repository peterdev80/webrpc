package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	pb "gprcWeb/calculatorpb"
	"log"
	"net"
	"time"
)

type Server struct {
	pb.CalculatorServer
}

func (s *Server) Add(context context.Context, req *pb.AddRequest) (*pb.AddResponse, error) {
	fmt.Println("Got a new Add request")
	num1 := req.GetNum1()
	num2 := req.GetNum2()
	sum := num1 + num2
	result := &pb.AddResponse{Result: sum}
	return result, nil
}

func (s *Server) Fibonacci(req *pb.FibonacciRequest, stream pb.Calculator_FibonacciServer) error {
	fmt.Println("Got a new Fibonacci request")
	count := req.GetCount()
	c := make(chan int, 3)
	go func(count int, c chan int) {
		if count == 1 {
			fmt.Println("COunt 1")
			c <- 1
		} else if count == 2 {
			fmt.Println("Count 2")
			c <- 1
			c <- 1
		} else {
			num1 := 1
			num2 := 1
			c <- num1
			c <- num2
			for i := 3; i <= count; i++ {
				time.Sleep(time.Millisecond * 300)
				result := num1 + num2
				num1 = num2
				num2 = result
				c <- result
			}
			close(c)
		}
	}(int(count), c)

	for num := range c {
		fmt.Println("Reading from channel", num)
		stream.Send(&pb.FibonacciResponse{Number: int32(num)})
	}

	return nil
}

// Unary returns a server interceptor function to authenticate and authorize unary RPC
func Unary() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		md, ok := metadata.FromIncomingContext(ctx)
		log.Println(md)
		var values []string
		if ok {
			values = md["custom-header-1"]
		}

		log.Println("--> unary interceptor: ", info.FullMethod, "Token:", values)
		return handler(ctx, req)
	}
}
func main() {

	fmt.Println("Starting Calculator server")
	lis, err := net.Listen("tcp", ":8000")

	if err != nil {
		log.Fatalf("Error while listening : %v", err)
	}

	s := grpc.NewServer(grpc.UnaryInterceptor(Unary()))
	cd := &Server{}
	pb.RegisterCalculatorServer(s, cd)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Error while serving : %v", err)
	}

}
