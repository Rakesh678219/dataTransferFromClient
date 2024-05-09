package main

import (
	"io"
	"log"
	"net"
	"os"
	"time"

	pb "github.com/Rakesh678219/dataTransferFromClient/protos/chunker" // Import the generated protobuf code

	"google.golang.org/grpc"
)

const (
    port = ":50051"
)

type server struct{
    pb.UnimplementedFileServiceServer
}

func (s *server) UploadFile(stream pb.FileService_UploadFileServer) error {
    // Create a buffer to hold the file data
    var data []byte
    // Record the start time
     startTime := time.Now()
    for {
        chunk, err := stream.Recv()
        if err == io.EOF {
            // File transmission is complete
            log.Println("File transmission complete")
            // Calculate the time taken
            duration := time.Since(startTime)
            log.Printf("Time taken: %s\n", duration)
            file, err := os.Create("/root/output_file.txt")
                if err != nil {
                    log.Println("error in creating file")
                    return nil
                }

                defer file.Close()

                // Write some data to the file
                _, err = file.Write(data)
                if err != nil {
                     log.Println("error in writing file")
                    return nil
                }

            // Process the file data, e.g., save it to disk
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
