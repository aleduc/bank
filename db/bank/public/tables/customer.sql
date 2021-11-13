-- liquibase formatted sql

-- changeset aleduc:1 stripComments:false runOnChange:false endDelimiter:#
create table customer
(
  id bigserial primary key,
  t_login text,
  t_password text,
  t_name text,
  t_surname text,
  t_mobile_number text,
  t_address text
);
-- #
-- rollback drop table public.customer;

-- changeset aleduc:9 stripComments:false runOnChange:true endDelimiter:#
create index idx_login_password ON public.customer using btree (t_login, t_password);
-- #
-- rollback drop index public.index_login_password;
