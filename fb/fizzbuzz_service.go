package fb

import (
	"errors"
	"strconv"
)

func fizzbuzzList(fbRequest FizzbuzzRequest) ([]string, error) {
	addRequest(fbRequest)
	if fbRequest.Limit < 1 || fbRequest.Int1*fbRequest.Int2 == 0 {
		return nil, errors.New("limit neds to be greater than 1 and Int1&2 different from 0")
	}
	response := make([]string, fbRequest.Limit)
	for i := 0; i < fbRequest.Limit; i++ {
		response[i] = fizzbuzzInt(i+1, fbRequest.Int1, fbRequest.Int2, fbRequest.Str1, fbRequest.Str2)
	}
	return response, nil
}

func fizzbuzzInt(i, int1, int2 int, str1, str2 string) string {
	switch {
	case i%int1 == 0 && i%int2 == 0:
		return str1 + str2
	case i%int1 == 0:
		return str1
	case i%int2 == 0:
		return str2
	default:
		return strconv.Itoa(i)
	}
}
