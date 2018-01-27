package body

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestESQueryModifier(t *testing.T) {
	cfg := `{"keys":["foo","bar","x"]}`
	modifier, err := modifierFromJSON([]byte(cfg))
	if err != nil {
		t.Error(err)
		return
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Error("wrong method:", r.Method)
		}
		query := r.URL.Query()
		for _, k := range []string{"foo", "bar", "x"} {
			if query.Get(k) != "" {
				t.Error("the param", k, "is present in the querystring")
			}
		}
		for _, k := range []string{"y", "z"} {
			if query.Get(k) == "" {
				t.Error("the param", k, "is not present in the querystring")
			}
		}
	}))

	URL, err := url.Parse(ts.URL)
	if err != nil {
		t.Error(err)
		return
	}
	query := URL.Query()
	query.Add("foo", "1")
	query.Add("bar", "1")
	query.Add("x", "1")
	query.Add("y", "1")
	query.Add("z", "1")
	URL.RawQuery = query.Encode()

	req, err := http.NewRequest("GET", URL.String(), ioutil.NopCloser(bytes.NewReader([]byte{})))
	if err != nil {
		t.Error(err)
		return
	}

	if err := modifier.RequestModifier().ModifyRequest(req); err != nil {
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
	if _, err := modifierFromJSON([]byte(`"x"]}`)); err == nil {
		t.Errorf("error expected")
	}
}
