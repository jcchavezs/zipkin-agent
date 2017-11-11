package zipkinagent

import (
	"log"
	"sync"
	"time"
)

const defaultBatchInterval = 1
const defaultBatchSize = 100
const defaultMaxBacklog = 1000

type Collector struct {
	logger        log.Logger
	transporter   Transporter
	batchInterval time.Duration
	batchSize     int
	maxBacklog    int
	batch         *[]Span
	spansChan     chan *[]Span
	quit          chan struct{}
	shutdown      chan error
	sendMutex     *sync.Mutex
	batchMutex    *sync.Mutex
}

func NewCollector(t Transporter) (*Collector, error) {
	c := &Collector{
		transporter:   t,
		logger:        log.Logger{},
		batchInterval: defaultBatchInterval * time.Second,
		batchSize:     defaultBatchSize,
		maxBacklog:    defaultMaxBacklog,
		batch:         &[]Span{},
		spansChan:     make(chan *[]Span),
		quit:          make(chan struct{}, 1),
		shutdown:      make(chan error, 1),
		sendMutex:     &sync.Mutex{},
		batchMutex:    &sync.Mutex{},
	}

	go c.loop()

	return c, nil
}

func (c *Collector) Collect(s *[]Span) error {
	c.spansChan <- s
	return nil
}

func (c *Collector) Close() error {
	close(c.quit)
	return <-c.shutdown
}

func (c *Collector) loop() {
	var (
		nextSend = time.Now().Add(c.batchInterval)
		ticker   = time.NewTicker(c.batchInterval / 10)
		tickc    = ticker.C
	)
	defer ticker.Stop()

	for {
		select {
		case spans := <-c.spansChan:
			currentBatchSize := c.append(spans)
			if currentBatchSize >= c.batchSize {
				nextSend = time.Now().Add(c.batchInterval)
				go c.send()
			}
		case <-tickc:
			if time.Now().After(nextSend) {
				nextSend = time.Now().Add(c.batchInterval)
				go c.send()
			}
		case <-c.quit:
			c.shutdown <- c.send()
			return
		}
	}
}

func (c *Collector) append(spans *[]Span) (newBatchSize int) {
	c.batchMutex.Lock()
	defer c.batchMutex.Unlock()

	*c.batch = append(*c.batch, *spans...)
	if len(*c.batch) > c.maxBacklog {
		dispose := len(*c.batch) - c.maxBacklog
		c.logger.Printf("backlog too long, disposing %d spans.\n", dispose)
		*c.batch = (*c.batch)[dispose:]
	}
	newBatchSize = len(*c.batch)
	return
}

func (c *Collector) send() error {
	// in order to prevent sending the same batch twice
	c.sendMutex.Lock()
	defer c.sendMutex.Unlock()

	// Select all current spans in the batch to be sent
	c.batchMutex.Lock()
	sendBatch := (*c.batch)[:]
	c.batchMutex.Unlock()

	// Do not send an empty batch
	if len(sendBatch) == 0 {
		return nil
	}

	c.transporter.Send(&sendBatch)

	// Remove sent spans from the batch
	c.batchMutex.Lock()
	*c.batch = (*c.batch)[len(sendBatch):]
	c.batchMutex.Unlock()

	return nil
}
