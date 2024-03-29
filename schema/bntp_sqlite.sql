PRAGMA foreign_keys = ON;

-- TODO: It might be worth exploring tables which support caching, ideas below:
/*
CREATE TABLE table_cache_states
(
    table_name  NOT NULL TEXT PRIMARY KEY,
    is_dirty    NOT NULL INTEGER
);

CREATE TABLE dirty_entries
(
    dirty_rows_pk TEXT NOT NULL PRIMARY KEY,
    table_name    TEXT NOT NULL
);
*/

CREATE TABLE tags
(
    id         INTEGER    PRIMARY KEY NOT NULL,
    tag        TEXT       NOT NULL,
    parent_tag INTEGER    REFERENCES tags(id) DEFERRABLE INITIALLY DEFERRED,
    -- Stores list of parent ids from root to self
    -- e.g. "1;2;3"
    path       TEXT       NOT NULL UNIQUE,
    -- Stores lis of children ids
    -- e.g. "1;2;3"
    children   TEXT       NOT NULL
);

CREATE TABLE bookmark_types
(
    id            INTEGER PRIMARY KEY NOT NULL,
    bookmark_type TEXT    NOT NULL UNIQUE
);

CREATE TABLE bookmarks
(
    id               INTEGER   PRIMARY KEY NOT NULL,
    is_read          INTEGER   NOT  NULL DEFAULT 0,
    title            TEXT      UNIQUE,
    url              TEXT      NOT NULL UNIQUE,
    bookmark_type_id INTEGER   REFERENCES bookmark_types(id) DEFERRABLE INITIALLY DEFERRED,
    is_collection    INTEGER   NOT NULL DEFAULT 0,
    created_at       TIMESTAMP NOT NULL,
    updated_at       TIMESTAMP NOT NULL,
    deleted_at       TIMESTAMP
);

CREATE TABLE document_types
(
    id            INTEGER PRIMARY KEY NOT NULL,
    document_type TEXT    NOT NULL UNIQUE
);

CREATE TABLE documents
(
    id               INTEGER   PRIMARY KEY NOT NULL,
    path             TEXT      NOT NULL UNIQUE,
    document_type_id INTEGER   REFERENCES document_types(id) DEFERRABLE INITIALLY DEFERRED,
    created_at       TIMESTAMP NOT NULL,
    updated_at       TIMESTAMP NOT NULL,
    deleted_at       TIMESTAMP
);

CREATE TABLE links
(
    source_id      INTEGER  NOT NULL REFERENCES documents(id) DEFERRABLE INITIALLY DEFERRED,
    destination_id INTEGER  NOT NULL REFERENCES documents(id) DEFERRABLE INITIALLY DEFERRED,

    PRIMARY KEY(source_id, destination_id),
    CHECK(source_id != destination_id)
);

CREATE TABLE bookmark_contexts
(
    bookmark_id INTEGER  NOT NULL REFERENCES bookmarks(id) DEFERRABLE INITIALLY DEFERRED,
    tag_id      INTEGER  NOT NULL REFERENCES tags(id) DEFERRABLE INITIALLY DEFERRED,

    PRIMARY KEY(tag_id, bookmark_id)
);

CREATE TABLE document_contexts
(
    document_id INTEGER  NOT NULL REFERENCES documents(id) DEFERRABLE INITIALLY DEFERRED,
    tag_id      INTEGER  NOT NULL REFERENCES  tags(id) DEFERRABLE INITIALLY DEFERRED,

    PRIMARY KEY(tag_id, document_id)
);
