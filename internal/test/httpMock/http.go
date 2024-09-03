package httpMock

import (
	"errors"
	"github.com/Sabyradinov/golang-hex-layout/common"
)

type HttpClientMock struct {
	TestError bool
}

func (h HttpClientMock) RequestJSON(request common.HttpArgs) (httpStatus int, responseBody []byte, err error) {
	if h.TestError {
		return 500, nil, errors.New("some error")
	}
	return 200, nil, nil
}
