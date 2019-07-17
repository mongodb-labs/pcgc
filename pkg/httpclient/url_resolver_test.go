package httpclient

import "testing"

const localhost = "http://localhost"

func TestResolvePathWithoutExpansions(t *testing.T) {
	resolver := NewURLResolver(localhost)
	got := resolver.Of("/groups")

	want := localhost + "/groups"
	if got != want {
		t.Errorf("Of(\"/groups\") = %s; want %s", got, want)
	}
}

func TestResolvePathWithExpansions(t *testing.T) {
	resolver := NewURLResolver(localhost)
	got := resolver.Of("/groups/%s", "1")

	want := localhost + "/groups/1"
	if got != want {
		t.Errorf("Of(\"/groups/1\") = %s; want %s", got, want)
	}
}

func TestResolvePathWithPrefixWithoutExpansions(t *testing.T) {
	resolver := NewURLResolverWithPrefix(localhost, "prefix")
	got := resolver.Of("/groups")

	want := localhost + "/prefix/groups"
	if got != want {
		t.Errorf("Of(\"/groups\") = %s; want %s", got, want)
	}
}

func TestResolvePathWithPrefixWithExpansions(t *testing.T) {
	resolver := NewURLResolverWithPrefix(localhost, "prefix")
	got := resolver.Of("/groups/%s", "1")

	want := localhost + "/prefix/groups/1"
	if got != want {
		t.Errorf("Of(\"/groups/1\") = %s; want %s", got, want)
	}
}

func TestRelativePathsAreCleaned(t *testing.T) {
	resolver := NewURLResolver(localhost)
	got := resolver.Of("/../../groups")

	want := localhost + "/groups"
	if got != want {
		t.Errorf("Of(\"/groups\") = %s; want %s", got, want)
	}
}
