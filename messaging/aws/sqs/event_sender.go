package messaging

import (
	"encoding/json"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
)

type SqsEventSender struct {
	Client   *sqs.SQS
	QueueURL string
}

func (receiver *SqsEventSender) SendEvent(event interface{}) error {
	body, err := json.Marshal(event)
	if err == nil {
		input := sqs.SendMessageInput{
			MessageBody: aws.String(string(body)),
			QueueUrl:    &receiver.QueueURL,
		}
		_, err = receiver.Client.SendMessage(&input)
	}
	return err
}
