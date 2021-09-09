package fb

import (
	"fmt"
	"log"
	"strings"
	"testing"
)

func TestFizzbuzz(t *testing.T) {
	request := FizzbuzzRequest{Int1: 2, Int2: 3, Limit: 6, Str1: "two", Str2: "three"}
	list, err := fizzbuzzList(request)
	fmt.Println(list)
	if err != nil || strings.Join(list, ",") != "1,two,three,two,5,twothree" {
		t.Fatalf("unexpected error")
	}
}

func TestMostFrequent(t *testing.T) {
	request, err := mostFrequentRequest()
	if err != nil {
		t.Fatalf(fmt.Sprint(err))
	}
	fmt.Println("habab")
	log.Printf("Most frequent request:%v", request)
}
