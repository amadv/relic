# Final Implementation Plan: Rename to "relic"

## Confirmed Decisions

1. **Tagline**: "Bring Out Your Dead"
2. **License**: MIT License (free use/distribute/modify + attribution required)
3. **Install Path**: `~/.local/bin` (user-local, no sudo needed)

## Files to Modify/Rename

### 1. Rename Files
- `kew.go` → `relic.go`
- `kew.1` → `relic.1`

### 2. Update File Contents

#### **go.mod**
```go
module relic
```

#### **relic.go** (line ~230)
```go
fmt.Fprintf(os.Stderr, "usage: relic <in> <out>\n")
```

#### **.gitignore**
```
relic
public/
```

#### **relic.1** (all instances)
- `.Dt KEW 1` → `.Dt RELIC 1`
- All `.Nm kew` → `.Nm relic`
- Update NAME section description

#### **example/index.md**
```markdown
### relic example site index
```

### 3. Create README.md

```markdown
# relic
### Bring Out Your Dead

**relic** is a minimal static site generator written in Go, inspired by the legendary [werc](http://werc.cat-v.org/) by Uriel and the cat-v team from Plan 9.

## Acknowledgment

This project owes its existence to **werc**, the pioneering web system from Plan 9. werc demonstrated that web frameworks can be elegant, minimal, and fast while treating the filesystem as the source of truth.

## Why relic?

While werc is fantastic, it has specific requirements that don't fit every use case:

- **werc** is a dynamic web system requiring a server and Plan 9 utilities (rc-shell)
- **werc** needs a running process to serve content

**relic** takes a different approach:

- **Static generation** - build your site once, deploy anywhere
- **Zero runtime dependencies** - just HTML, CSS, and JavaScript
- **Fast builds** - generate entire sites in milliseconds
- **Simple deployment** - host on any static file server (nginx, Apache, GitHub Pages, etc.)
- **No Plan 9 required** - works on any system with Go

## Philosophy

- werc treats the site as a dynamic system
- relic treats the site as a build output

Both approaches are valid. Choose werc for dynamic, interactive sites. Choose relic for fast, simple, static sites.

## Quick Start

```bash
make build
make run
```

Your site will be generated in `public/`.

## Installation

```bash
# Install to ~/.local/bin (default)
make install

# Or install to custom location
make install PREFIX=/usr/local
```

## Usage

```
relic <input_dir> <output_dir>
```

Example:
```bash
relic example public
```

## Features

- **Automatic navigation** - generates sidebar from directory structure
- **Collapsible tree** - expandable/collapsible directory navigation
- **Active path highlighting** - shows current location in the tree
- **Markdown support** - uses lowdown for fast Markdown processing
- **Simple templates** - customizable HTML templates
- **Zero JS required** - works without JavaScript for basic navigation

## Configuration

Edit `config.go` to customize:
- `SiteName` - site title in header
- `SiteTitle` - HTML page title
- `HeaderSubtitle` - subtitle shown in header
- `HeaderLinks` - navigation links in header
- `FooterText` - footer content

## License

MIT License

Copyright (c) 2025 [Your Name]

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
```

### 4. Create Makefile

```makefile
.PHONY: build clean install run test

BINARY=relic
PREFIX?=$(HOME)/.local
INSTALL_PATH=$(PREFIX)/bin

build:
	go build -o $(BINARY) .

clean:
	rm -f $(BINARY)
	rm -rf public/

install: build
	install -d $(INSTALL_PATH)
	install -m 755 $(BINARY) $(INSTALL_PATH)/

run: build
	./$(BINARY) example public

test: build
	./$(BINARY) example public
	@echo "✓ Site built successfully in public/"
```

## Implementation Steps

1. Rename `kew.go` → `relic.go`
2. Rename `kew.1` → `relic.1`
3. Update `go.mod` module name
4. Update usage message in `relic.go`
5. Update `.gitignore`
6. Update `relic.1` content
7. Update `example/index.md`
8. Create new `README.md`
9. Create `Makefile`
10. Rebuild site to verify

## Post-Implementation

After these changes, the project will be fully renamed to "relic" with:
- ✅ New README acknowledging werc
- ✅ MIT License with attribution
- ✅ Simple Makefile for building
- ✅ User-local install by default (`~/.local/bin`)
- ✅ All references updated

**Ready to proceed?** Say "implement" and I'll execute all changes.
