package fb

import (
	"fmt"
	"strings"
	"testing"
)

func TestFizzbuzz(t *testing.T) {
	request := FizzbuzzRequest{Int1: 2, Int2: 3, Limit: 6, Str1: "two", Str2: "three"}
	list, err := fizzbuzzList(request)
	t.Logf("for parameters 2,3,6,two,three , this list was obtained:%v", list)
	if err != nil || strings.Join(list, ",") != "1,two,three,two,5,twothree" {
		t.Fatalf("unexpected error")
	}
}

func TestFizzbuzzFail(t *testing.T) {
	request1 := FizzbuzzRequest{Int1: 2, Int2: 3, Limit: -3, Str1: "two", Str2: "three"}
	_, err1 := fizzbuzzList(request1)
	request2 := FizzbuzzRequest{Int1: 2, Int2: 0, Limit: 5, Str1: "two", Str2: "three"}
	_, err2 := fizzbuzzList(request2)
	if err1 == nil || err2 == nil {
		t.Fatalf("expected error, found none")
	}
}

func TestMostFrequent(t *testing.T) {
	request, err := mostFrequentRequest()
	if err != nil {
		t.Fatalf(fmt.Sprint(err))
	}
	t.Logf("Most frequent request and count:%v", request)
}
