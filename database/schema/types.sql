-- Enum data types declaration --
CREATE TYPE user_role AS ENUM (
   'admin',
   'manager',
   'listener'
);

CREATE TYPE collection_format AS ENUM (
   'album',
   'ep',
   'single'
);
