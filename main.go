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
	if err != nil {
		fmt.Printf("Recv Error: %v\n", err)
		return
	}

	m := map[string]string{}
	err = json.Unmarshal(bytes, &m)
	if err != nil {
		fmt.Printf("Recv Error: %v\n", err)
		return
	}

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
	if err != nil {
		fmt.Printf("Merge Wav Error: %v\n", err)
		return
	}

	fmt.Printf("Write Rsp\n")
	_, err = w.Write(rst)
	if err != nil {
		fmt.Printf("Write Error: %v\n", err)
		return
	}
}
