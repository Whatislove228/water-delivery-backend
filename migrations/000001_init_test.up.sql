CREATE TABLE IF NOT EXISTS schema_migrations_test (
                                                      id BIGSERIAL PRIMARY KEY,
                                                      created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
    );