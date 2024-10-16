CREATE TABLE cafes (
                              id UUID PRIMARY KEY,
                              name VARCHAR(255) NOT NULL,
                              address TEXT NOT NULL,
                              latitude DOUBLE PRECISION NOT NULL,
                              longitude DOUBLE PRECISION NOT NULL,
                              rating DOUBLE PRECISION,
                              hours_of_operation TEXT,
                              has_wifi BOOLEAN,
                              has_outlets BOOLEAN,
                              is_independent BOOLEAN,
                              photo_url TEXT
);