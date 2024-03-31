CREATE TABLE News (
  id SERIAL PRIMARY KEY,
  title TEXT NOT NULL,
  content TEXT NOT NULL
);

CREATE TABLE NewsCategories (
  id SERIAL NOT NULL PRIMARY KEY,
  news_id SERIAL NOT NULL REFERENCES News(id) ON DELETE CASCADE
);
