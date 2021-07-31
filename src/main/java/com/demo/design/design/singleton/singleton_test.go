package singleton

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

/**
 * 懒汉,线程安全，延迟加载
 */
type SingletonA struct {
}

var mu sync.Mutex
var instanceA *SingletonA

func GetInstanceA() *SingletonA {
	mu.Lock()
	defer mu.Unlock()
	if instanceA == nil {
		instanceA = &SingletonA{}
	}
	return instanceA
}

func TestSingletonA(t *testing.T) {
	start := time.Now()
	wg := sync.WaitGroup{}
	wg.Add(100)
	instances := [100]*SingletonA{}
	for i := 0; i < 100; i++ {
		go func(index int) {
			instances[index] = GetInstanceA()
			wg.Done()
		}(i)
	}
	wg.Wait()
	for i := 0; i < 100; i++ {
		fmt.Printf("地址：%p\n", instances[i])
	}
	since := time.Since(start)
	fmt.Println("耗时：", since)
}

type SingletonB struct {
}

var instanceB *SingletonB
var once sync.Once

func GetInstanceB() *SingletonB {
	once.Do(func() {
		instanceB = &SingletonB{}
	})
	return instanceB
}

func TestSingletonB(t *testing.T) {
	start := time.Now()
	wg := sync.WaitGroup{}
	wg.Add(100)
	instances := [100]*SingletonB{}
	for i := 0; i < 100; i++ {
		go func(index int) {
			instances[index] = GetInstanceB()
			wg.Done()
		}(i)
	}
	wg.Wait()
	for i := 0; i < 100; i++ {
		fmt.Printf("地址：%p\n", instances[i])
	}
	since := time.Since(start)
	fmt.Println("耗时：", since)
}
