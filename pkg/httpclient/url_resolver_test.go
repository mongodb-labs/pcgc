// Copyright 2019 MongoDB Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
