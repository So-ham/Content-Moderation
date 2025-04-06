package notf

import (
	"log"
	"os"

	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewClient() NotfServiceClient {
	serverAddr := os.Getenv("NOTF_GRPC_ADDR")
	conn, err := grpc.Dial(serverAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		if os.Getenv("ENV") != "LOCAL" {
			log.Fatalf("fail to dial: %v", err)

		}
	}
	client := NewNotfServiceClient(conn)
	return client
}
