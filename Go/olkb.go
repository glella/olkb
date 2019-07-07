package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

func get_data(s string) string {
	resp, err := http.Get(s)
	if err != nil {
		fmt.Println("Could not reach the page.")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Could not parse page contents.")
	}
	return string(body)
}

func main() {

	data := get_data("https://orders.olkb.com")
	re := regexp.MustCompile(`<li>10000\d{4}`)
	order_number := "100007000" // put your own order number

	temp := re.FindAllString(data, -1)
	var orders []string
	for _, each := range temp {
		trim := strings.TrimLeft(each, "<li>")
		orders = append(orders, trim)
	}
	order_count := len(orders)
	order_position := 0
	if Include(orders, order_number) {
		order_position = Index(orders, order_number) + 1
	}

	fmt.Printf("olkb position: %d\n", order_position)
	fmt.Println("---")
	fmt.Printf("Total orders: %d\n", order_count)
	fmt.Printf("Order #: %s\n", order_number)

}

// Slice of Strings support functions

// Index returns the first index of the target string `t`, or
// -1 if no match is found.
func Index(vs []string, t string) int {
	for i, v := range vs {
		if v == t {
			return i
		}
	}
	return -1
}

// Include returns `true` if the target string t is in the
// slice.
func Include(vs []string, t string) bool {
	return Index(vs, t) >= 0
}

// Any returns `true` if one of the strings in the slice
// satisfies the predicate `f`.
func Any(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if f(v) {
			return true
		}
	}
	return false
}

// All returns `true` if all of the strings in the slice
// satisfy the predicate `f`.
func All(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if !f(v) {
			return false
		}
	}
	return true
}

// Filter returns a new slice containing all strings in the
// slice that satisfy the predicate `f`.
func Filter(vs []string, f func(string) bool) []string {
	vsf := make([]string, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

// Map returns a new slice containing the results of applying
// the function `f` to each string in the original slice.
func Map(vs []string, f func(string) string) []string {
	vsm := make([]string, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}
