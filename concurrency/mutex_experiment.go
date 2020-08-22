package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyz" + "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()),
)

func RandString(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

type safeMap struct {
	Map map[string]string
	mux sync.Mutex
}

func (m *safeMap) safeGet(key string) string {
	m.mux.Lock()
	rv := m.Map[key]
	m.mux.Unlock()
	return rv
}

func (m *safeMap) safeSet(key string, value string) {
	m.mux.Lock()
	m.Map[key] = value
	m.mux.Unlock()
}

func randomKeyValueMutator(m *safeMap, key string) {
	m.safeSet(key, RandString(10))
}

func main() {
	var mux sync.Mutex
	var wg sync.WaitGroup
	kool := safeMap{
		make(map[string]string),
		mux,
		// the zero val for a mutex is an unlocked mutex
	}
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() { randomKeyValueMutator(&kool, "nice"); wg.Done() }()
	}
	wg.Wait()
	fmt.Println(kool.safeGet("nice"))
}
