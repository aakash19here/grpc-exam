package clients

import (
	"context"
	"fmt"
	"time"

	"github.com/aakash19here/grpc_exam/proto/generated/exampb"
)

func Unary(client exampb.ExamServiceClient) {
	fmt.Println("Enter student ID and exam ID (e.g., 123 math101):")
	var studentId, examId string

	fmt.Scanf("%s %s", &studentId, &examId)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	resp, err := client.GetExamResult(ctx, &exampb.GetExamResultRequest{
		StudentId: studentId,
		ExamId:    examId,
	})

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Student Name: %s\n", resp.StudentName)
	fmt.Printf("Subject: %s\n", resp.Subject)
	fmt.Printf("Marks Obtained: %d out of %d\n", resp.MarksObtained, resp.TotalMarks)
	fmt.Printf("Grade: %s\n", resp.Grade)
	fmt.Println("Unary RPC call completed successfully.")
}
