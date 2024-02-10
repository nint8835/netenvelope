CREATE TABLE vlans
(
    id   INTEGER PRIMARY KEY,
    tag  INTEGER NOT NULL UNIQUE,
    name TEXT
);

CREATE TABLE prefixes
(
    id          INTEGER PRIMARY KEY,
    bitmask     INTEGER NOT NULL,
    size        INTEGER NOT NULL,
    description TEXT,
    vlan_id     INTEGER,

    FOREIGN KEY (vlan_id) REFERENCES vlans (id)
);

CREATE TABLE addresses
(
    id          INTEGER PRIMARY KEY,
    address     INTEGER NOT NULL,
    description TEXT
);
