create table "user" (
  id serial,
  "name" text
);

create table image (
  id serial,
  "name" text,
  path text
);

create table diagnostic (
  id serial,
  definedNumber text,
  SKU text,
  images image[] -- hmm, wtf
)

