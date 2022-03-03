package main

import (
	"errors"
	"hash/crc32"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"net/http/fcgi"
	"path"
	"sort"
	"strconv"
	"unicode/utf8"

	"github.com/coding-robots/goiwl/tokenizer"
)

const (
	// Minimum text length in bytes for analysis.
	// If text has fewer bytes, the error is shown.
	MinTextLen = 200
	// Maximum text length in bytes for analysis.
	// If text has more bytes, it's cut, but on letter boundary,
	// not on byte boundary.
	MaxTextLen = 30000
)

var (
	templates *template.Template
	// Sorted authors names
	sortedAuthors []string
	// Maps CRC32 of author name to author name.
	authorByCrc map[string]string
	// Maps author name to CRC32
	crcByAuthor map[string]string
	// Maps author URL-friendly name to author name.
	authorByURLName map[string]string
	// Maps the other way around
	urlNameByAuthor map[string]string
	// Maps author name to author information.
	infoByAuthor map[string]*AuthorInfo
)

type page struct {
	Title       string
	Description string
	Error       string
}

type indexPage struct {
	page
	Text string
}

type resultPage struct {
	page
	Author        string
	Crc           string
	AuthorURLSafe string
	Info          *AuthorInfo
	IsShared      bool
}

type writerPage struct {
	page
	Author string
	Info   *AuthorInfo
}

type aboutPage struct {
	page
	Authors         []string
	URLNameByAuthor map[string]string
}

type analyzingPage struct {
	page
	Delay int
	URL   string
}

// cutText returns a string cut to approximately len bytes on
// the nearest rune to len. The resulting text is never larger
// than len bytes, but may be smaller, including an empty string.
func cutText(s string, len int) string {
	for len > 0 {
		if utf8.RuneStart(s[len]) {
			return s[:len]
		}
		len--
	}
	return ""
}

func analyze(text string) (author string, err error) {
	if len(text) > MaxTextLen {
		text = cutText(text, MaxTextLen)
	}
	if len(text) < MinTextLen {
		err = errors.New("Text is too short.")
		return
	}
	author, _ = classifier.Classify(tokenizer.WordsSpecials(text))
	return
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	p := &indexPage{page: page{
		Description: "Check which famous writer you write like with this statistical analysis tool.",
	}}

	if r.Method == "POST" {
		author, err := analyze(r.FormValue("text"))
		if err != nil {
			p.Error = err.Error()
		} else {
			// You may wonder what's going on here, why the
			// user has to be shown this "analyzing" page with
			// loading indicator and wait 2 seconds until redirected,
			// even though we already know the results.
			// The thing is... people just don't believe that
			// the server does any actual text analysis if we
			// returned answer instantly. Seriously! So, we
			// need this kind of placebo delay.

			// Successful analysis, redirect to author page.
			p := new(analyzingPage)
			p.URL = path.Join("/", "b", crcByAuthor[author])
			p.Delay = rand.Intn(2) + 1

			if err := templates.Lookup("analyzing.html").Execute(w, p); err != nil {
				log.Printf("error: %s", err)
			}
			return
		}
	}

	if err := templates.Lookup("index.html").Execute(w, p); err != nil {
		log.Printf("error: %s", err)
	}
}

func makeResultHandler(isShared bool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		crc := path.Base(r.URL.Path)
		author, ok := authorByCrc[crc]
		if !ok {
			// Author not found
			http.NotFound(w, r)
			return
		}

		p := new(resultPage)
		p.Title = "You write like " + author
		p.Author = author
		p.Crc = crcByAuthor[author]
		p.Info = infoByAuthor[author]
		p.AuthorURLSafe = authorNameToURLSafeName(author)
		p.IsShared = isShared

		if err := templates.Lookup("result.html").Execute(w, p); err != nil {
			log.Printf("error: %s", err)
		}
	}
}

func writerRedirectHandler(w http.ResponseWriter, r *http.Request) {
	crc := path.Base(r.URL.Path)
	author, ok := authorByCrc[crc]
	if !ok {
		// Author not found
		http.NotFound(w, r)
		return
	}
	url := getAuthorURL(siteBaseURL, author)
	http.Redirect(w, r, url, http.StatusMovedPermanently)
}

func writerHandler(w http.ResponseWriter, r *http.Request) {
	urlName := path.Base(r.URL.Path)
	author, ok := authorByURLName[urlName]
	if !ok {
		// Author not found
		http.NotFound(w, r)
		return
	}

	p := new(writerPage)
	p.Title = author
	p.Author = author
	p.Info = infoByAuthor[author]

	if err := templates.Lookup("writer.html").Execute(w, p); err != nil {
		log.Printf("error: %s", err)
	}
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	p := new(aboutPage)
	p.Title = "About"
	p.Authors = sortedAuthors
	p.URLNameByAuthor = urlNameByAuthor

	if err := templates.Lookup("about.html").Execute(w, p); err != nil {
		log.Printf("error: %s", err)
	}
}

func loadAuthors() {
	authorByCrc = make(map[string]string)
	crcByAuthor = make(map[string]string)
	authorByURLName = make(map[string]string)
	urlNameByAuthor = make(map[string]string)
	infoByAuthor = make(map[string]*AuthorInfo)
	sortedAuthors = make([]string, 0)
	for _, author := range classifier.Categories() {
		sortedAuthors = append(sortedAuthors, author)
		crc := crc32.ChecksumIEEE([]byte(author))
		hexCrc := strconv.FormatUint(uint64(crc), 16)
		//XXX temporary
		//log.Printf("%s = %s", author, hexCrc)
		//----
		authorByCrc[hexCrc] = author
		crcByAuthor[author] = hexCrc
		authorByURLName[authorNameToURLSafeName(author)] = author
		urlNameByAuthor[author] = authorNameToURLSafeName(author)
		for _, info := range authorInfos {
			if info.Name == author {
				infoByAuthor[author] = info
			}
		}
	}
	sort.Strings(sortedAuthors)
	//XXX temporary
	log.Printf("number of words: %d", classifier.NumberOfTokens())
}

func serveStaticFile(filename string) http.HandlerFunc {
	data, err := embeddedFS.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write(data)
	})
}

func Serve(addr string) {
	// Parse templates.
	templates = template.Must(template.ParseFS(embeddedFS, "templates/*.html"))

	// Load authors.
	loadAuthors()

	// Add handlers.
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/b/", makeResultHandler(false))
	http.HandleFunc("/s/", makeResultHandler(true))
	http.HandleFunc("/w/", writerRedirectHandler)
	http.HandleFunc("/writer/", writerHandler)
	http.HandleFunc("/about/", aboutHandler)
	http.HandleFunc("/api", apiHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.FS(embeddedFS))))
	http.HandleFunc("/favicon.ico", serveStaticFile("static/favicon.ico"))
	http.HandleFunc("/robots.txt", serveStaticFile("static/robots.txt"))

	// Launch server.
	if addr == "" {
		log.Printf("Started FastCGI server")
		fcgi.Serve(nil, nil)
	} else {
		log.Printf("Started HTTP server at %s", addr)
		if err := http.ListenAndServe(addr, nil); err != nil {
			log.Fatal(err)
		}
	}
}
