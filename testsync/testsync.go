package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	//"time"

	"runtime"
)

var result int32

func setResult(r int32) {
	atomic.StoreInt32(&result, r)
	//result = r
}

func read(v *int32, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10000; i++ {
		if *v == 1 {
			return
		}
		setResult(0)
		//time.Sleep(10 * time.Millisecond)
	}
}

func write(v *int32, wg *sync.WaitGroup) {
	defer wg.Done()
	count := 0
	for i := 0; i < 10000; i++ {
		if count >= 8000 {
			*v = 1
			setResult(1)
			return
		}
		count++
		//time.Sleep(10 * time.Millisecond)
	}
}

func readAtomic(v *int32, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10000; i++ {
		if 1 == atomic.LoadInt32(v) {
			//fmt.Printf("read break\n")
			return
		}
		setResult(0)
		//time.Sleep(10 * time.Millisecond)
	}
}

func writeAtomic(v *int32, wg *sync.WaitGroup) {
	defer wg.Done()
	count := 0
	for i := 0; i < 10000; i++ {
		if count >= 100 {
			atomic.StoreInt32(v, 1)
			setResult(1)
			return
		}
		count++
		//time.Sleep(10 * time.Millisecond)
	}
}

func readSync(v *int32, wg *sync.WaitGroup, lock sync.Locker) {
	defer wg.Done()
	for i := 0; i < 10000; i++ {
		lock.Lock()
		if 1 == *v {
			lock.Unlock()
			return
		}
		setResult(0)
		lock.Unlock()
		//time.Sleep(10 * time.Millisecond)
	}
}

func writeSync(v *int32, wg *sync.WaitGroup, lock sync.Locker) {
	defer wg.Done()
	count := 0
	for i := 0; i < 10000; i++ {
		if count >= 100 {
			lock.Lock()
			*v = 1
			setResult(1)
			lock.Unlock()
			return
		}
		count++
		//time.Sleep(10 * time.Millisecond)
	}
}

func testNonAtomic() {
	var v int32 = 0
	var wg sync.WaitGroup

	wg.Add(5)
	go read(&v, &wg)
	go read(&v, &wg)
	go write(&v, &wg)
	go read(&v, &wg)
	go read(&v, &wg)

	wg.Wait()
	//fmt.Printf("result:%v\n", result)
}

func testAtomic() {
	var v int32 = 0
	var wg sync.WaitGroup

	wg.Add(5)
	go readAtomic(&v, &wg)
	go readAtomic(&v, &wg)
	go writeAtomic(&v, &wg)
	go readAtomic(&v, &wg)
	go readAtomic(&v, &wg)

	wg.Wait()
	//fmt.Printf("result:%v\n", result)
}

func testSync() {
	var v int32 = 0
	var wg sync.WaitGroup
	var mu sync.Mutex

	wg.Add(5)
	go readSync(&v, &wg, &mu)
	go readSync(&v, &wg, &mu)
	go writeSync(&v, &wg, &mu)
	go readSync(&v, &wg, &mu)
	go readSync(&v, &wg, &mu)

	wg.Wait()
	//fmt.Printf("result:%v\n", result)
}

func main() {
	runtime.GOMAXPROCS(4)

	setResult(0)
	fmt.Println("testing non atomic")
	for i := 0; i < 100; i++ {
		testNonAtomic()
		if result == 0 {
			fmt.Println("failed")
		}
	}

	fmt.Println("testing atomic")
	setResult(0)
	for i := 0; i < 100; i++ {
		testAtomic()
		if result == 0 {
			fmt.Println("failed")
		}
	}

	fmt.Println("testing sync")
	setResult(0)
	for i := 0; i < 100; i++ {
		testSync()
		if result == 0 {
			fmt.Println("failed")
		}
	}
}
