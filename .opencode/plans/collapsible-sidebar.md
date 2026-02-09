# Implementation Plan: Collapsible Sidebar Navigation

## Overview
Transform the static sidebar navigation into an expandable/collapsible directory tree where:
- All directories start collapsed by default
- The active path is auto-expanded
- Users can click directories to toggle expand/collapse
- No localStorage, minimal JavaScript

## Files to Modify

### 1. kew.go - render_nav() function (lines 109-146)

**Current behavior:** Generates flat nested `<ul>` lists with `<a>` tags

**New behavior:** 
- Wrap directory items in clickable `<span>` elements with toggle indicators (▶/▼)
- Add server-side logic to detect if current page is in a directory's subtree
- Auto-expand directories in the active path
- Wrap children in `<div class="nav-children">` with display:none/block

**Key changes:**
- Replace simple `<a>` for directories with `<span class="nav-toggle">` containing the link
- Add onclick="toggleNav(this)" handler
- Show ▼ for expanded directories (in active path), ▶ for collapsed
- Set style="display: block" for active path, "display: none" for others

**Implementation logic:**
```go
// Check if current page is in this directory's subtree
inPath := false
if c.Path != "" && cur != "" {
    curDir := cur
    if idx := strings.LastIndex(cur, "/"); idx != -1 {
        curDir = cur[:idx]
    }
    if strings.HasPrefix(curDir, c.Path) || cur == c.Path {
        inPath = true
    }
}
```

### 2. public/style.css - Add tree styles

**New CSS rules:**
```css
/* Tree toggle styling */
.nav-toggle {
    cursor: pointer;
    user-select: none;
}

.nav-toggle a {
    margin-left: 4px;
}

/* Directory indicator arrows */
.nav-toggle.collapsed::before {
    content: "▶ ";
    display: inline;
}

.nav-toggle.expanded::before {
    content: "▼ ";
    display: inline;
}

/* Children container */
.nav-children {
    margin-left: 20px;
}
```

### 3. public/template.html - Add JavaScript

**Add before closing </body> tag:**
```html
<script>
function toggleNav(el) {
    el.classList.toggle('expanded');
    el.classList.toggle('collapsed');
    const children = el.nextElementSibling;
    if (children) {
        children.style.display = children.style.display === 'none' ? 'block' : 'none';
    }
}
</script>
```

### 4. example/style.css and example/template.html
- Apply identical changes to maintain consistency

## Testing Steps

1. Rebuild site with `go run kew.go` (or appropriate build command)
2. Load any page - verify all directories are collapsed except active path
3. Click a directory - verify it expands and shows ▼
4. Click again - verify it collapses and shows ▶
5. Navigate to a nested page - verify path auto-expands on load

## Rollback
If issues arise, revert to previous commit or restore original files from backup.
