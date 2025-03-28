CREATE TABLE languages
(
    Id   SERIAL PRIMARY KEY,
    Name VARCHAR(255) UNIQUE NOT NULL
);

CREATE TABLE news_outlet
(
    Id           SERIAL PRIMARY KEY,
    Name         VARCHAR(255) UNIQUE NOT NULL,
    QueryUrl     TEXT                NOT NULL,
    HtmlSelector TEXT                NOT NULL,
    LanguageId   INT                 NOT NULL,
    Credibility  INT                 NOT NULL
);

ALTER TABLE news_outlet
    ADD CONSTRAINT fk_language
        FOREIGN KEY (LanguageId) REFERENCES languages (Id) ON DELETE CASCADE;