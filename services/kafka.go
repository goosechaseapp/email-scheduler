package services

import (
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/confluentinc/confluent-kafka-go/v2/schemaregistry"
	"github.com/confluentinc/confluent-kafka-go/v2/schemaregistry/serde"
	"github.com/confluentinc/confluent-kafka-go/v2/schemaregistry/serde/protobuf"
	"goosechase.ai/email-scheduler/config"
	"goosechase.ai/email-scheduler/proto/proto"
	"goosechase.ai/email-scheduler/util/log"
)

type kafkaInstance struct {
	producer   *kafka.Producer
	serializer *protobuf.Serializer
}

var Kafka kafkaInstance

func (k *kafkaInstance) InitializeConfluentKafka() {
	bootstrapServers := config.Env("KAFKA_BOOTSTRAP_SERVER")
	schemaRegistryUrl := config.Env("SCHEMA_REGISTRY_URL")

	producer, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": bootstrapServers})
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create Kafka producer")
	}

	log.Info().Msg("Kafka producer created")

	schemaRegistry, err := schemaregistry.NewClient(schemaregistry.NewConfig(schemaRegistryUrl))
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create Schema Registry client")
	}

	log.Info().Msg("Schema Registry client created")

	serializer, err := protobuf.NewSerializer(schemaRegistry, serde.ValueSerde, protobuf.NewSerializerConfig())
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create Protobuf serializer")
	}

	log.Info().Msg("Protobuf serializer created")

	k.producer = producer
	k.serializer = serializer
}

func (k *kafkaInstance) produce(topic string, payload []byte) error {
	deliveryChan := make(chan kafka.Event)

	err := k.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          payload,
		Key:            []byte(string(rune(time.Now().Unix()))),
	}, deliveryChan)

	if err != nil {
		log.Error().Err(err).Msg("Failed to produce message")
		return err
	}

	e := <-deliveryChan
	m := e.(*kafka.Message)

	close(deliveryChan)

	if m.TopicPartition.Error != nil {
		log.Error().Err(m.TopicPartition.Error).Msg("Failed to deliver message")
		return m.TopicPartition.Error
	}

	log.Info().Str("topic", *m.TopicPartition.Topic).Str("partition", string(m.TopicPartition.Partition)).Int64("offset", int64(m.TopicPartition.Offset)).Msg("Message delivered")
	return nil

}

func (k *kafkaInstance) ProduceEmail(email *proto.SendEmailDocument) error {
	topic := config.Env("KAFKA_TOPIC")

	payload, err := k.serializer.Serialize(topic, email)
	if err != nil {
		log.Error().Err(err).Msg("Failed to serialize email")
		return err
	}

	return k.produce(topic, payload)
}

func (k *kafkaInstance) Close() {
	if k.producer != nil {
		k.producer.Close()
	}
}
