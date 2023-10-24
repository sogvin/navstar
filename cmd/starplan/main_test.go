package main

import (
	"os"
	"testing"
	"time"
)

func Test_main(t *testing.T) {
	os.Args = []string{"starplan", "--bind", ":9183"}
	go main()
	<-time.After(5 * time.Millisecond)
}
