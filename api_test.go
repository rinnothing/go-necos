package go_necos

import (
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestGet(t *testing.T) {
	type testCases struct {
		name   string
		path   string
		query  url.Values
		result string
	}

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

	tableTests := []testCases{
		{
			name:   "empty_root",
			path:   "",
			query:  url.Values{},
			result: "",
		},
		{
			name:   "empty_directory",
			path:   "/directory",
			query:  url.Values{},
			result: "/directory",
		},
		{
			name:   "query_root",
			path:   "",
			query:  url.Values{"oh": {"hello", "there"}},
			result: "?oh=hello&oh=there",
		},
		{
			name:   "query_directory",
			path:   "/directory",
			query:  url.Values{"oh": {"hello", "there"}},
			result: "/directory?oh=hello&oh=there",
		},
	}

	for _, cs := range tableTests {
		t.Run(cs.name, func(t *testing.T) {
			t.Parallel()
			var answer string

			err := c.Get(cs.path, cs.query, &answer)
			require.NoError(t, err)

			require.Equal(t, cs.result, answer)
		})
	}
}
