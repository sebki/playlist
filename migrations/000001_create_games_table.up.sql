CREATE TABLE IF NOT EXISTS games (
    id bigserial PRIMARY KEY,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    modified_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    title text NOT NULL,
    descr text NOT NULL,
    year_published integer NOT NULL,
    bg_type text[] NOT NULL,
    thumbnail text,
    image text,
    min_player integer NOT NULL,
    max_player integer NOT NULL,
    min_playtime integer NOT NULL,
    max_playtime integer NOT NULL,
    min_age integer NOT NULL,
    max_age integer NOT NULL,
    version integer NOT NULL DEFAULT 1
)