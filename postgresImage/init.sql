CREATE TABLE doc (
  id SERIAL PRIMARY KEY,
  image TEXT,
  title TEXT,
  description TEXT,
  link TEXT,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE post (
  id SERIAL PRIMARY KEY,
  user_id INTEGER,
  tag TEXT,
  topics_name TEXT,
  md_plain_text TEXT,
  md_html_text TEXT,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  post_image TEXT,
  min_read TEXT
);

CREATE TABLE tag (
  id SERIAL PRIMARY KEY,
  tag TEXT
);
