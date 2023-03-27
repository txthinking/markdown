package main

import (
	"bytes"
	"log"
	"os"
	"regexp"
	"strings"
	"text/template"

	ghtoc "github.com/ekalinin/github-markdown-toc.go"
	"github.com/shurcooL/github_flavored_markdown"
)

func main() {
	b, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Println(err)
		os.Exit(1)
		return
	}
	md := string(b)
	ga := strings.Replace(strings.Replace(regexp.MustCompile(`<!--(UA-|G-).*-->`).FindString(md), "<!--", "", 1), "-->", "", 1)
	s, err := bytes.NewBufferString(md).ReadString('\n')
	if err != nil {
		log.Println(err)
		os.Exit(1)
		return
	}
	if !strings.HasPrefix(s, "# ") {
		log.Println("First line must starts with # ")
		os.Exit(1)
		return
	}
	title := strings.TrimSpace(strings.Replace(s, "# ", "", 1))
	md = strings.Replace(md, s, "", 1)
	var toc, sidebar string
	if strings.Contains(md, "<!--TOC-->") || strings.Contains(md, "<!--SIDEBAR-->") {
		file, err := os.CreateTemp(os.TempDir(), "_")
		if err != nil {
			log.Println(err)
			os.Exit(1)
			return
		}
		defer os.Remove(file.Name())
		if err := os.WriteFile(file.Name(), []byte(md), 0644); err != nil {
			log.Println(err)
			os.Exit(1)
			return
		}
		l := *(ghtoc.NewGHDoc(file.Name(), false, 0, 0, false, "", 2, false).GetToc())
		for i, _ := range l {
			ll := strings.Split(l[i], "](")
			l[i] = ll[0] + "](" + strings.ReplaceAll(ll[1], "_", "-")
		}
		if strings.Contains(md, "<!--TOC-->") {
			toc = strings.Join(l, "\n")
		}
		if strings.Contains(md, "<!--SIDEBAR-->") {
			sidebar = strings.Join(l, "\n")
		}
	}

	b, err = static.ReadFile("static/template.html")
	if err != nil {
		log.Println(err)
		os.Exit(1)
		return
	}
	tmpl, err := template.New("").Parse(string(b))
	if err != nil {
		log.Println(err)
		os.Exit(1)
		return
	}
	name := strings.TrimRight(os.Args[1], ".md") + ".html"
	if len(os.Args) >= 3 {
		name = os.Args[2]
	}
	out, err := os.OpenFile(name, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
		os.Exit(1)
		return
	}
	defer out.Close()
	if err != nil {
		log.Println(err)
		os.Exit(1)
		return
	}
	err = tmpl.Execute(out, map[string]interface{}{
		"ga":          ga,
		"title":       title,
		"titlehtml":   string(github_flavored_markdown.Markdown([]byte("# " + title))),
		"tochtml":     string(github_flavored_markdown.Markdown([]byte(toc))),
		"sidebarhtml": string(github_flavored_markdown.Markdown([]byte(sidebar))),
		"bodyhtml":    string(github_flavored_markdown.Markdown([]byte(md))),
	})
	if err != nil {
		log.Println(err)
		os.Exit(1)
		return
	}
}
