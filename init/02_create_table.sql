CREATE TABLE users {
  user_id char(12) not null,
  name_jp char(20) not null,
  name_en char(40) not null,
  nickname char(30),
  constraint PK_users primary key (user_id)
} ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE INDEX IDX_users_01
ON users(user_id);
