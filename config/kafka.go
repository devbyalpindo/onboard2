package config

import (
	"backend-onboarding2/model/entity"
	"backend-onboarding2/repository/user_repository"
	"encoding/json"
	"fmt"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/jinzhu/gorm"
	"github.com/riferrei/srclient"
)

func InitConsumer() (*kafka.Consumer, error) {
	consume, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"group.id":          "onboarding",
		"auto.offset.reset": "latest",
	})
	if err != nil {
		return nil, err
	}

	kafkaTopics := []string{"users"}
	if err = consume.SubscribeTopics(kafkaTopics, nil); err != nil {
		return nil, err
	}

	return consume, nil
}

func GetLastSchema(topic string) (*entity.Schema, error) {

	var schemaResult = entity.Schema{}
	topic = topic + "-value"
	schemaRegistryClient := srclient.CreateSchemaRegistryClient("http://localhost:8081")

	schema, errSchema := schemaRegistryClient.GetLatestSchema(topic)
	if errSchema != nil {
		return nil, errSchema
	}
	schemaResult.SchemaUser = schema

	return &schemaResult, nil
}

func ListenerKafka(consumer *kafka.Consumer, schema *entity.Schema, hbaseConn *gorm.DB) {
	log.Printf("Listening Message...")
	for {
		msg, err := consumer.ReadMessage(-1)
		if err != nil {
			log.Printf("Error %v\n", err.Error())
			continue
		}

		var userModel *entity.User

		fmt.Println(msg.Value)

		var errJson = json.Unmarshal(msg.Value, &userModel)
		if errJson != nil {
			log.Printf("json %v\n", errJson.Error())
		}
		userRepository := user_repository.NewUserRepository(hbaseConn)

		switch userModel.Operation {
		case "create":
			log.Print("create user")
			userRepository.AddUsers(userModel)
		case "update":
			log.Print("update user")
			userRepository.AddUsers(userModel)
			log.Print("operation not found")
		}

	}
}
