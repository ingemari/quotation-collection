CREATE TABLE IF NOT EXISTS quotes (
                                      id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                                      author TEXT NOT NULL,
                                      quote TEXT NOT NULL,
                                      created_at TIMESTAMP NOT NULL DEFAULT NOW()
);
