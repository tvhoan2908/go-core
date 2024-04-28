-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE categories (
	id bigserial NOT NULL,
	created_at timestamp NULL,
	updated_at timestamp NULL,
	deleted_at timestamp NULL,
	name varchar(255) NOT NULL,
	slug varchar(255) NOT NULL,
	description text NULL,
	user_id int8 NULL,
	parent_id int8 NULL,
	CONSTRAINT categories_pkey PRIMARY KEY (id),
	CONSTRAINT categories_slug_key UNIQUE (slug)
);
CREATE TABLE media (
	id bigserial NOT NULL,
	file_name varchar(255) NOT NULL,
	file_mime varchar(255) NULL,
	file_size int8 NULL,
	file_type int2 DEFAULT 1 NULL, -- Loai File-Image,File,Video,Audio...
	"path" text NOT NULL,
	user_id int8 NULL,
	created_at timestamptz NULL,
	CONSTRAINT media_pkey PRIMARY KEY (id)
);
CREATE TABLE users (
	id bigserial NOT NULL,
	created_at timestamptz NULL,
	updated_at timestamptz NULL,
	deleted_at timestamptz NULL,
	username varchar(255) NOT NULL,
	"password" text NOT NULL,
	email varchar(255) NULL,
	full_name text NULL,
	status int2 DEFAULT 1 NULL, -- 1-Visible,2-Banned,3-Disabled
	account_type int2 DEFAULT 2 NULL, -- 1-Administrator,2-Normal Account
	token_expired_at timestamptz NULL, -- Thoi gian Token het han
	CONSTRAINT users_email_key UNIQUE (email),
	CONSTRAINT users_pkey PRIMARY KEY (id),
	CONSTRAINT users_username_key UNIQUE (username)
);
CREATE INDEX idx_users_deleted_at ON users USING btree (deleted_at);
CREATE INDEX user_idx_token_expired ON users USING btree (token_expired_at);
-- Column comments
COMMENT ON COLUMN users.status IS '1-Visible,2-Banned,3-Disabled';
COMMENT ON COLUMN users.account_type IS '1-Administrator,2-Normal Account';
COMMENT ON COLUMN users.token_expired_at IS 'Thoi gian Token het han';

CREATE TABLE roles (
	id bigserial NOT NULL,
	created_at timestamptz NULL,
	updated_at timestamptz NULL,
	deleted_at timestamptz NULL,
	"name" varchar(255) NOT NULL,
	description text NULL,
	user_id int8 NULL,
	CONSTRAINT roles_pkey PRIMARY KEY (id)
);
CREATE INDEX idx_roles_deleted_at ON roles USING btree (deleted_at);
CREATE TABLE user_roles (
	user_id int8 NOT NULL,
	role_id int8 NOT NULL,
	CONSTRAINT user_roles_pkey PRIMARY KEY (user_id, role_id)
);

CREATE TABLE modules (
	id bigserial NOT NULL,
	created_at timestamptz NULL,
	updated_at timestamptz NULL,
	deleted_at timestamptz NULL,
	"name" varchar(255) NOT NULL,
	description text NULL,
	CONSTRAINT modules_name_key UNIQUE (name),
	CONSTRAINT modules_pkey PRIMARY KEY (id)
);
CREATE INDEX idx_modules_deleted_at ON modules USING btree (deleted_at);

CREATE TABLE permissions (
	id bigserial NOT NULL,
	created_at timestamptz NULL,
	updated_at timestamptz NULL,
	deleted_at timestamptz NULL,
	"name" varchar(255) NOT NULL,
	description text NULL,
	module_id int8 NULL,
	CONSTRAINT permissions_name_key UNIQUE (name),
	CONSTRAINT permissions_pkey PRIMARY KEY (id),
  CONSTRAINT module_id_fk FOREIGN KEY (module_id) REFERENCES modules (id)
);
CREATE INDEX idx_permissions_deleted_at ON permissions USING btree (deleted_at);

CREATE TABLE role_permissions (
	role_id int8 NOT NULL,
	permission_id int8 NOT NULL,
	CONSTRAINT role_permissions_pkey PRIMARY KEY (role_id, permission_id)
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE IF EXISTS role_permissions;
DROP TABLE IF EXISTS permissions;
DROP TABLE IF EXISTS modules;
DROP TABLE IF EXISTS user_roles;
DROP TABLE IF EXISTS roles;
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS categories;
DROP TABLE IF EXISTS media;
-- +goose StatementEnd
