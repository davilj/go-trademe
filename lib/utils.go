package trademe

import (
	"testing"
)

func HandleTestError(t *testing.T, errMsg string, err error) {
	if err != nil {
		t.Errorf(errMsg, err)
	}
}

func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}
