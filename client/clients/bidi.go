package clients

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/aakash19here/grpc_exam/proto/generated/exampb"
)

func Bidi(client exampb.ExamServiceClient) {
	reader := bufio.NewReader(os.Stdin)
	stream, err := client.LiveExamQuery(context.Background())
	done := make(chan struct{})

	if err != nil {
		fmt.Errorf("Failed to create stream")
		return
	}

	// receive chunks
	go func() {
		for {
			res, err := stream.Recv()
			if err != nil {
				if err == io.EOF {
					break
				}
				log.Fatalf("Error receiving response: %v", err)
				break
			}
			fmt.Printf("🎓 %s | %s: %d/%d (%s)\n",
				res.StudentName, res.Subject, res.MarksObtained, res.TotalMarks, res.Grade)
		}
		close(done)
	}()

	fmt.Print("Enter student_id and exam_id (or 'exit'): ")

	// send data
	for {
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)

		if line == "exit" {
			stream.CloseSend()
		}

		parts := strings.Fields(line)
		if len(parts) != 2 {
			fmt.Println("Usage <student_id> <exam_id>")
			continue
		}

		req := &exampb.GetExamResultRequest{
			StudentId: parts[0],
			ExamId:    parts[1],
		}

		if err := stream.Send(req); err != nil {
			log.Printf("send error: %v", err)
			break
		}
	}

	<-done
	fmt.Println("👋 Session ended.")

}
