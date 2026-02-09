package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type NavNode struct {
	Name     string
	Path     string
	Files    []NavNode
	Children []NavNode
}

func build_nav(dir string, root string) (NavNode, bool) {
	var node NavNode
	node.Name = title_from_name(filepath.Base(dir))

	entries, err := os.ReadDir(dir)
	if err != nil {
		return node, false
	}

	for _, e := range entries {
		full := filepath.Join(dir, e.Name())

		if e.IsDir() {
			child, ok := build_nav(full, root)
			if ok {
				node.Children = append(node.Children, child)
			}
			continue
		}

		if strings.HasSuffix(e.Name(), ".md") {
			if e.Name() == "index.md" {
				rel_dir, _ := filepath.Rel(root, dir)
				if rel_dir == "." {
					node.Path = "index.html"
				} else {
					node.Path = rel_dir + "/index.html"
				}
				continue
			}
			rel, _ := filepath.Rel(root, full)
			html := strings.TrimSuffix(rel, ".md") + ".html"

			node.Files = append(node.Files, NavNode{
				Name: title_from_name(e.Name()),
				Path: html,
			})
		}
	}

	if len(node.Files) == 0 && len(node.Children) == 0 && node.Path == "" {
		return node, false
	}

	return node, true
}

func copy_file(src string, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	return err
}

func markdown_to_html(path string) (string, error) {
	cmd := exec.Command(
		"lowdown",
		"-T", "html",
		"--html-no-skiphtml",
		"--html-no-escapehtml",
	)

	in, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer in.Close()

	var out strings.Builder
	cmd.Stdin = in
	cmd.Stdout = &out
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return "", err
	}

	return out.String(), nil
}

func render_nav(n NavNode, b *strings.Builder, cur string) {
	b.WriteString("<ul>\n")

	for _, f := range n.Files {
		p := f.Path
		if !strings.HasPrefix(p, "/") {
			p = "/" + p
		}

		isCurrent := p == cur
		if isCurrent {
			b.WriteString(`<li><a href="` + p + `" class="thisPage">»<i> ` + f.Name + `</i></a></li>` + "\n")
		} else {
			b.WriteString(`<li><a href="` + p + `">› ` + f.Name + `</a></li>` + "\n")
		}
	}

	for _, c := range n.Children {
		// Check if current page is in this directory's subtree
		inPath := false
		isCurrent := false
		if cur != "" {
			curNorm := strings.TrimPrefix(cur, "/")

			if c.Path != "" {
				// Directory has an index page
				childNorm := strings.TrimPrefix(c.Path, "/")
				isCurrent = curNorm == childNorm

				// Get directory of current page
				curDir := curNorm
				if idx := strings.LastIndex(curNorm, "/"); idx != -1 {
					curDir = curNorm[:idx]
				}

				// Get directory of this child node
				childDir := childNorm
				if idx := strings.LastIndex(childNorm, "/"); idx != -1 {
					childDir = childNorm[:idx]
				}

				if strings.HasPrefix(curDir, childDir) || isCurrent {
					inPath = true
				}
			} else {
				// Directory has no index page, check by name
				curDir := curNorm
				if idx := strings.LastIndex(curNorm, "/"); idx != -1 {
					curDir = curNorm[:idx]
				}
				// Check if current page's directory matches this directory's name
				if curDir == c.Name {
					inPath = true
				}
			}
		}

		b.WriteString("<li>")

		// Directory name with toggle
		if inPath {
			b.WriteString(`<span class="nav-toggle expanded" onclick="toggleNav(this)">`)
		} else {
			b.WriteString(`<span class="nav-toggle collapsed" onclick="toggleNav(this)">`)
		}

		if c.Path != "" {
			p := c.Path
			if !strings.HasPrefix(p, "/") {
				p = "/" + p
			}
			if inPath {
				b.WriteString(`<a href="` + p + `" class="thisPage">»<i> ` + c.Name + `/</i></a>`)
			} else {
				b.WriteString(`<a href="` + p + `">› ` + c.Name + `/</a>`)
			}
		} else {
			if inPath {
				b.WriteString(`»<i> ` + c.Name + `/</i>`)
			} else {
				b.WriteString(`› ` + c.Name + `/`)
			}
		}
		b.WriteString("</span>")

		// Children container with display state
		if inPath {
			b.WriteString(`<div class="nav-children" style="display: block;">`)
		} else {
			b.WriteString(`<div class="nav-children" style="display: none;">`)
		}
		render_nav(c, b, cur)
		b.WriteString("</div></li>\n")
	}

	b.WriteString("</ul>\n")
}

func replace_md_references(s string) string {
	r := strings.NewReplacer(
		/* common cases */
		".md)", ".html)",
		".md\"", ".html\"",
		".md'", ".html'",
		".md)", ".html)",
		".md#", ".html#",
		".md>", ".html>",
		".md ", ".html ",
		".md,", ".html,",
	)
	return r.Replace(s)
}

func title_from_name(name string) string {
	name = strings.TrimSuffix(name, ".md")
	name = strings.ReplaceAll(name, "-", " ")
	return name
}

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "usage: relic <in> <out>\n")
		os.Exit(1)
	}

	src := os.Args[1]
	out := os.Args[2]

	/* load template */
	tmpl, err := os.ReadFile(filepath.Join(src, TemplateFile))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	/* build nav */
	rootnav, _ := build_nav(src, src)

	/* walk site */
	err = filepath.WalkDir(src, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}

		rel, _ := filepath.Rel(src, path)
		outpath := filepath.Join(out, rel)

		if d.IsDir() {
			return os.MkdirAll(outpath, 0755)
		}

		if strings.HasSuffix(path, ".md") {
			html, err := markdown_to_html(path)
			if err != nil {
				return err
			}
			html = replace_md_references(html)

			relhtml := strings.TrimSuffix(rel, ".md") + ".html"
			cur := relhtml
			if !strings.HasPrefix(cur, "/") {
				cur = "/" + cur
			}
			var navbuf strings.Builder
			render_nav(rootnav, &navbuf, cur)

			page := string(tmpl)
			page = strings.ReplaceAll(page, "{{TITLE}}", SiteTitle)
			page = strings.ReplaceAll(page, "{{NAV}}", navbuf.String())
			page = strings.ReplaceAll(page, "{{CONTENT}}", html)
			page = strings.ReplaceAll(page, "{{FOOTER}}", FooterText)
			page = strings.ReplaceAll(page, "{{SITE_NAME}}", SiteName)
			page = strings.ReplaceAll(page, "{{HEADER_LINKS}}", HeaderLinks)
			page = strings.ReplaceAll(page, "{{HEADER_SUBTITLE}}", HeaderSubtitle)

			outpath = strings.TrimSuffix(outpath, ".md") + ".html"
			return os.WriteFile(outpath, []byte(page), 0644)
		}

		return copy_file(path, outpath)
	})

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
