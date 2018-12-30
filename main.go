package main

import (
	"fmt"
	"net/http"
	"os"
	"poetry"
)

/*
Anthonys-MacBook-Pro:go mcclayac$ godoc fmt Fprintf | more
func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error)
Fprintf formats according to a format specifier and writes to w. It
returns the number of bytes written and any write error encountered.
*/

func poemHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	poemName := r.Form["name"][0]
	//fileName := "doggie.txt"
	p, err := poetry.LoadPoem(poemName)

	if err != nil {
		http.Error(w, "File Not Found", http.StatusInternalServerError)
		fmt.Printf("An Error occured reading file %s \n", poemName)
		//os.Exit(-1)
		return
	}

	_, err = fmt.Fprintf(w, "Poem Name: %s \n\n%s\n\n", poemName, p)

	if err != nil {
		fmt.Printf("An Error occured reading file %s \n", poemName)
		os.Exit(-1)
	}

}

func main() {

	/*fileName := "doggie.txt"
	p, err := poetry.LoadPoem(fileName)

	if err != nil {
		fmt.Printf("An Error occured reading file %s \n", fileName)
		os.Exit(-1)
	}

	fmt.Printf("%s\n", p)
	*/

	http.HandleFunc("/poem", poemHandler)
	http.ListenAndServe(":8088", nil)

}

/*
type Values map[string][]string
Values maps a string key to a list of values. It is typically used for
query parameters and form values. Unlike in the http.Header map, the
keys in a Values map are case-sensitive.

func ParseQuery(query string) (Values, error)
ParseQuery parses the URL-encoded query string and returns a map listing
the values specified for each key. ParseQuery always returns a non-nil
map containing all the valid query parameters found; err describes the
first decoding error encountered, if any.

Query is expected to be a list of key=value settings separated by
ampersands or semicolons. A setting without an equals sign is
interpreted as a key set to an empty value.

func (v Values) Add(key, value string)
Add adds the value to key. It appends to any existing values associated
with key.

func (v Values) Del(key string)
Del deletes the values associated with key.

func (v Values) Encode() string
Encode encodes the values into ``URL encoded'' form ("bar=baz&foo=quux")
sorted by key.

func (v Values) Get(key string) string
Get gets the first value associated with the given key. If there are no
values associated with the key, Get returns the empty string. To access
multiple values, use the map directly.

func (v Values) Set(key, value string)
Set sets the key to value. It replaces any existing values.

*/
