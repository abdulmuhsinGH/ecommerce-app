
CREATE TABLE IF NOT EXISTS oauth_tokens (
  id     uuid DEFAULT uuid_generate_v4(),
  access TEXT  NOT NULL,
  refresh TEXT  NOT NULL,
  code TEXT NOT NULL,
  expires_in integer,
  created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
  data   JSONB NULL,
  CONSTRAINT oauth_tokens_pkey PRIMARY KEY (id)
);