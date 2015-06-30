package node

import (
	"bytes"
	"math/rand"
	"time"
)

func timestamp() string {
	now := time.Now().UTC()
	// For whatever reason the API seems picky about +00:00 vs Z for the TZD in the timestamp
	// this means the constant for RFC3339 doesn't work but the below string does
	//layout := time.RFC3339
	layout := "2006-01-02T15:04:05-07:00"

	return now.Format(layout)
}

func string_to_reader(body string) *bytes.Buffer {
	return bytes.NewBuffer([]byte(body))
}

func sleep_splay(ms int) {
	max_splay := (ms / 5)
	sleep_time := ms + rand.Intn(max_splay)

	time.Sleep(time.Duration(sleep_time) * time.Millisecond)
}
