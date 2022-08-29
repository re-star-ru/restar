create table "user" (
  id serial,
  "name" text
);

create table image (
  id serial,
  parent_id integer not null,
  "name" text,
  path text
);

create table diagnostic (
  id serial,
  definedNumber text,
  SKU text
)

