CREATE TABLE users (
    id serial UNIQUE NOT NULL PRIMARY KEY,
    nickname varchar UNIQUE NOT NULL,
    email varchar UNIQUE NOT NULL,
    password varchar UNIQUE NOT NULL
);

CREATE TABLE clothes (
    id serial UNIQUE NOT NULL PRIMARY KEY,
    photo_id varchar UNIQUE NOT NULL,
    owner_id int NOT NULL REFERENCES users(id),
    class varchar,
    brand varchar,
    color int
);
--ALTER TABLE clothes ADD CONSTRAINT pkClothes PRIMARY KEY (id);
--ALTER TABLE clothes ADD CONSTRAINT fkClothesUserId FOREIGN KEY (owner_id) REFERENCES users(id) ON DELETE RESTRICT;

CREATE TABLE looks (
    id serial UNIQUE NOT NULL PRIMARY KEY,
    photo_id int UNIQUE NOT NULL,
    owner_id int NOT NULL,
    description text,
    season varchar,
    temperature_range int4range,
    purpose varchar,
    priority int
);
--ALTER TABLE looks ADD CONSTRAINT pk_looks PRIMARY KEY (id);

CREATE TABLE looks_clothes (
    id serial UNIQUE NOT NULL PRIMARY KEY,
    look_id int NOT NULL REFERENCES looks(id),
    cloth_id int NOT NULL REFERENCES clothes(id)
);
--ALTER TABLE looks_clothes ADD CONSTRAINT pk_looks_clothes PRIMARY KEY (id);
--ALTER TABLE looks_clothes ADD CONSTRAINT fk_looks_clothes_look_id FOREIGN KEY (look_id) REFERENCES looks(id);
--ALTER TABLE looks_clothes ADD CONSTRAINT fk_looks_clothes_cloth_id FOREIGN KEY (cloth_id) REFERENCES clothes(id);
