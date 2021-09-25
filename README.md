# bntp.go
Libraries and CLIs for my personal all-in-one productivity system including components like bookmarks, notes, todos, projects, etc. 

Neovim integration available at https://github.com/JonasMuehlmann/bntp.nvim

# Why?

Being a huge productivity nerd and life long learner, I value having efficient tools and workflows but have not found any one tool, which satisfied my needs.
Combining multiple tools is suboptimal too, even though using Notion and Todoist has serverd me quite alright so far.
Being a software developer, I thought, I would use my skills to my advantage and build the very tool I am trying to find, implementing just what I need, exactly how I want it and combining the best of all the other tools out there while having complete control and not depending on others.

Tools and systems, which have or, after research and considerations, might influence this project:
- Notion
- Roam Research
- Obsidian
- Todoist
- The *Zettelkasten* Method
- David Allens *Getting Things Done* (GTD)

# Goals

- **Synergy** by using a single tool serving as a **complete** productivity system
- **No reliance on external services** (local data synced through my own means)
- Integrate with and built upon my **neovim/terminal based workflows and tools** (no bloated and unergonomic graphical or web based interfaces)
- **Scriptable**
- **Modular** (Interfaces built **on top of** CLIs built **on top of different** libraries)
- **Portable** through import/export to/from and use of common file formats (`YML`,`CSV`,`Markdown`) using a local `SQLite` database only where it makes sense
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
