package amqp_manager

import (
	"context"
	"fmt"
	"log"
	"pack.ag/amqp"
)

type AMQPConsumer struct {
	manager  *AMQPManager
	client   *amqp.Client
	session  *amqp.Session
	receiver *amqp.Receiver
}

func NewAMQPConsumer(manager *AMQPManager) *AMQPConsumer {
	return &AMQPConsumer{
		manager: manager,
	}
}

func (consumer *AMQPConsumer) generateReceiver() error {
	receiver, err := consumer.session.NewReceiver(amqp.LinkCredit(20))
	if err == nil {
		consumer.receiver = receiver
	}

	return err
}

func (consumer *AMQPConsumer) init() (err error) {
	if consumer.session != nil {
		err = consumer.generateReceiver()
		if err == nil {
			return
		}
	}

	if consumer.client != nil {
		_ = consumer.client.Close()
	}

	manager := consumer.manager
	client, err := amqp.Dial(manager.address, amqp.ConnSASLPlain(manager.username, manager.password))
	if err != nil {
		return
	}
	consumer.client = client

	session, err := client.NewSession()
	if err != nil {
		return
	}
	consumer.session = session

	return consumer.generateReceiver()
}

func (consumer *AMQPConsumer) Start(ctx context.Context, onMessage func(*amqp.Message)) (err error) {
	err = consumer.init()
	if err != nil {
		return
	}
	fmt.Println("AMQP connect success")

	msgChan, errChan := consumer.run(ctx)

	go func() {
		for {
			select {
			case msg := <-msgChan:
				onMessage(msg)
			case err = <-errChan:
				log.Println(err)
				return
			}
		}
	}()

	return
}

func (consumer *AMQPConsumer) run(ctx context.Context) (chan *amqp.Message, chan error) {
	msgChan := make(chan *amqp.Message, 10)
	errChan := make(chan error, 10)

	go func() {
		childCtx, cancel := context.WithCancel(ctx)
		defer func() {
			_ = consumer.receiver.Close(childCtx)
			_ = consumer.session.Close(childCtx)
			_ = consumer.client.Close()
		}()

		for {
			message, err := consumer.receiver.Receive(ctx)
			if err != nil {
				errChan <- err
				cancel()
				break

			} else {
				if message != nil {
					msgChan <- message
					err = message.Accept()
					if err != nil {
						errChan <- err
					}
				}
			}
		}
	}()

	return msgChan, errChan
}
