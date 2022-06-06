

-- TODO: It might be worth exploring tables which support caching, ideas below:
/*
CREATE TABLE table_cache_states
(
    table_name  NOT NULL VARCHAR(255) PRIMARY KEY,
    is_dirty    NOT NULL INTEGER
);

CREATE TABLE dirty_entries
(
    dirty_rows_pk VARCHAR(255) NOT NULL PRIMARY KEY,
    table_name    VARCHAR(255) NOT NULL
);
*/

CREATE TABLE tags
(
    id         INTEGER    PRIMARY KEY,
    parent_tag INTEGER    ,
    tag        VARCHAR(255)       NOT NULL UNIQUE,
    depth      INTEGER    NOT NULL DEFAULT 0
);

CREATE TABLE bookmark_types
(
    id   INTEGER PRIMARY KEY,
    Type VARCHAR(255)    NOT NULL UNIQUE
);

CREATE TABLE bookmarks
(
    id               INTEGER   PRIMARY KEY,
    is_read          INTEGER   NOT  NULL DEFAULT 0,
    title            VARCHAR(255)      UNIQUE,
    url              VARCHAR(255)      NOT NULL UNIQUE,
    bookmark_type_id INTEGER   ,
    is_collection    INTEGER   NOT NULL DEFAULT 0,
    created_at       DATETIME NOT NULL,
    updated_at       DATETIME NOT NULL,
    deleted_at       DATETIME
);

CREATE TABLE document_types
(
    id            INTEGER PRIMARY KEY,
    document_type VARCHAR(255)    NOT NULL UNIQUE
);

CREATE TABLE documents
(
    id               INTEGER   PRIMARY KEY,
    path             VARCHAR(255)      NOT NULL UNIQUE,
    document_type_id INTEGER   ,
    created_at       DATETIME NOT NULL,
    updated_at       DATETIME NOT NULL,
    deleted_at       DATETIME
);

CREATE TABLE links
(
    source_id      INTEGER  NOT NULL ,
    destination_id INTEGER  NOT NULL ,

    PRIMARY KEY(source_id, destination_id),
    CHECK(source_id != destination_id)
);

CREATE TABLE bookmark_contexts
(
    bookmark_id INTEGER  NOT NULL ,
    tag_id      INTEGER  NOT NULL ,

    PRIMARY KEY(tag_id, bookmark_id)
);

CREATE TABLE document_contexts
(
    document_id INTEGER  NOT NULL ,
    tag_id      INTEGER  NOT NULL ,

    PRIMARY KEY(tag_id, document_id)
);
