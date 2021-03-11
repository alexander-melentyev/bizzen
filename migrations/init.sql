-- Org

CREATE TABLE org (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR (50) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    creator VARCHAR (50) NOT NULL,
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updater VARCHAR (50) NOT NULL,
    deleted_at TIMESTAMP,
    deleter VARCHAR (50)
);

CREATE TABLE org_history (
    id BIGINT NOT NULL,
    name VARCHAR (50) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    creator VARCHAR (50) NOT NULL,
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updater VARCHAR (50) NOT NULL,
    deleted_at TIMESTAMP,
    deleter VARCHAR (50)
);

CREATE INDEX ON org_history(id, updated_at);

CREATE FUNCTION org_history()
    RETURNS TRIGGER
    LANGUAGE PLPGSQL
    AS
$$
BEGIN
    INSERT INTO org_history (
        id,
        name,
        created_at,
        creator,
        updated_at,
        updater,
        deleted_at,
        deleter
    )
    VALUES (
        NEW.id,
        NEW.name,
        NEW.created_at,
        NEW.creator,
        NEW.updated_at,
        NEW.updater,
        NEW.deleted_at,
        NEW.deleter
    );

    RETURN NEW;
END
$$;

CREATE TRIGGER org_history_trigger
    AFTER INSERT OR UPDATE
    ON org
    FOR EACH ROW
    EXECUTE PROCEDURE org_history();

-- Org struct

CREATE TABLE org_struct (
    id BIGSERIAL PRIMARY KEY,
    org_id BIGINT REFERENCES org (id) NOT NULL,
    name VARCHAR (50) NOT NULL,
    parent_id BIGINT,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    creator VARCHAR (50) NOT NULL,
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updater VARCHAR (50) NOT NULL,
    deleted_at TIMESTAMP,
    deleter VARCHAR (50)
);

CREATE TABLE org_struct_history (
    id BIGINT NOT NULL,
    org_id BIGINT NOT NULL,
    name VARCHAR (50) NOT NULL,
    parent_id BIGINT,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    creator VARCHAR (50) NOT NULL,
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updater VARCHAR (50) NOT NULL,
    deleted_at TIMESTAMP,
    deleter VARCHAR (50)
);

CREATE INDEX ON org_struct_history(org_id, updated_at);

CREATE FUNCTION org_struct_history()
    RETURNS TRIGGER
    LANGUAGE PLPGSQL
    AS
$$
BEGIN
    INSERT INTO org_struct_history (
        id,
        org_id,
        name,
        parent_id,
        created_at,
        creator,
        updated_at,
        updater,
        deleted_at,
        deleter
    )
    VALUES (
        NEW.id,
        NEW.org_id,
        NEW.name,
        NEW.parent_id,
        NEW.created_at,
        NEW.creator,
        NEW.updated_at,
        NEW.updater,
        NEW.deleted_at,
        NEW.deleter
    );

    RETURN NEW;
END
$$;

CREATE TRIGGER org_struct_history_trigger
    AFTER INSERT OR UPDATE
    ON org_struct
    FOR EACH ROW
    EXECUTE PROCEDURE org_struct_history();