![img](https://img.shields.io/badge/semver-2.0.0-green) [![Conventional Commits](https://img.shields.io/badge/Conventional%20Commits-1.0.0-yellow.svg)](https://conventionalcommits.org) [![pre-commit](https://img.shields.io/badge/pre--commit-enabled-brightgreen?logo=pre-commit&logoColor=white)](https://github.com/pre-commit/pre-commit)

# bntp.go
Libraries and CLIs for my personal all-in-one productivity system including components like bookmarks, notes, todos, projects, etc. 

Neovim integration available at https://github.com/JonasMuehlmann/bntp.nvim
Suggestions and ideas are welcome.

## Installation

```go get github.com/JonasMuehlmann/bntp.go```

Documentation available at https://pkg.go.dev/github.com/JonasMuehlmann/bntp.go.
The CLI is documented at docs/cli_help/

# Why?

Being a huge productivity nerd and life long learner, I value having efficient tools and workflows but have not found any one tool, which satisfied my needs.
Combining multiple tools is suboptimal too, even though using Notion and Todoist has serverd me quite alright so far.
Being a software developer, I thought, I would use my skills to my advantage and build the very tool I am trying to find, implementing just what I need, exactly how I want it and combining the best of all the other tools out there while having complete control and not depending on others.

Tools and systems, which have or, after research and considerations, might influence this project:
- Notion
- Roam Research
- Obsidian
- Todoist
- Jira
- The *Zettelkasten* Method
- David Allens *Getting Things Done* (GTD)

# Goals

- **Synergy** by using a single tool serving as a **complete** productivity system
- **No reliance on external services** (local data synced through my own means)
- Integrate with and built upon my **neovim/terminal based workflows and tools** (no bloated and unergonomic graphical or web based interfaces)
- **Scriptable** (e.g. periodic bookmark imports)
- **Modular** (Interfaces built **on top of** CLIs built **on top of different** libraries)
- **Hackable** through a GRPC based plugin framework (Decide how you want to store data; Local FS? Cloud Storage? Import from new file format? Put something in a Redis cache?)
- **Customizable** through the Plugin framework and extensive customization and hooking into the core code (e.g. trigger upload after change or trigger versioning, if you don't store your data in a git repo)
- **Portable** through import/export to/from common file formats (`YML`,`CSV`,`Markdown`) using a local `SQLite` database only where it makes sense
- **Discoverable** through context given by a hierarchical tag structure fuzzy searching and using filters
- **Extensive association** through links/backlinks between documents and declaration of sources and *related pages*
- **High performance** by working completely offline, having no bloat and using a high performance backend written in `go`
- **Quick and smooth usage** because of all of the above(especially the neovim frontend)

# Features
- Bidirectional linking between documents (creating a link to `B` inside `A` creates a backlink to `A` in `B`)
- Hierarchical tag structure with `YML` import/export (tags can have sub tags: `software_development` has sub tags `software_engineering` and `computer_science`, searching for `software_developmebt` also shows documents with the tags `software_engineering` and `computer_science`)

# Content types
- Bookmarks with `CSV` import/export
- Todos with `CSV` import/export
- Documents (`Markdown` based files categorized for different purposes)
    - Notes
    - Projects (notes with associated todos)
    - Roadmaps (collection of projects, notes, todos and bookmarks)
