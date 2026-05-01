CREATE TABLE products (
                          id BIGSERIAL PRIMARY KEY,
                          name TEXT NOT NULL,
                          description TEXT,
                          price_cents INTEGER NOT NULL CHECK (price_cents >= 0),
                          volume_ml INTEGER NOT NULL CHECK (volume_ml > 0),
                          is_active BOOLEAN NOT NULL DEFAULT TRUE,
                          created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
                          updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);