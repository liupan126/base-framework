-- ----------------------------
-- 数据库升级日志表
-- ----------------------------
CREATE TABLE migration_history (
  id integer PRIMARY KEY NOT NULL ,
  created_at timestamp NOT NULL  default CURRENT_TIMESTAMP,
  version CHAR(50) NOT NULL ,
  data text NOT NULL,
  UNIQUE(version)
);
COMMENT ON TABLE migration_history IS '数据库升级';
comment on column migration_history.id is '主键';
comment on column migration_history.created_at is '创建时间';
comment on column migration_history.version is '升级版本';
comment on column migration_history.data is '升级内容';

CREATE SEQUENCE migration_history_id_seq START WITH 1 INCREMENT BY 1 NO MINVALUE NO MAXVALUE CACHE 1;

alter table migration_history alter column id set default nextval('migration_history_id_seq');

-- ----------------------------
-- 系统用户
-- ----------------------------

create  table users (
  id integer PRIMARY KEY NOT NULL ,
  created_at timestamp NOT NULL  default CURRENT_TIMESTAMP,
  updated_at timestamp NOT NULL  default CURRENT_TIMESTAMP ,
  username CHAR(50) NOT NULL,
  UNIQUE(username)
);

COMMENT ON TABLE users IS '系统用户';
comment on column users.created_at is '创建时间';
comment on column users.updated_at is '更新时间';
CREATE SEQUENCE users_id_seq START WITH 1 INCREMENT BY 1 NO MINVALUE NO MAXVALUE CACHE 1;

alter table users alter column id set default nextval('users_id_seq');