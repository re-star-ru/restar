create table "user" ( id bigserial primary key, "name" text );

create table
    image (
        id bigserial primary key,
        "name" text,
        path text
    );

create table product ( id bigserial primary key, name text );

drop table if exists diagnostic;

create table
    diagnostic (
        id bigserial,
        "version" serial,
        primary key (id, "version"),
        created_at timestamptz not null,
        updated_at timestamptz not null,
        defined_number text not null,
        sku text not null
    );

drop view if exists diagnostic_view;

CREATE VIEW DIAGNOSTIC_VIEW AS 
	select distinct on (id) *
	from diagnostic
	order by id desc, version
DESC; 

-- explain select * from diagnostic_view limit 500;

-- create schema history;

-- create table history.diagnostic (

-- 	"version" serial

-- ) inherits (public.diagnostic);

-- create table diagnostic_product (

-- 	product_id bigserial references product on delete restrict,

-- 	order_id bigserial references diagnostic on delete cascade,

-- 	quantity integer,

-- 	primary key (product_id, order_id)

-- );

-- create table diagnostic_image (

-- 	image_id bigserial references image on delete restrict,

-- 	order_id bigserial references diagnostic on delete cascade,

-- 	primary key (image_id, order_id)

-- );