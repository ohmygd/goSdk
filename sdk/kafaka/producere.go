package kafaka

import (
	"fmt"
	"github.com/Shopify/sarama"
	"time"
)

// 生产者

func Producer() {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner

	producer, err := sarama.NewSyncProducer([]string{"172.16.19.140:9092"}, config)
	if err != nil {
		panic(err)
	}
	defer producer.Close()

	go func(){

		msg := &sarama.ProducerMessage{
			Topic:     "go-test",
			Partition: int32(1),
			Key:       sarama.StringEncoder("key1"),
		}

		for {
			msg.Value = sarama.ByteEncoder(time.Now().String() + "----------")
			paritition, offset, err := producer.SendMessage(msg)
			if err != nil {
				fmt.Println("Send Message Fail")
			}
			fmt.Printf("Partion = %d, offset = %d\n", paritition, offset)
			time.Sleep(time.Second)
		}
	}()

	go func(){

		msg := &sarama.ProducerMessage{
			Topic:     "go-test1",
			Partition: int32(10),
			Key:       sarama.StringEncoder("key2"),
		}

		for {
			msg.Value = sarama.ByteEncoder(time.Now().String() + "mc")
			paritition, offset, err := producer.SendMessage(msg)
			if err != nil {
				fmt.Println("Send Message Fail")
			}
			fmt.Printf("Partion = %d, offset = %d\n", paritition, offset)
			time.Sleep(time.Second)
		}
	}()

	select{}
}

func SendInfo(id int) {

}