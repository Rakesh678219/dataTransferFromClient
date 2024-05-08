// client.go
package main

import (
	"context"
	"io"
	"log"
	"os"

	pb "path_to_your_protos" // Import the generated protobuf code

	"google.golang.org/grpc"
)

const (
    address = "localhost:50051"
)

func main() {
    // Set up a connection to the server
    conn, err := grpc.Dial(address, grpc.WithInsecure())
    if err != nil {
        log.Fatalf("did not connect: %v", err)
    }
    defer conn.Close()
    c := pb.NewFileServiceClient(conn)

    // Open the file to upload
    file, err := os.Open("file_to_upload.txt")
    if err != nil {
        log.Fatalf("failed to open file: %v", err)
    }
    defer file.Close()

    // Create a stream to send chunks of the file
    stream, err := c.UploadFile(context.Background())
    if err != nil {
        log.Fatalf("failed to upload file: %v", err)
    }

    // Read and send chunks of the file until EOF
    buffer := make([]byte, 1024)
    for {
        bytesRead, err := file.Read(buffer)
        if err == io.EOF {
            break
        }
        if err != nil {
            log.Fatalf("failed to read file: %v", err)
        }
        // Send the chunk to the server
        if err := stream.Send(&pb.FileChunk{Chunk: buffer[:bytesRead]}); err != nil {
            log.Fatalf("failed to send chunk: %v", err)
        }
    }

    // Close the stream and receive the server response
    response, err := stream.CloseAndRecv()
    if err != nil {
        log.Fatalf("failed to receive response: %v", err)
    }
    log.Printf("Upload response: %v", response)
}
