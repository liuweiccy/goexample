package event

// go get github.com/asaskevich/EventBus
import (
	"fmt"
	"github.com/asaskevich/EventBus"
	"sync"
	"testing"
	"time"
)

const N int  = 10000000

func TestEventBus(t *testing.T) {
	eventBus := EventBus.New()
	wait := new(sync.WaitGroup)
	wait.Add(N)

	s := time.Now()
	_ = eventBus.Subscribe("vv", func(wait *sync.WaitGroup) {
		wait.Done()
	})

	err := eventBus.Subscribe("test", func(wait *sync.WaitGroup) {
		wait.Done()
	})

	if err != nil {
		fmt.Println("发生错误：", err)
	}

	for i := 0; i < N/2; i++ {
		eventBus.Publish("vv", wait)
		eventBus.Publish("test", wait)
	}

	wait.Wait()


	fmt.Println(time.Now().Sub(s))
}

func BenchmarkEventBus(b *testing.B) {
	eventBus := EventBus.New()
	_ = eventBus.Subscribe("vv", func(msg string) {

	})
	for i := 0; i < b.N; i++ {
		eventBus.Publish("vv", "test1")
	}
}