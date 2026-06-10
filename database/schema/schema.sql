-- DB tables creation --
CREATE TABLE IF NOT EXISTS status (
  _id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  name VARCHAR(15) NOT NULL,
  slug VARCHAR(15) UNIQUE NOT NULL
);

CREATE TABLE IF NOT EXISTS users (
  _id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  username VARCHAR(50) UNIQUE NOT NULL,
  password VARCHAR(255) NOT NULL,
  role user_role NOT NULL,
  picture UUID,
  token CHARACTER(32),
  token_exp TIMESTAMP,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  status_id UUID NOT NULL REFERENCES status(_id)
);

CREATE TABLE IF NOT EXISTS genres (
  _id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  name VARCHAR(255) NOT NULL,
  slug VARCHAR(255) UNIQUE NOT NULL
);

CREATE TABLE IF NOT EXISTS artists (
  _id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  stagename VARCHAR(255) UNIQUE NOT NULL,
  bio TEXT,
  picture UUID UNIQUE
);

CREATE TABLE IF NOT EXISTS members (
  _id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  status_id UUID NOT NULL REFERENCES status(_id),
  artist_id UUID NOT NULL REFERENCES artists(_id),
  group_id UUID NOT NUll REFERENCES artists(_id)
);

CREATE TABLE IF NOT EXISTS songs (
  _id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  title VARCHAR(255) NOT NULL,
  audio UUID UNIQUE NOT NULL,
  lyrics UUID UNIQUE NOT NULL,
  duration SMALLINT,
  explicit BOOLEAN,
  artist_id UUID REFERENCES artists(_id)
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
  cover UUID,
  artist_id UUID REFERENCES artists(_id)
);

CREATE TABLE IF NOT EXISTS collection_songs (
  _id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  song_index SMALLINT NOT NULL,
  song_id UUID NOT NULL REFERENCES songs(_id),
  collection_id UUID NOT NULL REFERENCES collections(_id)
);

CREATE TABLE IF NOT EXISTS playlist (
  _id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  name VARCHAR(255) NOT NULL,
  history BOOLEAN NOT NULL,
  picture UUID,
  user_id UUID NOT NULL REFERENCES users(_id),
  status_id UUID NOT NULL REFERENCES status(_id)
);

CREATE TABLE IF NOT EXISTS playlist_songs (
  _id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  song_index SMALLINT NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  user_id UUID NOT NULL REFERENCES users(_id),
  song_id UUID NOT NULL REFERENCES songs(_id),
  collection_id UUID NOT NULL REFERENCES collections(_id),
  playlist_id UUID NOT NULL REFERENCES playlist(_id)
);

CREATE TABLE IF NOT EXISTS playlist_collabs (
  _id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  user_id UUID NOT NULL REFERENCES users(_id),
  playlist_id UUID NOT NULL REFERENCES playlist(_id),
  status_id UUID NOT NULL REFERENCES status(_id)
);
