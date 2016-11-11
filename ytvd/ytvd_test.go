package ytvd

import (
	"testing"
	"log"
	"strings"
	"fmt"
	"os"
)
func check(e error) {
    if e != nil {
        panic(e)
    }
}

func TestDownloadVideo(t *testing.T) {
	tN := "Xfs49Yc2qOw"
	if fn, err := DownloadVideo(tN); fn != nil {
		log.Printf("%+v", fn)
		if !strings.Contains(fn.Url, tN) {
			t.Error("In url i not found ")
		}
		f, err := os.Open(fn.ContentFileName)
		check(err)
		b1 := make([]byte, 256)
		n1, err := f.Read(b1)
		
		t.Log(fmt.Sprintf("%d bytes: %s\n", n1, string(b1)))
		
	} else {
		log.Panic(err)
		t.Error(err)
	}
}

