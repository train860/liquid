// Package filters is an internal package that defines the standard Liquid filters.
package filters

import (
	"encoding/json"
	"fmt"
	"html"
	"math"
	"net/url"
	"reflect"
	"regexp"
	"strings"
	"time"
	"unicode"

	"github.com/osteele/liquid/values"
	"github.com/osteele/tuesday"
)

// A FilterDictionary holds filters.
type FilterDictionary interface {
	AddFilter(string, interface{})
}

// AddStandardFilters defines the standard Liquid filters.
func AddStandardFilters(fd FilterDictionary) { // nolint: gocyclo
	// value filters
	fd.AddFilter("default", func(bindings map[string]interface{}, value, defaultValue interface{}) interface{} {
		if value == nil || value == false || values.IsEmpty(value) {
			value = defaultValue
		}
		return value
	})
	fd.AddFilter("json", func(bindings map[string]interface{}, a interface{}) interface{} {
		result, _ := json.Marshal(a)
		return result
	})

	// array filters
	fd.AddFilter("compact", func(bindings map[string]interface{}, a []interface{}) (result []interface{}) {
		for _, item := range a {
			if item != nil {
				result = append(result, item)
			}
		}
		return
	})
	fd.AddFilter("concat", func(bindings map[string]interface{}, a, b []interface{}) (result []interface{}) {
		result = make([]interface{}, 0, len(a)+len(b))
		return append(append(result, a...), b...)
	})
	fd.AddFilter("join", joinFilter)
	fd.AddFilter("map", func(bindings map[string]interface{}, a []interface{}, key string) (result []interface{}) {
		keyValue := values.ValueOf(key)
		for _, obj := range a {
			value := values.ValueOf(obj)
			result = append(result, value.PropertyValue(keyValue).Interface())
		}
		return result
	})
	fd.AddFilter("reverse", reverseFilter)
	fd.AddFilter("sort", sortFilter)
	// https://shopify.github.io/liquid/ does not demonstrate first and last as filters,
	// but https://help.shopify.com/themes/liquid/filters/array-filters does
	fd.AddFilter("first", func(bindings map[string]interface{}, a []interface{}) interface{} {
		if len(a) == 0 {
			return nil
		}
		return a[0]
	})
	fd.AddFilter("last", func(bindings map[string]interface{}, a []interface{}) interface{} {
		if len(a) == 0 {
			return nil
		}
		return a[len(a)-1]
	})
	fd.AddFilter("uniq", uniqFilter)

	// date filters
	fd.AddFilter("date", func(bindings map[string]interface{}, t time.Time, format func(string) string) (string, error) {
		f := format("%a, %b %d, %y")
		return tuesday.Strftime(f, t)
	})

	// number filters
	fd.AddFilter("abs", func(bindings map[string]interface{}, a float64) float64 {
		return math.Abs(a)
	})
	fd.AddFilter("ceil", func(bindings map[string]interface{}, a float64) int {
		return int(math.Ceil(a))
	})
	fd.AddFilter("floor", func(bindings map[string]interface{}, a float64) int {
		return int(math.Floor(a))
	})
	fd.AddFilter("modulo", func(bindings map[string]interface{}, a, b float64) float64 {
		return math.Mod(a, b)
	})
	fd.AddFilter("minus", func(bindings map[string]interface{}, a, b float64) float64 {
		return a - b
	})
	fd.AddFilter("plus", func(bindings map[string]interface{}, a, b float64) float64 {
		return a + b
	})
	fd.AddFilter("times", func(bindings map[string]interface{}, a, b float64) float64 {
		return a * b
	})
	fd.AddFilter("divided_by", func(bindings map[string]interface{}, a float64, b interface{}) interface{} {
		switch q := b.(type) {
		case int, int16, int32, int64:
			return int(a) / q.(int)
		case float32, float64:
			return a / b.(float64)
		default:
			return nil
		}
	})
	fd.AddFilter("round", func(bindings map[string]interface{}, n float64, places func(int) int) float64 {
		pl := places(0)
		exp := math.Pow10(pl)
		return math.Floor(n*exp+0.5) / exp
	})

	// sequence filters
	fd.AddFilter("size", func(bindings map[string]interface{}, value interface{}) int {
		return values.Length(value)
	})

	// string filters
	fd.AddFilter("append", func(bindings map[string]interface{}, s, suffix string) string {
		return s + suffix
	})
	fd.AddFilter("capitalize", func(bindings map[string]interface{}, s, suffix string) string {
		if len(s) == 0 {
			return s
		}
		return strings.ToUpper(s[:1]) + s[1:]
	})
	fd.AddFilter("downcase", func(bindings map[string]interface{}, s, suffix string) string {
		return strings.ToLower(s)
	})
	fd.AddFilter("escape", func(bindings map[string]interface{}, s, suffix string) string {
		return html.EscapeString(s)
	})
	fd.AddFilter("escape_once", func(bindings map[string]interface{}, s, suffix string) string {
		return html.EscapeString(html.UnescapeString(s))
	})
	fd.AddFilter("newline_to_br", func(bindings map[string]interface{}, s string) string {
		return strings.Replace(s, "\n", "<br />", -1)
	})
	fd.AddFilter("prepend", func(bindings map[string]interface{}, s, prefix string) string {
		return prefix + s
	})
	fd.AddFilter("remove", func(bindings map[string]interface{}, s, old string) string {
		return strings.Replace(s, old, "", -1)
	})
	fd.AddFilter("remove_first", func(bindings map[string]interface{}, s, old string) string {
		return strings.Replace(s, old, "", 1)
	})
	fd.AddFilter("replace", func(bindings map[string]interface{}, s, old, new string) string {
		return strings.Replace(s, old, new, -1)
	})
	fd.AddFilter("replace_first", func(bindings map[string]interface{}, s, old, new string) string {
		return strings.Replace(s, old, new, 1)
	})
	fd.AddFilter("sort_natural", sortNaturalFilter)
	fd.AddFilter("slice", func(bindings map[string]interface{}, s string, start int, length func(int) int) string {
		ss := []rune(s)
		n := length(1)
		if start < 0 {
			start = len(ss) + start
		}
		end := start + n
		if end > len(ss) {
			end = len(ss)
		}
		return string(ss[start:end])
	})
	fd.AddFilter("split", splitFilter)
	fd.AddFilter("strip_html", func(bindings map[string]interface{}, s string) string {
		// TODO this probably isn't sufficient
		return regexp.MustCompile(`<.*?>`).ReplaceAllString(s, "")
	})
	fd.AddFilter("strip_newlines", func(bindings map[string]interface{}, s string) string {
		return strings.Replace(s, "\n", "", -1)
	})
	fd.AddFilter("strip", func(bindings map[string]interface{}, s string) string {
		return strings.TrimSpace(s)
	})
	fd.AddFilter("lstrip", func(bindings map[string]interface{}, s string) string {
		return strings.TrimLeftFunc(s, unicode.IsSpace)
	})
	fd.AddFilter("rstrip", func(bindings map[string]interface{}, s string) string {
		return strings.TrimRightFunc(s, unicode.IsSpace)
	})
	fd.AddFilter("truncate", func(bindings map[string]interface{}, s string, length func(int) int, ellipsis func(string) string) string {
		n := length(50)
		el := ellipsis("...")
		// runes aren't bytes; don't use slice
		re := regexp.MustCompile(fmt.Sprintf(`^(.{%d})..{%d,}`, n-len(el), len(el)))
		return re.ReplaceAllString(s, `$1`+el)
	})
	fd.AddFilter("truncatewords", func(bindings map[string]interface{}, s string, length func(int) int, ellipsis func(string) string) string {
		el := ellipsis("...")
		n := length(15)
		re := regexp.MustCompile(fmt.Sprintf(`^(?:\s*\S+){%d}`, n))
		m := re.FindString(s)
		if m == "" {
			return s
		}
		return m + el
	})
	fd.AddFilter("upcase", func(bindings map[string]interface{}, s, suffix string) string {
		return strings.ToUpper(s)
	})
	fd.AddFilter("url_encode", func(bindings map[string]interface{}, s string) string {
		return url.QueryEscape(s)
	})
	fd.AddFilter("url_decode", func(bindings map[string]interface{}, s string) (string, error) {
		return url.QueryUnescape(s)
	})

	// debugging filters
	// inspect is from Jekyll
	fd.AddFilter("inspect", func(bindings map[string]interface{}, value interface{}) string {
		s, err := json.Marshal(value)
		if err != nil {
			return fmt.Sprintf("%#v", value)
		}
		return string(s)
	})
	fd.AddFilter("type", func(value interface{}) string {
		return fmt.Sprintf("%T", value)
	})
}

func joinFilter(bindings map[string]interface{}, a []interface{}, sep func(string) string) interface{} {
	ss := make([]string, 0, len(a))
	s := sep(" ")
	for _, v := range a {
		if v != nil {
			ss = append(ss, fmt.Sprint(v))
		}
	}
	return strings.Join(ss, s)
}

func reverseFilter(bindings map[string]interface{}, a []interface{}) interface{} {
	result := make([]interface{}, len(a))
	for i, x := range a {
		result[len(result)-1-i] = x
	}
	return result
}

var wsre = regexp.MustCompile(`[[:space:]]+`)

func splitFilter(bindings map[string]interface{}, s, sep string) interface{} {
	result := strings.Split(s, sep)
	if sep == " " {
		// Special case for Ruby, therefore Liquid
		result = wsre.Split(s, -1)
	}
	// This matches Ruby / Liquid / Jekyll's observed behavior.
	for len(result) > 0 && result[len(result)-1] == "" {
		result = result[:len(result)-1]
	}
	return result
}

func uniqFilter(bindings map[string]interface{}, a []interface{}) (result []interface{}) {
	seenMap := map[interface{}]bool{}
	seen := func(item interface{}) bool {
		if k := reflect.TypeOf(item).Kind(); k < reflect.Array || k == reflect.Ptr || k == reflect.UnsafePointer {
			if seenMap[item] {
				return true
			}
			seenMap[item] = true
			return false
		}
		// the O(n^2) case:
		for _, other := range result {
			if eqItems(item, other) {
				return true
			}
		}
		return false
	}
	for _, item := range a {
		if !seen(item) {
			result = append(result, item)
		}
	}
	return
}

func eqItems(a, b interface{}) bool {
	if reflect.TypeOf(a).Comparable() && reflect.TypeOf(b).Comparable() {
		return a == b
	}
	return reflect.DeepEqual(a, b)
}
