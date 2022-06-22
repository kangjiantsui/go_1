package lmstfy

import (
	"fmt"
	"github.com/bitleak/lmstfy/client"
	"go_01/utils"
	"os"
	"strconv"
	"testing"
)

var c *client.LmstfyClient

func TestMain(m *testing.M) {
	c = client.NewLmstfyClient("localhost", 7777, "test-ns", "01G5K9SRA0C50MQ28P5FH9SWWD")
	c.ConfigRetry(3, 50)
	code := m.Run()
	os.Exit(code)
}

func TestAck(t *testing.T) {
	_ = c.Ack("q1", "01G5RFC9YQM9DRBYTMQNHG0000")
}

func TestPeekQueue(t *testing.T) {
	var job *client.Job
	job, _ = c.PeekQueue("q1")
	if job == nil {
		return
	}
	t.Log(job.ID)
}

func TestPublish(t *testing.T) {
	delay := uint32(99)
	jobID, err := c.Publish("q1", []byte(strconv.Itoa(int(delay))), 0, 3, delay)
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Log("jobID", jobID)
}

func TestBatchPublish(t *testing.T) {
	jobId, err := c.BatchPublish("q1", []interface{}{"1", "2"}, 0, 4, 5)
	if err != nil {
		return
	}
	t.Log(jobId)
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
