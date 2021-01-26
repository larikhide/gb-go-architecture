package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

type Job struct {
	payload []byte
}

type Worker struct {
	wg      *sync.WaitGroup
	num     int // only for example
	jobChan <-chan *Job
}

func main() {
	var numThreads, numJobs, workTime int
	flag.IntVar(&numThreads, "c", 1, "number of threads")
	flag.IntVar(&numJobs, "n", 1, "number of requests")
	flag.IntVar(&workTime, "t", 1, "time of work in seconds")

	flag.Parse()

	if numJobs > 1 && workTime > 1 {
		log.Fatalln("specify only one flag 't' or 'n'")
	}

	if numJobs < 1 || workTime < 1 {
		log.Fatalln("flags 't' or 'n' should be more or equal than 1")
	}

	wg := &sync.WaitGroup{}
	jobChan := make(chan *Job)
	for i := 0; i < numThreads; i++ { //тут создается 5 воркеров
		worker := NewWorker(i+1, wg, jobChan)
		wg.Add(1)
		go worker.Handle()
	}
	// отметка по времени, когда началось выполнение джоб
	start := time.Now()

	isNumJobsPassed := numJobs != 1 // проверка был ли введен флаг n
	for i := 0; i < numJobs; i++ {
		if !isNumJobsPassed && time.Since(start) < time.Duration(workTime)*time.Second {
			numJobs++ //если был задан флаг на время выполнения, а не количество запросов, то в течении этого времени плодятся новые джобы
		}
		jobChan <- &Job{
			payload: []byte(fmt.Sprintf("Some message %d", i)),
		}
	}

	close(jobChan)
	wg.Wait()

	fmt.Println(time.Since(start))
	fmt.Printf("%.2f RPS\n", float64(numJobs)/float64(time.Since(start)))
}

// Handle просто шаблон
func (w *Worker) Handle() {
	defer w.wg.Done()
	for job := range w.jobChan {
		log.Printf("worker %d processing job with payload %s", w.num, string(job.payload))
	}
}

func (w *Worker) HandleDDoS() {
	defer w.wg.Done()
	for job := range w.jobChan {
		resp, err := http.Get("https://google.com")
		if err != nil {
			fmt.Println(err)
			continue
		}
		log.Printf("response is %s and job = %s ", resp.Header, string(job.payload))
		defer resp.Body.Close()
	}
}

func NewWorker(num int, wg *sync.WaitGroup, jobChan <-chan *Job) *Worker {
	return &Worker{
		wg:      wg,
		num:     num,
		jobChan: jobChan,
	}
}
