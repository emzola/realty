CREATE TABLE IF NOT EXISTS properties (
  	id bigserial PRIMARY KEY,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    title text NOT NULL,
    description text NOT NULL,
    city text NOT NULL,
    location text NOT NULL,
    latitude numeric NOT NULL,
    longitude numeric NOT NULL,
    type text[] NOT NULL,
    category text[] NOT NULL,
    features hstore NOT NULL,
    price numeric NOT NULL,
    currency text[] NOT NULL,
    nearby hstore NOT NULL,
    amenities text[] NOT NULL,
    version integer NOT NULL DEFAULT 1
);
