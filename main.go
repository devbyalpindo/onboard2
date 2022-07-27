package main

import (
	"backend-onboarding2/config"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"os/signal"
)

func main() {
	consumerKafka, err := config.InitConsumer()
	if err != nil {
		log.Fatalln("creating Kafka Consumer: ", err)
	}
	defer consumerKafka.Close()

	schemaKafka, errSchema := config.GetLastSchema("users")
	if errSchema != nil {
		log.Fatalln("schema user: ", errSchema)
	}
	hbase, errHbase := config.ConnectHbase()
	if errHbase != nil {
		log.Print("error hbase connection:", errHbase)
	}
	defer hbase.Close()
	if errHbase == nil {
		router := gin.Default()

		go func() {
			if err := router.Run(":8000"); err != nil {
				log.Fatalf("listen: %s\n", err)
			}
		}()

		config.ListenerKafka(consumerKafka, schemaKafka, hbase)
		quit := make(chan os.Signal, 1)

		signal.Notify(quit, os.Interrupt)
		<-quit
		log.Println("Shutdown Server...")

	}

}
