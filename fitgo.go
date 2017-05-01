/*
Package fitgo is purely to help  parsing text in online commands:

	import . "github.com/kavehmz/fitgo"
	...

	Lines("https://www.cloudflare.com/ips-v4").Grep(`131\.`).Echo().Count()
	Lines("/tmp/my_text").Grep(`Word`).Count().Grep(`2`).Echo()
*/
package fitgo

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strconv"
)

// Stream is a structure that will be passed between calls.
type Stream struct {
	lines chan string
}

// Stdin is the shortcode for reading lines from stdin
func Stdin() *Stream {
	return Lines("")
}

// S is the shortcode for reading lines from stdin
func S() *Stream {
	return Stdin()
}

// Lines will fetch and parse a text file or a URL.
func Lines(s string) *Stream {
	tmp := initStream()
	if r, e := url.Parse(s); e == nil && r.Scheme != "" {
		go tmp.readURL(r.String())
	} else if s != "" {
		go tmp.readFile(s)
	} else {
		go tmp.readStdIn(s)
	}
	return tmp
}

// L is the shortcode for reading lines from stdin
func L(s string) *Stream {
	return Lines(s)
}

// Grep will filter lines based on passed regular expression
func (u *Stream) Grep(s string) *Stream {
	tmp := initStream()
	go func() {
		r, _ := regexp.Compile(s)
		for l := range u.lines {
			if r.Match([]byte(l)) {
				tmp.lines <- l
			}
		}
		close(tmp.lines)
	}()
	return tmp
}

// G will filter lines based on passed regular expression
func (u *Stream) G(s string) *Stream {
	return u.Grep(s)
}

// Replace will Replace lines based on passed regular expression and replace string
// It is using regexp.ReplaceAllString
func (u *Stream) Replace(s string, t string) *Stream {
	tmp := initStream()
	go func() {
		r, _ := regexp.Compile(s)
		for l := range u.lines {
			tmp.lines <- r.ReplaceAllString(l, t)
		}
		close(tmp.lines)
	}()
	return tmp
}

// R will Replace lines based on passed regular expression and replace string
// It is using regexp.ReplaceAllString
func (u *Stream) R(s string, t string) *Stream {
	return u.Replace(s, t)
}

// Count will return count of lines in the current stream
func (u *Stream) Count() *Stream {
	tmp := initStream()
	go func() {
		count := 0
		for range u.lines {
			count++
		}
		tmp.lines <- strconv.Itoa(count)
		close(tmp.lines)
	}()
	return tmp
}

// C will return count of lines in the current stream
func (u *Stream) C() *Stream {
	return u.Count()
}

// Echo will print the current stream lines
func (u *Stream) Echo() *Stream {
	for l := range u.lines {
		fmt.Println(l)
	}
	return u
}

// E will print the current stream lines
func (u *Stream) E() *Stream {
	return u.Echo()
}

func initStream() *Stream {
	return &Stream{make(chan string)}
}

func (u *Stream) readURL(s string) *Stream {
	resp, err := http.Get(s)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		u.lines <- scanner.Text()
	}
	close(u.lines)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return u
}

func (u *Stream) readFile(s string) *Stream {
	file, err := os.Open(s)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		u.lines <- scanner.Text()
	}
	close(u.lines)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return u
}

func (u *Stream) readStdIn(s string) *Stream {
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		u.lines <- scanner.Text()
	}
	close(u.lines)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return u
}
