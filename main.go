package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
)

func main() {
	http.HandleFunc("/", handler)

	fmt.Println("ListenAndServe")
	http.ListenAndServe(":8000", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	bytes, err := io.ReadAll(r.Body)
	ep(err)

	m := map[string]string{}
	err = json.Unmarshal(bytes, &m)
	ep(err)

	w.WriteHeader(200)
	w.Header().Add("Content-Type", "audio/mpeg")

	fmt.Printf("Recv: %v\n", m)

	texts := SplitText(m["text"])
	chatName := m["chat_name"]
	speed := m["speed"]

	wavs := make([][]byte, len(texts))

	var wg sync.WaitGroup
	for i, text := range texts {
		wg.Add(1)
		go func(idx int, t string) {
			defer wg.Done()
			fmt.Printf("\tProcess Part %v: Len:(%v) %v\n", idx, len(t), t)
			wavs[idx] = convert(t, chatName, speed)
			fmt.Printf("\tProcess Part %v: Done\n", idx)
		}(i, text)
	}
	wg.Wait()

	fmt.Printf("Merge Wav\n")
	rst, err := MergeWAVBytes(wavs)
	ep(err)

	fmt.Printf("Write Rsp\n")
	w.Write(rst)
}

func ep(err error) {
	if err != nil {
		panic(err)
	}
}
