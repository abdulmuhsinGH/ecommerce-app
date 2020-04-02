
CREATE TABLE IF NOT EXISTS oauth_tokens (
  id uuid DEFAULT uuid_generate_v4(),
  access TEXT  NULL,
  refresh TEXT  NULL,
  code TEXT NULL,
  expires_at TIMESTAMPTZ NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
  data   text NULL,
  CONSTRAINT oauth_tokens_pkey PRIMARY KEY (id)
);