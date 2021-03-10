CREATE TABLE organizations (
    id BIGSERIAL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    creator VARCHAR (50) NOT NULL,
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updater VARCHAR (50) NOT NULL,
    deleted_at TIMESTAMP,
    deleter VARCHAR (50)
);

CREATE TABLE organizations_history (
    row_id BIGSERIAL PRIMARY KEY,
    organization_id INTEGER NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    creator VARCHAR (50) NOT NULL,
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updater VARCHAR (50) NOT NULL,
    deleted_at TIMESTAMP,
    deleter VARCHAR (50)
);

CREATE INDEX ON organizations_history(organization_id, updated_at);

CREATE FUNCTION organizations_history()
    RETURNS TRIGGER
    LANGUAGE PLPGSQL
    AS
$$
    BEGIN
        INSERT INTO organizations_history (
            organization_id,
            created_at,
            creator,
            updated_at,
            updater,
            deleted_at,
            deleter
        )
        VALUES (
            NEW.id,
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

CREATE TRIGGER organizations_history_trigger
    AFTER INSERT OR UPDATE
    ON organizations
    FOR EACH ROW
    EXECUTE PROCEDURE organizations_history();