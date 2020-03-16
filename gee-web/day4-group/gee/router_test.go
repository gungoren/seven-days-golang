package gee

import (
	"fmt"
	"reflect"
	"testing"
)

func newTestRouter() *router {
	r := newRouter()
	r.addRoute("GET", "/", nil)
	r.addRoute("GET", "/hello/:name", nil)
	r.addRoute("GET", "/hello/b/c", nil)
	r.addRoute("GET", "/hi/:name", nil)
	r.addRoute("GET", "/assets/*filepath", nil)
	return r
}

func TestParsePattern(t *testing.T) {
	ok := reflect.DeepEqual(parsePattern("/p/:name"), []string{"p", ":name"})
	ok = ok && reflect.DeepEqual(parsePattern("/p/*"), []string{"p", "*"})
	ok = ok && reflect.DeepEqual(parsePattern("/p/*name/*"), []string{"p", "*name"})
	if !ok {
		t.Fatal("test parsePattern failed")
	}
}

func TestGetRoute(t *testing.T) {
	r := newTestRouter()
	n, ps := r.getRoute("GET", "/hello/mehmet")

	if n == nil {
		t.Fatal("nil shouldn't be returned")
	}

	if n.pattern != "/hello/:name" {
		t.Fatal("should match /hello/:name")
	}

	if ps["name"] != "mehmet" {
		t.Fatal("name should be equal to 'mehmet'")
	}

	fmt.Printf("matched path: %s, params['name']: %s\n", n.pattern, ps["name"])
}

func TestGetRoute2(t *testing.T) {
	r := newTestRouter()
	n, ps := r.getRoute("GET", "/assets/file1.txt")
	ok := n.pattern == "/assets/*filepath" && ps["filepath"] == "file1.txt"
	if !ok {
		t.Fatal("pattern should be /assets/*filepath & filepath should be file1.txt")
	}

	n1, ps1 := r.getRoute("GET", "assets/css/test.css")
	ok = n1.pattern == "/assets/*filepath" && ps1["filepath"] == "css/test.css"
	if !ok {
		t.Fatal("pattern should be /assets/*filepath & filepath should be css/test.css")
	}
}

func TestGetRoutes(t *testing.T) {
	r := newTestRouter()
	nodes := r.getRoutes("GET")
	for i, n := range nodes {
		fmt.Println(i+1, n)
	}

	if len(nodes) != 5 {
		t.Fatal("the number of routes should be 4")
	}
}
