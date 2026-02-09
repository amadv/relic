# Implementation Plan: New Layout and Styling

## Overview
Update the site with new flexbox-based layout, red color scheme, and improved navigation indicators. Add header links and site name placeholders.

## Changes Summary

### 1. config.go - Add New Constants
Add constants for header links and site name:
```go
const SiteName = "dérive"
const HeaderLinks = "<a href=\"https://hevel.derivelinux.org\">hevel</a><a href=\"https://ports.derivelinux.org\">ports</a><a href=\"https://pkg.derivelinux.org\">pkg</a>"
const HeaderSubtitle = "is not ready"
```

### 2. kew.go - Update render_nav() Function
**Current behavior:** Uses `@ ` for current page, `: ` for files, `/` for directories

**New behavior:**
- Current page: `»<i> filename</i>` (with italic)
- Other files: `› filename`
- Directories: `dirname/` (unchanged format)

**Key changes in render_nav() (around line 118-122):**
```go
// For files
sym := "› "
style := ""
if p == cur {
    sym = "»<i> "
    style = " class=\"thisPage\""
}
b.WriteString(`<li><a href="` + p + `"` + style + `">` + sym + f.Name + "</a></li>\n")
// Note: Need to close </i> tag for current page
```

**For directories (around line 126-140):**
```go
// Similar logic - use » for current directory, › for others
// Wrap directory name in <i> tags when it's the current page
```

### 3. public/style.css - Complete Replacement
Replace entire file content with the new CSS provided by user.

### 4. public/template.html - New Structure
```html
<!DOCTYPE HTML>
<html>
<head>
    <title>{{TITLE}}</title>
    <link rel="stylesheet" href="/style.css" type="text/css">
    <meta charset="UTF-8">
</head>
<body>
<header>
    <nav>
        <div>{{HEADER_LINKS}}</div>
    </nav>
    <h1><a href="/">{{SITE_NAME}} <span id="headerSubTitle">{{HEADER_SUBTITLE}}</span></a></h1>
</header>

    <nav id="side-bar">
        <div>
            <p class="sideBarTitle">pages:</p>
            {{NAV}}
        </div>
    </nav>

<article>
    {{CONTENT}}
</article>

<footer>
    {{FOOTER}}
</footer>

<script>
function toggleNav(el) {
    el.classList.toggle('expanded');
    el.classList.toggle('collapsed');
    const children = el.nextElementSibling;
    if (children) {
        children.style.display = children.style.display === 'none' ? 'block' : 'none';
    }
</script>
</body>
</html>
```

### 5. example/template.html - Apply Same Changes
Same structure as public/template.html

### 6. example/style.css - Apply Same Changes
Same CSS as public/style.css

### 7. Update Template Processing in kew.go
Around line 254-257, add new replacements:
```go
page = strings.ReplaceAll(page, "{{SITE_NAME}}", SiteName)
page = strings.ReplaceAll(page, "{{HEADER_LINKS}}", HeaderLinks)
page = strings.ReplaceAll(page, "{{HEADER_SUBTITLE}}", HeaderSubtitle)
```

## Questions for User

1. **Current page styling** - In the example HTML, parent directories in the active path are also marked with `»` (not just the current file). Should I:
   - A) Mark only the current file with `»` and others with `›`
   - B) Mark all directories in the active path with `»` (like the example)

2. **Navigation expansion** - Should the collapsible tree behavior be kept with the new layout? The new CSS might need adjustments for the `.nav-toggle` and `.nav-children` classes.

3. **Header links format** - The HeaderLinks constant contains raw HTML. Should I:
   - A) Keep it as a single constant with HTML
   - B) Make it an array/slice of link objects for better structure

## Testing Steps

1. Rebuild with `go build -o kew .`
2. Regenerate site with `./kew example public`
3. Verify:
   - New red color scheme applied
   - Flexbox layout working (sidebar on left, content on right)
   - Header shows site name and links
   - Navigation shows `»` for current page, `›` for others
   - Collapsible tree still functional

## Implementation Order

1. config.go - Add new constants
2. kew.go - Update render_nav() and template processing
3. public/style.css - Replace CSS
4. public/template.html - New HTML structure
5. example/style.css - Replace CSS
6. example/template.html - New HTML structure
7. Test and rebuild
