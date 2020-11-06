package fb

import (
	"log"
	"testing"
)

func TestFbLoop(t *testing.T) {
	len := 10
	i := FbLoop(len)
	log.Printf("fb %d is %d", len, i)
}

func TestFbrecursion(t *testing.T) {
	len := 10
	i := Fbrecursion(len)
	log.Printf("fb %d is %d", len, i)
}
