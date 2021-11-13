-- liquibase formatted sql

-- changeset aleduc:1 stripComments:false runOnChange:false endDelimiter:#
create table account
(
  id bigserial primary key,
  id_customer bigint not null,
  b_is_deleted boolean default false,
  n_number integer,
  n_sort_code integer,
  n_balance numeric default 0 not null
);
-- #
-- rollback drop table public.account;
