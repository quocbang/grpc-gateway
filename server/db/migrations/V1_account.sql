-- account table
CREATE TABLE IF NOT EXISTS grpc_gateway_database.account (
  username text NOT NULL PRIMARY KEY, 
	email text NOT NULL, 
	is_user_verified boolean NOT NULL DEFAULT false, 
	password bytea NOT NULL, 
	role smallint NOT NULL 
);

-- account: primary key index
CREATE UNIQUE INDEX "idx_account_email"
	ON "grpc_gateway"."account"
	USING btree (email);

-- account: email unique index
CREATE UNIQUE INDEX "idx_email"
	ON "grpc_gateway"."account"
	USING btree (email);
CREATE TABLE IF NOT EXISTS account_verify (

);