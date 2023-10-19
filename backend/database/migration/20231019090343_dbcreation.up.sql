CREATE TABLE "weekness" (
  "weeknessid" serial PRIMARY KEY,
  "weekness" varchar
);

CREATE TABLE "origin" (
  "originid" serial PRIMARY KEY,
  "origin" varchar
);

CREATE TABLE "organization" (
  "organizationid" serial PRIMARY KEY,
  "organization" varchar,
  "type" varchar
);

CREATE TABLE "weapon" (
  "weaponid" serial PRIMARY KEY,
  "weapon" varchar,
  "weaponclass" varchar
);

CREATE TABLE "power" (
  "powerid" serial PRIMARY KEY,
  "power" varchar
);

CREATE TABLE "user_detail" (
  "id" serial PRIMARY KEY,
  "username" varchar UNIQUE,
  "password" varchar,
  "name" varchar,
  "birthdate" timestamp,
  "deathdate" timestamp,
  "originid" integer,
  "organizationid" integer,
  "retired" boolean,
  CONSTRAINT fk_originid FOREIGN KEY (originid) REFERENCES origin (originid),
  CONSTRAINT fk_organizationid FOREIGN KEY (organizationid) REFERENCES organization (organizationid)
);

CREATE TABLE "fight" (
  "fightid" serial PRIMARY KEY,
  "hero_user_id" integer,
  "villain_user_id" integer,
  "winner_user_id" integer,
  CONSTRAINT fk_hero_user_id FOREIGN KEY (hero_user_id) REFERENCES user_detail (id),
  CONSTRAINT fk_villain_user_id FOREIGN KEY (villain_user_id) REFERENCES user_detail (id),
  CONSTRAINT fk_winner_user_id FOREIGN KEY (winner_user_id) REFERENCES user_detail (id)
);

CREATE TABLE "userweapon" (
  "userid" integer,
  "weaponid" integer,
  CONSTRAINT fk_userid FOREIGN KEY (userid) REFERENCES user_detail (id),
  CONSTRAINT fk_weaponid FOREIGN KEY (weaponid) REFERENCES weapon (weaponid)
);

CREATE TABLE "userpower" (
  "userid" integer PRIMARY KEY,
  "powerid" integer,
  CONSTRAINT fk_userpower_id FOREIGN KEY (userid) REFERENCES user_detail (id),
  CONSTRAINT fk_powerid FOREIGN KEY (powerid) REFERENCES "power" (powerid)
);

CREATE TABLE "userweekness" (
  "userid" integer,
  "weeknessid" integer,
  CONSTRAINT fk_userweekness_id FOREIGN KEY (userid) REFERENCES user_detail (id),
  CONSTRAINT fk_weeknessid FOREIGN KEY (weeknessid) REFERENCES weekness (weeknessid)
);

CREATE TABLE "user_class" (
  "userid" integer,
  "class" varchar,
  CONSTRAINT fk_userclass_id FOREIGN KEY (userid) REFERENCES user_detail (id)
);

CREATE TABLE "user_type" (
  "userid" integer,
  "hero" boolean,
  "villain" boolean,
  CONSTRAINT fk_usertype_id FOREIGN KEY (userid) REFERENCES user_detail (id)
);

CREATE TABLE "mentor" (
  "userid" integer,
  "mentorid" integer,
  CONSTRAINT fk_usermentor_id FOREIGN KEY (userid) REFERENCES user_detail (id),
  CONSTRAINT fk_mentorid FOREIGN KEY (mentorid) REFERENCES user_detail (id)
);

CREATE TABLE "sidekick" (
  "userid" integer,
  "sidekickid" integer,
  CONSTRAINT fk_usersidekick_id FOREIGN KEY (userid) REFERENCES user_detail (id),
  CONSTRAINT fk_sidekickid FOREIGN KEY (sidekickid) REFERENCES user_detail (id)
);