package go_necos

import (
	"encoding/json"
	"github.com/stretchr/testify/require"
	"maps"
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

	query := url.Values{"oh": {"hello", "there"}, "obi": {"wan"}}
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
			query:  query,
			result: "/?" + query.Encode(),
		},
		{
			name:   "query_directory",
			path:   "/directory",
			query:  query,
			result: "/directory?" + query.Encode(),
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

func TestPost(t *testing.T) {
	type testCases struct {
		name   string
		path   string
		query  url.Values
		result string
	}

	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
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

	query := url.Values{"oh": {"hello", "there"}, "obi": {"wan"}}
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
			query:  query,
			result: "/?" + query.Encode(),
		},
		{
			name:   "query_directory",
			path:   "/directory",
			query:  query,
			result: "/directory?" + query.Encode(),
		},
	}

	wg := sync.WaitGroup{}
	for _, cs := range tableTests {
		wg.Add(1)

		t.Run(cs.name, func(t *testing.T) {
			t.Parallel()

			defer wg.Done()
			var answer string

			err := c.Post(cs.path, cs.query, &answer)

			require.NoError(t, err)

			require.Equal(t, cs.result, answer)
		})
	}

	go func() {
		wg.Wait()
		s.Close()
	}()
}

func TestQuery(t *testing.T) {
	type testCases struct {
		name         string
		path         string
		query        url.Values
		defaultQuery url.Values
		result       string
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

	query := url.Values{"oh": {"hello", "there"}, "obi": {"wan"}}
	defaultQuery := url.Values{"oh": {"my", "god"}, "hii": {}}

	biQuery := maps.Clone(query)
	for k, v := range defaultQuery {
		if _, ok := biQuery[k]; ok {
			continue
		}

		biQuery[k] = v
	}

	tableTests := []testCases{
		{
			name:         "empty_both",
			path:         "/",
			defaultQuery: url.Values{},
			query:        url.Values{},
			result:       "/",
		},
		{
			name:         "empty_default",
			path:         "/",
			defaultQuery: url.Values{},
			query:        query,
			result:       "/?" + query.Encode(),
		},
		{
			name:         "empty_query",
			path:         "/",
			defaultQuery: defaultQuery,
			query:        url.Values{},
			result:       "/?" + defaultQuery.Encode(),
		},
		{
			name:         "clash_query",
			path:         "/",
			defaultQuery: defaultQuery,
			query:        query,
			result:       "/?" + biQuery.Encode(),
		},
	}

	wg := sync.WaitGroup{}
	for _, cs := range tableTests {
		c := Client{
			DefaultQuery: cs.defaultQuery,
			Domain:       s.URL,
		}

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
