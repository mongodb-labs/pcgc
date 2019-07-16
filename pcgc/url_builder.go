package pcgc

import (
	"fmt"
	"net/url"
	"path"
)

type baseURL struct {
	base *url.URL
}

// URLResolver contract for resolving any paths against the given base URL
type URLResolver interface {
	Of(path string, v ...interface{}) string
}

// NewURLResolver builds a new API URL which can be used to build any path
func NewURLResolver(base string) URLResolver {
	return parseBaseURL(base)
}

// NewURLResolverWithPrefix builds a new API URL using a prefix for all other paths
func NewURLResolverWithPrefix(base string, prefix string) URLResolver {
	var err error

	result := parseBaseURL(base)

	// augment the URL with a prefix
	result.base, err = url.Parse(prefix)
	panicOnUnrecoverableError(err)

	return result
}

func parseBaseURL(base string) baseURL {
	result := baseURL{}
	var err error
	result.base, err = url.Parse(base)
	panicOnUnrecoverableError(err)
	return result
}

// Of builds a URL by concatenating the base URL with the specified path, replacing all needed parts
func (u baseURL) Of(apiPath string, parts ...interface{}) string {
	result, err := u.base.Parse(fmt.Sprintf(path.Clean(apiPath), parts))
	panicOnUnrecoverableError(err)
	return result.String()
}
