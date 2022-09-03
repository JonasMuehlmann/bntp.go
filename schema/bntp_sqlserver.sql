

-- TODO: It might be worth exploring tables which support caching, ideas below:
/*
CREATE TABLE table_cache_states
(
    table_name  NOT NULL VARCHAR(255) PRIMARY KEY,
    is_dirty    NOT NULL BIGINT
);

CREATE TABLE dirty_entries
(
    dirty_rows_pk VARCHAR(255) NOT NULL PRIMARY KEY,
    table_name    VARCHAR(255) NOT NULL
);
*/

CREATE TABLE tags
(
    id         BIGINT    PRIMARY KEY,
    tag        VARCHAR(255)       NOT NULL,
    parent_tag BIGINT    REFERENCES tags(id) DEFERRABLE INITIALLY DEFERRED,
    -- Stores list of parent ids from root to self
    -- e.g. "1;2;3"
    path       VARCHAR(255)       NOT NULL UNIQUE,
    -- Stores lis of children ids
    -- e.g. "1;2;3"
    children   VARCHAR(255)       NOT NULL
);

CREATE TABLE bookmark_types
(
    id            BIGINT PRIMARY KEY,
    bookmark_type VARCHAR(255)    NOT NULL UNIQUE
);

CREATE TABLE bookmarks
(
    id               BIGINT   PRIMARY KEY,
    is_read          BIGINT   NOT  NULL DEFAULT 0,
    title            VARCHAR(255)      UNIQUE,
    url              VARCHAR(255)      NOT NULL UNIQUE,
    bookmark_type_id BIGINT   REFERENCES bookmark_types(id) DEFERRABLE INITIALLY DEFERRED,
    is_collection    BIGINT   NOT NULL DEFAULT 0,
    created_at       DATETIME NOT NULL,
    updated_at       DATETIME NOT NULL,
    deleted_at       DATETIME
);

CREATE TABLE document_types
(
    id            BIGINT PRIMARY KEY,
    document_type VARCHAR(255)    NOT NULL UNIQUE
);

CREATE TABLE documents
(
    id               BIGINT   PRIMARY KEY,
    path             VARCHAR(255)      NOT NULL UNIQUE,
    document_type_id BIGINT   REFERENCES document_types(id) DEFERRABLE INITIALLY DEFERRED,
    created_at       DATETIME NOT NULL,
    updated_at       DATETIME NOT NULL,
    deleted_at       DATETIME
);

CREATE TABLE links
(
    source_id      BIGINT  NOT NULL REFERENCES documents(id) DEFERRABLE INITIALLY DEFERRED,
    destination_id BIGINT  NOT NULL REFERENCES documents(id) DEFERRABLE INITIALLY DEFERRED,

    PRIMARY KEY(source_id, destination_id),
    CHECK(source_id != destination_id)
);

CREATE TABLE bookmark_contexts
(
    bookmark_id BIGINT  NOT NULL REFERENCES bookmarks(id) DEFERRABLE INITIALLY DEFERRED,
    tag_id      BIGINT  NOT NULL REFERENCES tags(id) DEFERRABLE INITIALLY DEFERRED,

    PRIMARY KEY(tag_id, bookmark_id)
);

CREATE TABLE document_contexts
(
    document_id BIGINT  NOT NULL REFERENCES documents(id) DEFERRABLE INITIALLY DEFERRED,
    tag_id      BIGINT  NOT NULL REFERENCES  tags(id) DEFERRABLE INITIALLY DEFERRED,

    PRIMARY KEY(tag_id, document_id)
);
