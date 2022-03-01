/**
    package: sco_tracers
    filename: kafka
    author: diogo@gmail.com
    time: 2022/3/1 22:33
**/
package kafka

import (
	"fmt"

	"github.com/Shopify/sarama"
)

// 方案一:  这是最简单的方案
//func Producer(topic string, data string) error {
//	config := sarama.NewConfig()
//	config.Producer.RequiredAcks = sarama.WaitForAll
//	config.Producer.Partitioner = sarama.NewRandomPartitioner
//	// config.Producer.Return.Successes = true // 只有同步消息才设置为true
//	producer, err := sarama.NewAsyncProducer([]string{""}, config)
//	if err != nil {
//		return err
//	}
//	defer producer.AsyncClose()
//
//	producer.Input() <- &sarama.ProducerMessage{Topic: topic, Key: nil, Value: sarama.StringEncoder(data)}
//	select {
//
//	// 同步消息才会接收到success
//	/*case msg := <-producer.Successes():
//	log.Printf("Produced message successes: [%s]\n", msg.Value)*/
//	case err := <-producer.Errors():
//		log.Println("Produced message failure: ", err)
//	default:
//		log.Println("Produced message default")
//	}
//	if err != nil {
//		return err
//	}
//	return nil
//}
// - ----------------------------

// -------------- 方案二 ------

// Producer out
var Producer *KafkaProducer

func InitProducer(kafkaUrls []string, topic string, config *sarama.Config) {
	p, err := NewKafkaProducer(kafkaUrls, topic, config)
	if err != nil {
		fmt.Println("New Kafka Producer Failed.")
	} else {
		Producer = p
	}
}

type KafkaProducer struct {
	sarama.AsyncProducer
	Topic string
}

func NewKafkaProducer(kafkaUrls []string, topic string, config *sarama.Config) (*KafkaProducer, error) {
	asyncProducer, err := sarama.NewAsyncProducer(kafkaUrls, config)
	if err != nil {
		return nil, err
	}

	// todo: check 看情况
	if err := NewTopic(kafkaUrls, topic, 3); err != nil {
		fmt.Printf("NewKafkaProducer: %v\n", err)
	}

	kp := &KafkaProducer{asyncProducer, topic}

	go func() {
		for {
			select {
			case suc := <-kp.Successes():
				fmt.Printf(
					"send msg to kafka topic successfully. partition: %d, offset: %d, timestamp: %s\n",
					suc.Partition,
					suc.Offset,
					suc.Timestamp.String(),
				)
			case fail := <-kp.Errors():
				fmt.Printf(
					"send msg to kafka topic fail. error: %s",
					fail.Err.Error(),
				)
			}
		}
	}()
	return kp, nil
}

func (kp *KafkaProducer) MakeMsg(key, value string) *sarama.ProducerMessage {
	msg := &sarama.ProducerMessage{
		Topic: kp.Topic,
		Key:   sarama.StringEncoder(key),
		//Value: sarama.ByteEncoder(value),
		Value: sarama.StringEncoder(value),
	}
	return msg
}

func (kp *KafkaProducer) SendMsg(key, value string) error {
	msg := kp.MakeMsg(key, value)
	kp.Input() <- msg
	return nil
}
