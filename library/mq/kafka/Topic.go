/**
    package: sco_tracers
    filename: kafka
    author: diogo@gmail.com
    time: 2022/3/1 23:01
**/
package kafka

import (
	"fmt"

	"github.com/Shopify/sarama"
)

func NewTopic(kafkaUrls []string, topicName string, partitionNum int32) error {
	config := sarama.NewConfig()
	// 看一下他的那个版本
	config.Version = sarama.V2_0_0_0
	admin, err := sarama.NewClusterAdmin(kafkaUrls, config)
	if err != nil {
		return fmt.Errorf("new kafka admin err: %v", err)
	}
	topicDetail := &sarama.TopicDetail{
		NumPartitions:     partitionNum,
		ReplicationFactor: 1,
	}
	err = admin.CreateTopic(topicName, topicDetail, false)
	if err != nil && err != sarama.ErrTopicAlreadyExists {
		return fmt.Errorf("create topic %s err: %v", topicName, err)
	}

	err = admin.Close()
	if err != nil {
		return fmt.Errorf("close kafka admin err: %v", err)
	}
	return nil
}
