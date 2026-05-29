-- DB tables creation --
CREATE TABLE IF NOT EXISTS users (
  _id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  username VARCHAR(50) UNIQUE NOT NULL,
  password VARCHAR(255) NOT NULL,
  role user_role NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS genres (
  _id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  name VARCHAR(255) NOT NULL,
  alt_name VARCHAR(255),
  slug VARCHAR(255) UNIQUE NOT NULL
);

CREATE TABLE IF NOT EXISTS artists (
  _id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  stage_name VARCHAR(255) NOT NULL,
  is_group BOOLEAN NOT NULL,
  bio TEXT,
  picture CHARACTER(12) UNIQUE
);

CREATE TABLE IF NOT EXISTS members (
  _id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  status member_status,
  artist_id UUID NOT NULL REFERENCES artists(_id),
  group_id UUID NOT NUll REFERENCES artists(_id)
);

CREATE TABLE IF NOT EXISTS songs (
  _id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  title VARCHAR(255) NOT NULL,
  audio CHARACTER(12) UNIQUE NOT NULL,
  lyrics CHARACTER(12) UNIQUE NOT NULL,
  duration SMALLINT,
  explicit BOOLEAN,
  owner_id UUID REFERENCES artists(_id)
);

CREATE TABLE IF NOT EXISTS song_ft (
  _id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  ft_id UUID NOT NULL REFERENCES artists(_id),
  song_id UUID NOT NULL REFERENCES songs(_id)
);

CREATE TABLE IF NOT EXISTS song_genres (
  _id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  genre_id UUID NOT NULL REFERENCES genres(_id),
  song_id UUID NOT NULL REFERENCES songs(_id)
);

CREATE TABLE IF NOT EXISTS collections (
  _id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  title VARCHAR(255) NOT NULL,
  format collection_format NOT NULL,
  release SMALLINT,
  description TEXT,
  cover CHARACTER(12),
  artist_id UUID REFERENCES artists(_id),
  user_id UUID REFERENCES users(_id)
);

CREATE TABLE IF NOT EXISTS collection_songs (
  _id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  song_index SMALLINT NOT NULL,
  song_id UUID NOT NULL REFERENCES songs(_id),
  collection_id UUID NOT NULL REFERENCES collections(_id)
);

CREATE TABLE IF NOT EXISTS collection_collab (
  _id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  collab_id UUID NOT NULL REFERENCES users(_id),
  playlist_id UUID NOT NULL REFERENCES collections(_id)
);
