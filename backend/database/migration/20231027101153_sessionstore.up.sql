CREATE TABLE "session_detail" (
  "session_id" varchar PRIMARY KEY,
  "user_id" int,
  "expires" timestamp,
  CONSTRAINT fk_userid_session FOREIGN KEY (user_id) REFERENCES user_detail (id)
);