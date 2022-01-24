package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strings"

	"github.com/shurcooL/github_flavored_markdown"
)

var html = `<!DOCTYPE html>
<html>
    <head>
        <meta charset="UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1">
        <link rel="stylesheet" href="https://cdn.jsdelivr.net/gh/sindresorhus/github-markdown-css@master/github-markdown.css">
        <title>TITLE</title>
		GAJS
        <style>
                .markdown-body {
                    box-sizing: border-box;
                    min-width: 200px;
                    max-width: 980px;
                    margin: 0 auto;
                    padding: 45px;
                }

                @media (max-width: 767px) {
                    .markdown-body {
                        padding: 15px;
                    }
                }
		</style>
        <body class="markdown-body">REPLACE</body>
</html>`
var gajs = `
		<script async src="https://www.googletagmanager.com/gtag/js?id=%s"></script>
		<script>
			window.dataLayer = window.dataLayer || [];
			function gtag(){dataLayer.push(arguments);}
			gtag('js', new Date());
			gtag('config', '%s');
		</script>
`

func main() {
	b, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Println(err)
		return
	}

	re := regexp.MustCompile(`<!--(UA-|G-).*-->`)
	ga := re.FindString(string(b))
	gaid := strings.Replace(strings.Replace(ga, "<!--", "", 1), "-->", "", 1)
	s := ""
	if ga != "" {
		s = fmt.Sprintf(gajs, gaid, gaid)
	}
	html = strings.Replace(html, "GAJS", s, 1)

	s, err = bytes.NewBuffer(b).ReadString('\n')
	if err != nil {
		log.Println(err)
		return
	}
	if !strings.HasPrefix(s, "# ") {
		log.Println("First line must starts with # ")
		return
	}
	title := strings.TrimSpace(strings.Replace(s, "# ", "", 1))
	if bytes.Contains(b, []byte("<!--TOC-->")) {
		out, err := exec.Command("sh", "-c", fmt.Sprintf("mdtoc1d %s", os.Args[1])).Output()
		if err != nil {
			log.Println(err)
			return
		}
		b = bytes.Replace(b, []byte("<!--TOC-->"), out, 1)
	}
	b = bytes.Replace([]byte(html), []byte("REPLACE"), github_flavored_markdown.Markdown(b), 1)
	b = bytes.Replace(b, []byte("TITLE"), []byte(title), 1)
	name := strings.TrimRight(os.Args[1], ".md") + ".article"
	if len(os.Args) >= 3 {
		name = os.Args[2]
	}
	if err := os.WriteFile(name, b, 0644); err != nil {
		log.Println(err)
		return
	}
}
