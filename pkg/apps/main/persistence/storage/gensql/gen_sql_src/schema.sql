CREATE TABLE users (
  id uuid PRIMARY KEY,
  name_full text NOT NULL,
  phone text NOT NULL,
  bio text
);

CREATE TABLE pings (
  phone_to text NOT NULL,
  phone_from text NOT NULL,
  time_created timestamp NOT NULL
);

-- NOTES:
-- - Warning: time_created in go ignores timezone

-- TODO:
-- - Should we have user_ids in pings ?
