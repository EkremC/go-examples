// https://www.calhoun.io/6-tips-for-using-strings-in-go/
package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func main() {
	// 1. multiline strings
	str := `This is a
multiline
string.`

	fmt.Println(str)

	str = `This string
	will have
	tabs in it`

	fmt.Println(str)

	// 2. Efficient concatenation
	var b bytes.Buffer

	for i := 0; i < 1000; i++ {
		b.WriteString(randString())
	}

	fmt.Println(b.String())

	// strings.Join
	var strs []string

	for i := 0; i < 1000; i++ {
		strs = append(strs, randString())
	}

	fmt.Println(strings.Join(strs, ""))

	// 3. Convert ints into strings
	i := 123
	t := strconv.Itoa(i)
	fmt.Println(t)

	// 4. Creating random strings
	fmt.Println(RandString(10))

	// 5. HasPrefix, HasSuffix
	fmt.Println(strings.HasPrefix("something", "some"))
	fmt.Println(strings.HasSuffix("something", "thing"))

	// 6.Strings can be converted into byte slices (and vice versa)
	var s string = "this is a string"
	fmt.Println(s)
	var by []byte
	by = []byte(s)
	fmt.Println(by)
	for i := range by {
		fmt.Println(string(by[i]))
	}
	s = string(by)
	fmt.Println(s)
}

var source = rand.NewSource(time.Now().UnixNano())

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func RandString(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[source.Int63()%int64(len(charset))]
	}
	return string(b)
}

func randString() string {
	return "abc-123-"
}
