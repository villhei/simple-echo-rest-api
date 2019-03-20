-- 
CREATE EXTENSION pgcrypto;

DROP TABLE shares;
CREATE TABLE IF NOT EXISTS shares (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  name TEXT NOT NULL,
  symbol TEXT NOT NULL,
  UNIQUE(symbol)
);

