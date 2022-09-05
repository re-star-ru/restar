create table "user" (
  id bigserial primary key,
  "name" text
);

create table image (
  id bigserial primary key,
  "name" text,
  path text
);

create table product (
  id bigserial primary key,
  name text
);

create table diagnostic (
  id bigserial primary key,
  createdAt timestamptz not null,
  updatedAt timestamptz not null ,

  definedNumber text not null,
  SKU text not null
);

create table diagnostic_product (
	product_id bigserial references product on delete restrict,
	order_id bigserial references diagnostic on delete cascade,
	quantity integer,
	primary key (product_id, order_id)
);

create table diagnostic_image (
	image_id bigserial references image on delete restrict,
	order_id bigserial references diagnostic on delete cascade,
	primary key (image_id, order_id)
);
