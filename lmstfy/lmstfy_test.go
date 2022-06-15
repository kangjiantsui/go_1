package lmstfy

import (
	"fmt"
	"github.com/bitleak/lmstfy/client"
	"go_01/utils"
	"os"
	"testing"
)

var c *client.LmstfyClient

func TestMain(m *testing.M) {
	c = client.NewLmstfyClient("localhost", 7777, "test-ns", "01G5K9SRA0C50MQ28P5FH9SWWD")
	code := m.Run()
	os.Exit(code)
}

func TestPublish(t *testing.T) {
	c.ConfigRetry(3, 50)
	jobID, err := c.Publish("q1", []byte("hello"), 0, 3, 5)
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Log("jobID", jobID)
}

func TestConsume(t *testing.T) {
	for {
		job, _ := c.Consume("q1", 10, 12)
		if job != nil {
			t.Log(utils.String(job.Data))
			_ = c.Ack("q1", job.ID)
		}
	}
}

func TestBatchConsume(t *testing.T) {
	jobs, err := c.BatchConsume([]string{"q1"}, 5, 10, 12)
	if err != nil {
		panic(err)
	}

	for _, job := range jobs {
		fmt.Println(string(job.Data))
		err := c.Ack("q1", job.ID)
		if err != nil {
			panic(err)
		}
	}
}
