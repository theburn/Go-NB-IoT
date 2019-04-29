package amqp

import (
	"Go-NB-IoT/configure"
	log "Go-NB-IoT/logging"
	"time"

	"github.com/streadway/amqp"
)

var amqpChannel *amqp.Channel
var amqpConn *amqp.Connection
var amqpQueue amqp.Queue

const (
	DefaultQueueName   = "NBIoTCallback"
	DefaultContentType = "application/json"
)

func InitAMQP() error {
	var err error
	amqp_url := configure.NBIoTConfig.AMQPParam.AMQPURL
	log.Debugf("amqp connect to %s", amqp_url)
	dialFlag := false
	// reconnect 10 times
	for i := 0; i < 10; i++ {
		amqpConn, err = amqp.Dial(amqp_url)
		if err != nil {
			log.Errorf("amqp Dial error: %s, conntinue retry...", err.Error())
			time.Sleep(5 * time.Second)
		} else {
			dialFlag = true
			break
		}
	}

	if !dialFlag {
		log.Errorf("amqp Dial error, maybe rabbitmq is down, return and quit")
		return err
	}

	amqpChannel, err = amqpConn.Channel()

	if err != nil {
		log.Errorf("amqp init Channel Error %s", err.Error())
		return err
	}

	return nil
}

func CloseConn() error {
	return amqpConn.Close()
}

func CloseChannel() error {
	return amqpChannel.Close()
}

func InitQueue(queueName string) error {

	var err error

	amqpQueue, err = amqpChannel.QueueDeclare(
		queueName, // name
		true,      // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)

	if err != nil {
		log.Errorf("amqp init NBIoTCallback Queue Error %s", err.Error())
		return err
	}

	return nil
}

func AMQPSend(queueName, contentType string, v []byte) error {
	// No Check Queue is exists
	// Fixme ?
	return amqpChannel.Publish(
		"",        // examqpChannelange
		queueName, // routing key
		false,     // mandatory
		false,     // immediate
		amqp.Publishing{
			ContentType: contentType,
			Body:        v,
		})
}

func AMQPRecv(queueName string) (<-chan amqp.Delivery, error) {
	// No Check Queue is exists
	// Fixme ?
	return amqpChannel.Consume(
		queueName, // queue
		"",        // consumer
		true,      // auto-ack
		false,     // exclusive
		false,     // no-local
		false,     // no-wait
		nil,       // args
	)
}
