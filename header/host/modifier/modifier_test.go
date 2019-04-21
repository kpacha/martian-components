package modifier

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestQueryModifier(t *testing.T) {
	cfg := `{"host":"some-host.tld"}`
	modifier, err := FromJSON([]byte(cfg))
	if err != nil {
		t.Error(err)
		return
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Host != "some-host.tld" {
			t.Error("wrong host header:", r.Host)
		}
	}))

	URL, err := url.Parse(ts.URL)
	if err != nil {
		t.Error(err)
		return
	}

	req, err := http.NewRequest("GET", URL.String(), ioutil.NopCloser(bytes.NewReader([]byte{})))
	if err != nil {
		t.Error(err)
		return
	}

	if err := modifier.ModifyRequest(req); err != nil {
		t.Error(err)
		return
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Error(err)
		return
	}
	if resp.StatusCode != 200 {
		t.Error("unexpected status code:", resp.StatusCode)
		return
	}
}
