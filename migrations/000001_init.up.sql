CREATE SCHEMA IF NOT EXISTS quest_of_today;

CREATE TABLE quest_of_today.questions (
  id uuid NOT NULL,
  question varchar(255) NOT NULL,
  answer varchar(255) NOT NULL,
  source_text varchar(255) NOT NULL,
  source_url varchar(255) NOT NULL,
  dato date NOT NULL,
  PRIMARY KEY (id)
);
