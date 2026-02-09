# Plan: Rename Project from "kew" to "relic"

## Overview
Rename the project from "kew" to "relic", update all references, create a new README, and add a Makefile. The README should acknowledge this is inspired by werc from Plan 9 while explaining the benefits of relic.

## Files to Modify/Rename

### 1. Rename Files
- `kew.go` → `relic.go`
- `kew.1` → `relic.1`

### 2. Update File Contents

#### **go.mod**
```go
module relic
```

#### **relic.go** (formerly kew.go)
- Line ~230: Change usage message from `"usage: kew <in> <out>\n"` to `"usage: relic <in> <out>\n"`

#### **README.md** - Complete rewrite
New content structure:
- Project name: **relic**
- Tagline: "a static site generator inspired by werc"
- Acknowledge werc by Uriel @ cat-v from Plan 9
- Explain benefits over werc:
  - Static site generator vs dynamic web system
  - No Plan 9 utils required
  - Faster build times
  - Simpler deployment
- Keep it simple and appreciative

#### **.gitignore**
```
relic
public/
```

#### **relic.1** (formerly kew.1)
- Update all `.Nm kew` to `.Nm relic`
- Update `.Dt KEW 1` to `.Dt RELIC 1`
- Update description

#### **example/index.md**
Change title from "kew example" to "relic example"

### 3. Create Makefile

```makefile
.PHONY: build clean install

BINARY=relic
INSTALL_PATH=/usr/local/bin

build:
	go build -o $(BINARY) .

clean:
	rm -f $(BINARY)
	rm -rf public/

install: build
	cp $(BINARY) $(INSTALL_PATH)/

run: build
	./$(BINARY) example public

test: build
	./$(BINARY) example public
	@echo "Site built successfully in public/"
```

### 4. Generated Files (will auto-update on rebuild)
- `public/index.html` - will regenerate with new content

### 5. Optional/Can Skip
- `.opencode/plans/*.md` - documentation references (low priority)
- `.git/config` - update after GitHub repo rename

## README.md Content Structure

```markdown
# relic

**relic** is a minimal static site generator written in Go, inspired by the legendary [werc](http://werc.cat-v.org/) by Uriel and the cat-v team.

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

## Usage

```
relic <input_dir> <output_dir>
```

## Features

- Automatic navigation generation from directory structure
- Collapsible sidebar with expandable tree view
- Markdown to HTML conversion using lowdown
- Simple, customizable templates
- No JavaScript required for basic functionality

## License

[Your license here]
```

## Implementation Steps

1. **Rename files**: `kew.go` → `relic.go`, `kew.1` → `relic.1`
2. **Update go.mod**: Change module name
3. **Update relic.go**: Fix usage message
4. **Create new README.md**: With werc acknowledgment and benefits
5. **Update .gitignore**: Change binary name
6. **Update relic.1**: Fix man page references
7. **Update example/index.md**: Change title
8. **Create Makefile**: Simple build automation
9. **Rebuild site**: Generate new output
10. **Test**: Verify everything works

## Questions for User

1. **Project tagline**: Should we keep "K, Enough of Werc" or change it to something like "Bring Out Your Dead" (matching the subtitle in config.go)?

2. **License**: What license should be mentioned in the README?

3. **Makefile install path**: Is `/usr/local/bin` acceptable, or would you prefer a different default?

4. **Man page**: Should the man page content be updated with new descriptions, or just the name changes?

5. **Logo**: Should the README reference a different logo file, or keep logo.png?

## Rollback Plan

If issues arise:
- Keep git history for easy revert
- All changes are in separate commits
- Generated files can be rebuilt
