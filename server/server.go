// server.go
package main

import (
	"io"
	"log"
	"net"

	pb "path_to_your_protos" // Import the generated protobuf code

	"google.golang.org/grpc"
)

const (
    port = ":50051"
)

type server struct{}

func (s *server) UploadFile(stream pb.FileService_UploadFileServer) error {
    // Create a buffer to hold the file data
    var data []byte

    for {
        chunk, err := stream.Recv()
        if err == io.EOF {
            // File transmission is complete
            log.Println("File transmission complete")
            // Process the file data, e.g., save it to disk
            // Example: ioutil.WriteFile("uploaded_file.txt", data, 0644)
            return stream.SendAndClose(&pb.UploadResponse{Success: true, Message: "File uploaded successfully"})
        }
        if err != nil {
            return err
        }
        // Append the received chunk to the data buffer
        data = append(data, chunk.Chunk...)
    }
}

func main() {
    lis, err := net.Listen("tcp", port)
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    s := grpc.NewServer()
    pb.RegisterFileServiceServer(s, &server{})
    log.Printf("Server listening on port %s", port)
    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}
