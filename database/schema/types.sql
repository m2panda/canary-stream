-- Enum data types declaration --
CREATE TYPE user_role AS ENUM (
  'admin',
  'listener'
);

CREATE TYPE member_status AS ENUM (
  'active',
  'ex'
);

CREATE TYPE collection_format AS ENUM (
  'album',
  'ep',
  'single',
  'playlist'
);
