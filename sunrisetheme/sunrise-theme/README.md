# Sunrise Hugo Theme

A Hugo theme based on sunrise color animations, designed for portfolio and project showcase websites with distinctive layouts for different content types.

## Features

### Four Distinctive Layout Types

1. **Landing Page** - Singular page with sunrise animation background
2. **Project Pages** - Full narrative layout with rich content blocks
3. **Simple Project Pages** - Minimal layout for lighter content
4. **Project List Pages** - Filterable grid/list view with tag-based filtering

### Advanced Tag System

- **Nested Tags** - Support for hierarchical tags (e.g., "writing" and "writing/poetry")
- **Tag Pages** - Individual pages for each tag with optional about content
- **Smart Filtering** - Parent tags show content from child tags, but not vice versa

### Content Management

- **Priority System** - `isPriority` flag for highlighting important content
- **Custom Sorting** - `sortPriority` for controlling display order
- **Call-to-Actions** - Configurable CTA, CTALink, and CTAPreamble for each item

## Installation

1. Clone or download this theme to your Hugo site's `themes` directory:
```bash
git clone [repo-url] themes/sunrise-theme
```

2. Add the theme to your site configuration:
```toml
theme = "sunrise-theme"
```

3. Copy the example configuration:
```bash
cp themes/sunrise-theme/exampleSite/config.toml config.toml
```

## Content Structure

### Projects

Create projects in `content/projects/`:

```yaml
---
title: "Project Name"
subtitle: "Optional subtitle"
type: "projects"
tags: ["tag1", "tag2", "nested/tag"]
isPriority: false
sortPriority: 1
isSimple: false
CTA: "View Project"
CTALink: "https://example.com"
CTAPreamble: "Check it out:"
---
```

### Tags

Create tag pages in `content/tags/`:

```yaml
---
title: "tag-name"
about: "Description of this tag category"
isPriority: false
CTA: "Learn More"
CTALink: "/about"
---
```

### Nested Tags

For nested tags like "writing/poetry":
- Create `content/tags/writing.md` for the parent
- Create `content/tags/writing-poetry.md` with `title: "writing/poetry"`
- The theme automatically handles the hierarchy

## Configuration

### Site Parameters

```toml
[params]
  # Landing page CTA
  landingCTA = "View My Work"
  landingCTALink = "/projects"
  landingCTAPreamble = "Ready to explore?"
```

### Taxonomies

```toml
[taxonomies]
  tag = "tags"
```

## Layout Types

### Project Pages

Regular projects use the full layout with:
- Header with title, subtitle, and tags
- Full content area
- Related projects section
- CTA section

### Simple Projects

Set `isSimple: true` for a minimal layout without related projects.

### Landing Page

The home page (`content/_index.md`) uses a special layout with sunrise animation background.

## Customization

### CSS

- `static/css/landing.css` - Landing page styles
- `static/css/project.css` - Project page styles  
- `static/css/project-list.css` - List view styles
- `static/css/tag-*.css` - Tag page styles

### JavaScript

- `static/js/project-list.js` - Filtering and view switching
- `static/js/nested-tags.js` - Nested tag functionality

### Sunrise Animation

The sunrise animation from the original color processing work can be integrated by:
1. Copying SVG animations to `static/images/`
2. Including them in the landing page layout
3. Adding CSS animations in `landing.css`

## Development

This theme is designed to be easily updatable and maintains a clean separation between:
- Layout templates (`layouts/`)
- Styling (`static/css/`)
- Functionality (`static/js/`)
- Content structure (`archetypes/`)

## License

MIT License - see LICENSE file for details.