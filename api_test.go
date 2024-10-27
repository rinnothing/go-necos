package go_necos

import (
	"encoding/json"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"net/url"
	"sync"
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

		ans := r.URL.Path
		if r.URL.RawQuery != "" {
			ans += "?" + r.URL.RawQuery
		}
		m, _ := json.Marshal(ans)
		_, _ = w.Write(m)
	}))

	c := Client{
		Domain: s.URL,
	}

	tableTests := []testCases{
		{
			name:   "empty_root",
			path:   "/",
			query:  url.Values{},
			result: "/",
		},
		{
			name:   "empty_directory",
			path:   "/directory",
			query:  url.Values{},
			result: "/directory",
		},
		{
			name:   "query_root",
			path:   "/",
			query:  url.Values{"oh": {"hello", "there"}},
			result: "/?oh=hello&oh=there",
		},
		{
			name:   "query_directory",
			path:   "/directory",
			query:  url.Values{"oh": {"hello", "there"}},
			result: "/directory?oh=hello&oh=there",
		},
	}

	wg := sync.WaitGroup{}
	for _, cs := range tableTests {
		wg.Add(1)

		t.Run(cs.name, func(t *testing.T) {
			t.Parallel()

			defer wg.Done()
			var answer string

			err := c.Get(cs.path, cs.query, &answer)

			require.NoError(t, err)

			require.Equal(t, cs.result, answer)
		})
	}

	go func() {
		wg.Wait()
		s.Close()
	}()
}
