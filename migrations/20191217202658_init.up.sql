CREATE TABLE
    album (
        id VARCHAR PRIMARY KEY,
        name VARCHAR NOT NULL,
        created_at TIMESTAMP NOT NULL,
        updated_at TIMESTAMP NOT NULL
    );

CREATE TABLE
    analysis (
        id VARCHAR PRIMARY KEY,
        staId VARCHAR NOT NULL,
        knowledgeMap JSON NOT NULL,
        created_at TIMESTAMP NOT NULL,
        updated_at TIMESTAMP NOT NULL
    );

CREATE TABLE
    showtechs (
        id VARCHAR PRIMARY KEY,
        parsed_data JSON NOT NULL,
        created_at TIMESTAMP NOT NULL,
        updated_at TIMESTAMP NOT NULL
    );