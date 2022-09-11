<a name="readme-top"></a>
![img](https://img.shields.io/badge/semver-2.0.0-green) [![Conventional Commits](https://img.shields.io/badge/Conventional%20Commits-1.0.0-yellow.svg)](https://conventionalcommits.org) [![pre-commit](https://img.shields.io/badge/pre--commit-enabled-brightgreen?logo=pre-commit&logoColor=white)](https://github.com/pre-commit/pre-commit) [![codecov](https://codecov.io/gh/JonasMuehlmann/bntp.go/branch/main/graph/badge.svg?token=ngxdtOfn1f)](https://codecov.io/gh/JonasMuehlmann/bntp.go)

# bntp.go
Libraries and CLIs for my personal all-in-one productivity system including components like bookmarks, notes, todos, projects, etc. 

Neovim integration available at https://github.com/JonasMuehlmann/bntp.nvim

Suggestions and ideas are welcome.
<p align="right">(<a href="#readme-top">back to top</a>)</p>

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
- **Function > aesthetics** because aesthetic customization wastes development time and distracts users
## Installation and setup
`go get github.com/JonasMuehlmann/bntp.go`

Determine OS-specific user configuration directories
```bash
bntp.go config paths
# /home/jonas/.config/bntp
# /home/jonas
```

 Set up a database
```bash
sqlite3 ~/.config/bntp/bntp_db.sql < $(bntp.go config get-schema sqlite)
```

Test your installation (Formatting through [`jq`](https://github.com/stedolan/jq))
```bash
bntp.go bookmark add '{"id": 1, "url": "example.com"}'
bntp.go bookmark list | jq
# [
#   {
#     "created_at": "0001-01-01T00:00:00Z",
#     "updated_at": "0001-01-01T00:00:00Z",
#     "deleted_at": null,
#     "url": "example.com",
#     "title": "foo",
#     "tagIDs": [],
#     "bookmark_type": null
#     "id": 1,
#   }
# ]

bntp.go bookmark remove --filter "BookmarkFilterUntitled"
# {"numAffectedRecords":1}

bntp.go bookmark list
# Should complain about non-existent data
```
<p align="right">(<a href="#readme-top">back to top</a>)</p>

# Usage

## Means of interaction

bntp.go's architecture has clear separation of concerns and abstractions allowing for extensive modularization (and plugin usage in the future) and usage at various levels of abstraction.

### As a schema

The simplest way to interact with bntp.go is by writing raw SQL against the [schema](schema/bntp_sqlite.sql), allowing very flexible interaction, even with different programming languages.

### As a library
#### As an ORM

bntp.go implements and ORM based on [sqlboiler](https://github.com/volatiletech/sqlboiler) for [various DBMS'](https://github.com/JonasMuehlmann/bntp.go/tree/main/model/repository).
Take the [`sqlite3` ORM](https://github.com/JonasMuehlmann/bntp.go/tree/main/model/repository/sqlite3) as an example.
The ORM is defined by most of the files not ending in `*_repository.go` or `*_repository_test.go`, e.g. [`bookmarks.go`](https://github.com/JonasMuehlmann/bntp.go/blob/main/model/repository/sqlite3/bookmarks.go).

#### Through a repository

Repositories are defined by the `*_repository.go` (e.g. [`bookmark_repostiroy.go`](https://github.com/JonasMuehlmann/bntp.go/blob/main/model/repository/bookmark_repository.go)) files in the [repository package](https://github.com/JonasMuehlmann/bntp.go/tree/main/model/repository) and implemented for various DBMS' in the packages sub directory's (e.g. [`sqlite3`](https://github.com/JonasMuehlmann/bntp.go/tree/main/model/repository/sqlite3)) `*_repository.go` files (e.g. [`bookmark_repository`](https://github.com/JonasMuehlmann/bntp.go/blob/main/model/repository/sqlite3/bookmark_repository.go)).

The interface based repository pattern allows seamless modularization, easy extension and even plugin-based implementations.
Whetever the repository is an in-memory KV-store, a remote postrgres DB or a filesystem based implementation is only a matter of specification in the config.

Repositories offer high level operations like `DeleteWhere()`, `GetAll()`, etc.
Their inputs/output are [domain models](https://github.com/JonasMuehlmann/bntp.go/tree/main/model/domain), but get translated to repository-specific data structures internally.

As inputs, they use repository-agnostic:
- entity models (e.g. [`Bookmark`](https://github.com/JonasMuehlmann/bntp.go/blob/a5673f8d95cba8f2529a92a071c495c3c790a2b5/model/domain/bookmark.go#L33-L44))
- `Filter`s (e.g. [`BookmarkFilter`](https://github.com/JonasMuehlmann/bntp.go/blob/a5673f8d95cba8f2529a92a071c495c3c790a2b5/model/domain/bookmark.go#L211-L231))
- `Updater`s (e.g. [`BookmarkUpdater`](https://github.com/JonasMuehlmann/bntp.go/blob/a5673f8d95cba8f2529a92a071c495c3c790a2b5/model/domain/bookmark.go#L268-L279))
- `Grouper`s (Coming soon)
- `Sorter`s (Coming soon)
- `Limiter`s (Coming soon)
- `MemberSelector`s (Coming soon)

#### Through managers

Managers (e.g. [`BookmarkManager`](https://github.com/JonasMuehlmann/bntp.go/blob/main/bntp/libbookmarks/bookmarkmanager.go)) are again entity-specific components, wrapping the underlying repository (e.g. [`BookmarkRepository`](https://github.com/JonasMuehlmann/bntp.go/blob/main/model/repository/sqlite3/bookmark_repository.go)) and enhancing it with extra logic for:
- Hook execution
- Caching (Coming soon)
- Inter-repository communication (e.g. Updating document contents after altering their entities)

### Through program interfaces

The interaction with bntp.go as a program (instead of a library) can be achieved in multiple ways:
- Through a [CLI](https://github.com/JonasMuehlmann/bntp.go/blob/main/model/repository/sqlite3/bookmark_repository.gogg) (e.g. [`bntp.go bookmark` command](https://github.com/JonasMuehlmann/bntp.go/blob/main/model/repository/sqlite3/bookmark_repository.go)) with TUI elements (Coming soon).
- Through gRPC (Coming soon)
- Through a REST API (Coming soon)

These allow scripting bntp.go to create an even richer feature set, allowing e.g. periodic import of bookmarks through unix cronjobs and the CLI.

### Through UIs
The modular architecture and various possible program interfaces allow building various kinds of UIs for bntp.go with little duplication and clear separation of concerns.
Examples:
- [Neovim integration](https://github.com/JonasMuehlmann/bntp.nvim) (First party-support) or other editor integration
- A web app through the REST API (Not a goal)
- A desktop app through gRPC (Not a goal)


# Features

- [`libdocuments`](https://github.com/JonasMuehlmann/bntp.go/tree/main/bntp/libdocuments) (`Markdown` Documents with bidirectional linking)
- [`libbookmarks`](https://github.com/JonasMuehlmann/bntp.go/tree/main/bntp/libbookmarks)
- [`libtags`](https://github.com/JonasMuehlmann/bntp.go/tree/main/bntp/libtags) (Hierarchical tag structure, allowing infinite nesting of parent-tag/sub-tag relationships)
- libtasks (Graph-based task system, coming soon)
- `(g)RPC` based remote plugins based on [`hashicorp/go-plugin`](https://github.com/hashicorp/go-plugin) (Coming soon)

# Getting help
API documentation available at https://pkg.go.dev/github.com/JonasMuehlmann/bntp.go.

CLI documentation available at  [docs/cli_help/bntp.go.md](docs/cli_help/bntp.go.md).

All other documentation will be in [docs/](docs/).

For further questions, please use the [discussions](https://github.com/JonasMuehlmann/bntp.go/discussions).

If any documentation has errors, is incomplete, confusing or could be improved in any other way, please [open an issue](https://github.com/JonasMuehlmann/bntp.go/issues/new).

When no proper exampls exist for a certain API, try to use the `*_test.go` files as examples in the meantime.

# Contributing

TODO: Reference ARCHITECTURE.md

# License

# Acknowledgement
