CREATE TABLE esample.users (
  user_id MEDIUMINT NOT NULL AUTO_INCREMENT,
  name_first CHAR(20) NOT NULL,
  name_familly CHAR(40) NOT NULL,
  nickname CHAR(30),
  constraint PK_users primary key (user_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE INDEX IDX_users_01
ON users(user_id);
