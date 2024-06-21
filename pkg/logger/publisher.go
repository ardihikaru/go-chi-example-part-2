package logger

// publisher provides the interface for the functionality of publisher (message broker)
//
//	FYI: to use this interface, we need to implement it somewhere in our code. :)
type publisher interface {
	PublishLogToLogManager(logType string, msg []byte) error    // FYI: Log Manager will handle how it will be sent into another channel
	PublishLogToElasticSearch(logType string, msg []byte) error // publishes to the elasticsearch database
	PublishLogToMessageBroker(logType string, msg []byte) error // publishes to the message broker (e.g. RabbitMQ)
}
