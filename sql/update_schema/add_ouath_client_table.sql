
CREATE TABLE IF NOT EXISTS oauth_clients (
  id     TEXT  NOT NULL,
  secret TEXT  NOT NULL,
  domain TEXT  NOT NULL,
  data   JSONB NOT NULL,
  CONSTRAINT oauth_clients_pkey PRIMARY KEY (id)
);