package go_necos

import (
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

type testCase struct {
	name   string
	path   string
	query  url.Values
	result string
}

func makeCase(name string, path string, query url.Values) testCase {
	return testCase{
		name:   name,
		path:   path,
		query:  query,
		result: path + query.Encode(),
	}
}

func TestGet(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			w.WriteHeader(http.StatusMethodNotAllowed)
		}

		ans := r.URL.Path + r.URL.RawQuery
		_, _ = w.Write([]byte(ans))
	}))
	defer s.Close()

	c := Client{
		Domain: s.URL,
	}

	cases := []testCase{
		makeCase("empty_root", "", url.Values{}),
		makeCase("empty_directory", "/directory", url.Values{}),
		makeCase("query_root", "", url.Values{"oh": {"hello", "there"}}),
		makeCase("query_directory", "/directory", url.Values{"oh": {"hello", "there"}}),
	}

	for _, cs := range cases {
		var answer string

		err := c.Get(cs.path, cs.query, &answer)
		require.NoError(t, err)

		require.Equal(t, cs.result, answer)
	}
}
