package servers

import (
	"fmt"
	"time"

	"github.com/aakash19here/grpc_exam/proto/generated/exampb"
)

func (s *ExamServiceServer) StreamExamResults(req *exampb.StreamExamResultsRequest, stream exampb.ExamService_StreamExamResultsServer) error {
	studentId := req.StudentId
	examIDs := req.ExamIds

	found := false

	for _, examID := range examIDs {
		key := fmt.Sprintf("%s_%s", studentId, examID)
		if result, ok := s.examData[key]; ok {
			time.Sleep(time.Second) // artifical delay
			stream.Send(result)

			found = true
		}

		if !found {
			return fmt.Errorf("exam results not found for student id %s and examIds: %v", studentId, examIDs)
		}
	}

	return nil
}
