resource "aws_sqs_queue" "queue_staging" {
  name = "TESTING_queue"
  message_retention_seconds = 300
}

output "sqs_queue" {
  value = aws_sqs_queue.queue_staging
}