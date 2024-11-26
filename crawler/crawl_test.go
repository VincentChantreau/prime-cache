package crawler

import (
	"testing"
)

func TestAsAbsoluteUrl(t *testing.T) {
	crawler := Crawler{}
	path := "/path"
	origin := "https://test.com"
	url, err := crawler.asAbsoluteUrl(&path, &origin)
	if err != nil {
		t.Fail()
	}
	if url.String() != "https://test.com/path" {
		t.Fail()
	}
}
