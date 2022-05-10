package bucket

import (
	"sync"
	"time"
)

// Counter - store the counter of checks and start time.
type Counter struct {
	count   int
	created time.Time
	updated time.Time
}

// Bucket - store map of Counter's.
type Bucket struct {
	prefix   string
	limit    int
	interval time.Duration
	ttl      time.Duration
	list     map[string]*Counter
	mx       sync.Mutex
}

func NewBucket(prefix string, limit int, interval, ttl time.Duration) *Bucket {
	return &Bucket{
		prefix:   prefix,
		limit:    limit,
		interval: interval,
		ttl:      ttl,
		list:     make(map[string]*Counter),
	}
}

func (b *Bucket) Allow(key string) bool {
	b.mx.Lock()
	currentTime := time.Now()
	counter, ok := b.list[key]
	defer b.mx.Unlock()

	//	New key for checking
	if !ok {
		counter = &Counter{
			count:   0,
			created: currentTime,
			updated: currentTime,
		}
		b.list[key] = counter
	}

	//	Reset counter for expired time interval
	if currentTime.Sub(counter.created) > b.interval {
		counter.count = 0
		counter.created = currentTime
		counter.updated = currentTime
	}

	//	Counter was out of limit for interval
	if counter.count >= b.limit {
		return false
	}

	// Update counter
	counter.count++
	counter.updated = time.Now()

	return true
}

//	Reset Counter in Bucket
func (b *Bucket) Reset(key string) {
	b.mx.Lock()
	currentTime := time.Now()
	counter, ok := b.list[key]
	defer b.mx.Unlock()

	//	New key for checking
	if !ok {
		counter = &Counter{
			count:   0,
			created: currentTime,
			updated: currentTime,
		}
		b.list[key] = counter
		return
	}

	counter.count = 0
	counter.created = currentTime
	counter.updated = currentTime
}

// BucketGC - garbage collector for Bucket. Delete items from list if last update of counter more that TTL.
func (b *Bucket) BucketGC() {
	b.mx.Lock()
	currentTime := time.Now()
	defer b.mx.Unlock()

	for key, counter := range b.list {
		if currentTime.Sub(counter.updated) > b.ttl {
			delete(b.list, key)
		}
	}
}
