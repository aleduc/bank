-- liquibase formatted sql

-- changeset aleduc:1 stripComments:false runOnChange:false endDelimiter:#
create table transaction
(
  id bigserial primary key,
  id_account bigint not null,
  d_time timestamp,
  d_actual_time timestamp,
  n_currency_code integer,
  n_amount numeric,
  n_transaction_type integer,
  n_transaction_status integer
);
-- #
-- rollback drop table public.transaction;
