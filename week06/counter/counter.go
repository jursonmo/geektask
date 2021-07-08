package counter

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
	"unsafe"
)

const (
	Success = iota
	Fail
	Timeout
	MaxEventType
)

type BucketCounter struct {
	sync.Once
	startTime  time.Time
	totalTime  time.Duration
	windowTime time.Duration //一个采样窗口的时间
	buckets    []Bucket
}

type Bucket struct {
	counters [MaxEventType]Counter //可以记录多种事件的计数, 每种事件计数之间避免伪共享
}

type count = int64

//avoid false sharing
type Counter struct {
	count
	pad [64 - unsafe.Sizeof(count(0))]byte
	//internal/cpu
	//pad   [cpu.CacheLinePadSize - unsafe.Sizeof(int64(0))%cpu.CacheLinePadSize]byte
}

//totalTime: 持续统计时间
//bucketSizeInMs: 多少毫秒一个采样窗口
func NewBucketCounter(totalTime time.Duration, bucketSizeInMs int) *BucketCounter {
	if bucketSizeInMs == 0 {
		bucketSizeInMs = 100 //
	}
	windowTime := time.Duration(bucketSizeInMs) * time.Millisecond
	size := int(totalTime / windowTime)

	return &BucketCounter{
		totalTime:  totalTime,
		windowTime: windowTime,
		buckets:    make([]Bucket, size+1), //size + 1 ,避免wrap覆盖
	}
}

func (bc *BucketCounter) SetStartTimeOnce() {
	bc.Do(func() {
		bc.startTime = time.Now()
	})
}

func (bc *BucketCounter) ReceiveEvent(eventType int, value int) error {
	if eventType >= MaxEventType {
		return fmt.Errorf("eventType(%d) >= MaxEventType(%d)", eventType, MaxEventType)
	}
	bc.SetStartTimeOnce()
	elapse := time.Since(bc.startTime)
	index := int(int64(elapse/bc.windowTime) % int64(len(bc.buckets)))
	c := &bc.buckets[index].counters[eventType]
	atomic.AddInt64(&c.count, count(value))
	return nil
}

func (bc *BucketCounter) GetValue(eventType int) int64 {
	if eventType >= MaxEventType {
		return 0
	}
	elapse := time.Since(bc.startTime)
	currentIndex := int(int64(elapse/bc.windowTime) % int64(len(bc.buckets)))

	//有size+1个buckets(窗口), 统计非当前窗口所有值相加
	result := int64(0)
	for i := 0; i < len(bc.buckets); i++ {
		if i == currentIndex {
			continue
		}
		c := &bc.buckets[i].counters[eventType]
		result += atomic.LoadInt64(&c.count)
	}
	return result
}
