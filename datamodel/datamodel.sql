create extension if not exists "pgcrypto";

drop table if exists "user" cascade;
create table "user" (
	id		integer	unique generated always as identity,
	email	text		unique
);

drop table if exists device cascade;
create table device (
	id			integer	unique generated always as identity,
	owner		integer	references "user" (id) on delete cascade,
	serial	text		unique
);

-- bootstrap database with a user
-- insert into "user" (name, email) values ('ayan george', 'ayan@ayan.net');

drop table if exists "authkey" cascade;
create table "authkey" (
	id			integer	unique generated always as identity,
	"user"	integer references "user" (id) on delete cascade,
	key			uuid		default gen_random_uuid()
);

-- generate an authkey for our 'ayan@ayan.net' user
-- insert into "authkey" ("user") values ( (select id from "user" where email='ayan@ayan.net')) returning key;

drop table if exists "authtoken";
create table "authtoken" (
	id				integer		generated always as identity,
	authkey		integer		references "authkey" (id) on delete cascade not null,
	token			uuid			default gen_random_uuid(),
	created		timestamp	default localtimestamp,
	accessed	timestamp	default localtimestamp
);

-- request an auth token for our key
-- insert into "authtoken" (authkey) values ( (select id from "authkey" where id=1)) returning key, created, accessed;

drop table if exists "temperature";
create table "temperature" (
	id				integer				generated always as identity not null,
	device		integer				references "device" (id) on delete cascade not null,
	time			timestamp 		default localtimestamp not null,	-- time this measurement was taken
	value			numeric(5,2)	not null													-- assume that we're going to store temps up to 999.99
);

drop table if exists "humidity";
create table "humidity" (
	id				integer				generated always as identity not null,
	device		integer				references "device" (id) on delete cascade not null,
	time			timestamp			default localtimestamp not null,-- time this measurement was taken
	value			numeric(5,2)	not null											-- assume that we're going to store temps up to 999.99
);


drop table if exists "co2";
create table "co2" (
	id				integer				generated always as identity not null,
	device		integer				references "device" (id) on delete cascade not null,
	time			timestamp			default localtimestamp not null,	-- time this measurement was taken
	value			numeric(5,2)	not null											-- assume that we're going to store temps up to 999.99
);

drop table if exists "health";
create table "health" (
	id				integer				generated always as identity not null,
	device		integer				references "device" (id) on delete cascade not null,
	time			timestamp			default localtimestamp not null,	-- time this measurement was taken
	value			text					check (char_length(value) <= 150) not null
);
