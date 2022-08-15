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

var github = `<!DOCTYPE html>
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

var blacklight = `<!DOCTYPE html>
<html lang="en">
  <head>
	GAJS
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="chrome=1">
    <meta name="HandheldFriendly" content="True">
    <meta name="MobileOptimized" content="320">
    <meta name="viewport" content="width=device-width, initial-scale=1.0, maximum-scale=1.0, user-scalable=0">

    <link href='https://fonts.googleapis.com/css?family=Open+Sans:400|Old+Standard+TT:400' rel='stylesheet' type='text/css'>
    <title>TITLE</title>

<style>
  * {
    border:0;
    font:inherit;
    font-size:100%;
    vertical-align:baseline;
    margin:0;
    padding:0;
    color: black;
  }

  body {
    font-family:'Open Sans', 'Myriad Pro', Myriad, sans-serif;
    font-size:17px;
    line-height:160%;
    color:#1d1313;
    max-width:700px;
    margin:auto;
  }

  p {
    margin: 20px 0;
  }

  a img {
    border:none;
  }

  img {
    margin: 10px auto 10px auto;
    max-width: 100%;
    display: block;
  }

  pre, code {
    font: 12px Consolas, "Liberation Mono", Menlo, Courier, monospace;
    background-color: #f7f7f7;
  }

  code {
    font-size: 12px;
    padding: 4px;
  }

  pre {
    margin-top: 0;
    margin-bottom: 16px;
    word-wrap: normal;
    padding: 16px;
    overflow: auto;
    font-size: 85%;
    line-height: 1.45;
  }

  pre>code {
    padding: 0;
    margin: 0;
    font-size: 100%;
    word-break: normal;
    white-space: pre;
    background: transparent;
    border: 0;
  }

  pre code {
    display: inline;
    max-width: auto;
    padding: 0;
    margin: 0;
    overflow: visible;
    line-height: inherit;
    word-wrap: normal;
    background-color: transparent;
    border: 0;
  }

  pre code::before,
  pre code::after {
    content: normal;
  }

  em,q,em,dfn {
    font-style:italic;
  }

  .sans,html .gist .gist-file .gist-meta {
    font-family:"Open Sans","Myriad Pro",Myriad,sans-serif;
  }

  .mono,pre,code,tt,p code,li code {
    font-family:Menlo,Monaco,"Andale Mono","lucida console","Courier New",monospace;
  }

  .heading,.serif,h1,h2,h3 {
    font-family:"Old Standard TT",serif;
  }

  strong {
    font-weight:600;
  }

  q:before {
    content:"\201C";
  }

  q:after {
    content:"\201D";
  }

  del,s {
    text-decoration:line-through;
  }

  blockquote {
    font-family:"Old Standard TT",serif;
    text-align:center;
    padding:50px;
  }

  blockquote p {
    display:inline-block;
    font-style:italic;
  }

  blockquote:before,blockquote:after {
    font-family:"Old Standard TT",serif;
    content:'\201C';
    font-size:35px;
    color:#403c3b;
  }

  blockquote:after {
    content:'\201D';
  }

  hr {
    width:40%;
    height: 1px;
    background:#403c3b;
    margin: 25px auto;
  }

  h1 {
    font-size:35px;
  }

  h2 {
    font-size:28px;
  }

  h3 {
    font-size:22px;
    margin-top:18px;
  }

  h1 a,h2 a,h3 a {
    text-decoration:none;
  }

  h1,h2 {
    margin-top:28px;
  }

  #sub-header, time {
    color:#403c3b;
    font-size:13px;
  }

  #sub-header {
    margin: 0 4px;
  }

  #nav h1 a {
    font-size:35px;
    color:#1d1313;
    line-height:120%;
  }

  .posts_listing a,#nav a {
    text-decoration: none;
  }

  li {
    margin-left: 20px;
  }

  ul li {
    margin-left: 5px;
  }

  ul li {
    list-style-type: none;
  }
  ul li:before {
    content:"\00BB \0020";
  }

  #nav ul li:before, .posts_listing li:before {
    content:'';
    margin-right:0;
  }

  #content {
    text-align:left;
    width:100%;
    font-size:15px;
    padding:60px 0 80px;
  }

  #content h1,#content h2 {
    margin-bottom:5px;
  }

  #content h2 {
    font-size:25px;
  }

  #content .entry-content {
    margin-top:15px;
  }

  #content time {
    margin-left:3px;
  }

  #content h1 {
    font-size:30px;
  }

  .highlight {
    margin: 10px 0;
  }

  .posts_listing {
    margin:0 0 50px;
  }

  .posts_listing li {
    margin:0 0 25px 15px;
  }

  .posts_listing li a:hover,#nav a:hover {
    text-decoration: underline;
  }

  #nav {
    text-align:center;
    position:static;
    margin-top:60px;
  }

  #nav ul {
    display: table;
    margin: 8px auto 0 auto;
  }

  #nav li {
    list-style-type:none;
    display:table-cell;
    font-size:15px;
    padding: 0 20px;
  }

  #links {
    margin: 50px 0 0 0;
  }

  #links :nth-child(2) {
    float:right;
  }

  #not-found {
    text-align: center;
  }

  #not-found a {
    font-family:"Old Standard TT",serif;
    font-size: 200px;
    text-decoration: none;
    display: inline-block;
    padding-top: 225px;
  }

  @media (max-width: 750px) {
    body {
      padding-left:20px;
      padding-right:20px;
    }

    #nav h1 a {
      font-size:28px;
    }

    #nav li {
      font-size:13px;
      padding: 0 15px;
    }

    #content {
      margin-top:0;
      padding-top:50px;
      font-size:14px;
    }

    #content h1 {
      font-size:25px;
    }

    #content h2 {
      font-size:22px;
    }

    .posts_listing li div {
      font-size:12px;
    }
  }

  @media (max-width: 400px) {
    body {
      padding-left:20px;
      padding-right:20px;
    }

    #nav h1 a {
      font-size:22px;
    }

    #nav li {
      font-size:12px;
      padding: 0 10px;
    }

    #content {
      margin-top:0;
      padding-top:20px;
      font-size:12px;
    }

    #content h1 {
      font-size:20px;
    }

    #content h2 {
      font-size:18px;
    }

    .posts_listing li div{
      font-size:12px;
    }
  }
</style>
  </head>

  <body>
    <section id=nav>
      <h1><a href="/">SITENAME</a></h1>
    </section>

<section id=content>
  <h1>TITLE</h1>
  <div id=sub-header>DATE</div>
  <div class="entry-content">REPLACE</div>
</body>
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

	re := regexp.MustCompile(`<!--THEME:(\s|\S)*?-->`)
	tm := re.FindString(string(b))
	tm = strings.Replace(strings.Replace(tm, "<!--THEME:", "", 1), "-->", "", 1)

	if tm == "" {
		if bytes.Contains(b, []byte("<!--TOC-->")) {
			out, err := exec.Command("sh", "-c", fmt.Sprintf("mdtoc1d %s", os.Args[1])).Output()
			if err != nil {
				log.Println(err)
				return
			}
			b = bytes.Replace(b, []byte("<!--TOC-->"), out, 1)
		}

		b = github_flavored_markdown.Markdown(b)

		name := strings.TrimRight(os.Args[1], ".md") + ".article"
		if len(os.Args) >= 3 {
			name = os.Args[2]
		}

		if err := os.WriteFile(name, b, 0644); err != nil {
			log.Println(err)
			return
		}
		return
	}

	html := ""
	if tm == "github" {
		html = github

		re := regexp.MustCompile(`<!--(UA-|G-).*?-->`)
		ga := re.FindString(string(b))
		ga = strings.Replace(strings.Replace(ga, "<!--", "", 1), "-->", "", 1)
		s := ""
		if ga != "" {
			s = fmt.Sprintf(gajs, ga, ga)
		}
		html = strings.Replace(html, "GAJS", s, 1)
	}
	if tm == "blacklight" {
		html = blacklight

		re := regexp.MustCompile(`<!--(UA-|G-).*?-->`)
		ga := re.FindString(string(b))
		ga = strings.Replace(strings.Replace(ga, "<!--", "", 1), "-->", "", 1)
		s := ""
		if ga != "" {
			s = fmt.Sprintf(gajs, ga, ga)
		}
		html = strings.Replace(html, "GAJS", s, 1)

		re = regexp.MustCompile(`<!--SITENAME:(\s|\S)*?-->`)
		sn := re.FindString(string(b))
		sn = strings.Replace(strings.Replace(sn, "<!--SITENAME:", "", 1), "-->", "", 1)
		html = strings.Replace(html, "SITENAME", sn, 1)

		re = regexp.MustCompile(`<!--DATE:(\s|\S)*?-->`)
		dt := re.FindString(string(b))
		dt = strings.Replace(strings.Replace(dt, "<!--DATE:", "", 1), "-->", "", 1)
		html = strings.Replace(html, "DATE", dt, 1)
	}

	s, err := bytes.NewBuffer(b).ReadString('\n')
	if err != nil {
		log.Println(err)
		return
	}
	if !strings.HasPrefix(s, "# ") {
		log.Println("First line must starts with # ")
		return
	}
	title := strings.TrimSpace(strings.Replace(s, "# ", "", 1))
	if tm == "blacklight" {
		b = bytes.Replace(b, []byte(s), []byte(""), 1)
	}

	if bytes.Contains(b, []byte("<!--TOC-->")) {
		out, err := exec.Command("sh", "-c", fmt.Sprintf("mdtoc1d %s", os.Args[1])).Output()
		if err != nil {
			log.Println(err)
			return
		}
		b = bytes.Replace(b, []byte("<!--TOC-->"), out, 1)
	}

	b = bytes.Replace([]byte(html), []byte("REPLACE"), github_flavored_markdown.Markdown(b), 1)
	b = bytes.Replace(b, []byte("TITLE"), []byte(title), 2)

	name := strings.TrimRight(os.Args[1], ".md") + ".article"
	if len(os.Args) >= 3 {
		name = os.Args[2]
	}

	if err := os.WriteFile(name, b, 0644); err != nil {
		log.Println(err)
		return
	}
}
