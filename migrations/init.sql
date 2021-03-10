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