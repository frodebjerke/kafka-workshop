package challenges

import (
    "testing"

    "github.com/frodebjerke/kafka-workshop/setup"
)

func TestExercise1(t *testing.T) {
    consumer := setup.GetConsumer()

    partitionConsumer, err := consumer.ConsumePartition("ex1", 0, -1)

    if err != nil {
        t.Errorf("Error creating partitionConsumer", err)
        return
    }

    message := <- partitionConsumer.Messages()
    val := string(message.Value)
    if val != "hello" {
        t.Errorf("Expected message to be hello got %v", val)
    }

    partitionConsumer.Close()
    consumer.Close()
}
