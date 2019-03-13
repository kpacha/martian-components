package modifier

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestESQueryModifier(t *testing.T) {
	cfg := `{"keys_to_extract":["foo","x"],"template":"{\"foo\":\"{{index .foo 0}}\",\"bar\":\"{{index .bar 0}}\",\"x\":\"{{index .x 0 }}\"}"}`
	modifier, err := FromJSON([]byte(cfg))
	if err != nil {
		t.Error(err)
		return
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Error("wrong method:", r.Method)
		}
		query := r.URL.Query()
		for _, k := range []string{"foo", "x"} {
			if query.Get(k) != "" {
				t.Error("the param", k, "is present in the querystring")
			}
		}
		for _, k := range []string{"bar", "y", "z"} {
			if query.Get(k) == "" {
				t.Error("the param", k, "is not present in the querystring")
			}
		}

		if ct := r.Header.Get("Content-Type"); ct != "application/json; charset=UTF-8" {
			t.Errorf("unexpected content-tupe: %s", ct)
		}

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			t.Error(err)
			return
		}
		r.Body.Close()

		if string(body) != `{"foo":"1","bar":"foobar","x":"booo"}` {
			t.Errorf("unexpected response: %s", string(body))
		}
	}))

	URL, err := url.Parse(ts.URL)
	if err != nil {
		t.Error(err)
		return
	}
	query := URL.Query()
	query.Add("foo", "1")
	query.Add("bar", "foobar")
	query.Add("x", "booo")
	query.Add("y", "1")
	query.Add("z", "1")
	URL.RawQuery = query.Encode()

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

func TestESQueryModifier_badDSL(t *testing.T) {
	if _, err := FromJSON([]byte(`"x"]}`)); err == nil {
		t.Errorf("error expected")
	}
}
