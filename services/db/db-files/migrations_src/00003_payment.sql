-- +goose Up
-- SQL in this section is executed when the migration is applied.

CREATE TABLE key_value_store (
  key varchar(255) NOT NULL PRIMARY KEY,
  str_value varchar(255) NOT NULL,
  int_value bigint NOT NULL
);

INSERT INTO key_value_store (key, str_value, int_value) VALUES ('eth_address_index', '', 0);
INSERT INTO key_value_store (key, str_value, int_value) VALUES ('eth_last_block', '0', 0);

INSERT INTO key_value_store (key, str_value, int_value) VALUES ('btc_address_index', '', 0);
INSERT INTO key_value_store (key, str_value, int_value) VALUES ('btc_last_block', '0', 0);

INSERT INTO key_value_store (key, str_value, int_value) VALUES ('xlm_last_ledger_id', '0', 0);

CREATE TYPE order_status AS ENUM ('waiting_for_payment', 'payment_received', 'waiting_user_transaction', 'payment_error', 'finished', 'error', 'under_pay', 'over_pay', 'no_coins_left', 'phase_expired');

CREATE TABLE user_order (
  id SERIAL PRIMARY KEY NOT null,
  user_id integer not null REFERENCES user_profile(id),

  /* ========== */
  /* order data */
  ico_phase_id int not null REFERENCES ico_phase(id),
  order_status order_status not null,
  token_amount bigint not null,
  /* the users public key for the payment */
  /* coins will be transfered here on success */
  stellar_user_public_key varchar(56) NOT NULL,

  exchange_currency_id int not null REFERENCES exchange_currency(id),
  exchange_currency_denomination_amount varchar(64) not null, /* denomination for selected currency */

  payment_network payment_network not null, /* this is just used as information when filtering the orders */

  /* ========== */
  /* chain data */
  address_index bigint NOT NULL, /* used in eth and btc for generating the address */
  /* bitcoin 34 characters */
  /* ethereum 42 characters */
  /* stellar 56 characters */
  payment_address varchar(56) NOT NULL, /* public key in the target network, based on payment_network */
  payment_seed varchar(56) NOT NULL, /* this is either the seed on stellar, or the privatekey on other crypto */

  stellar_transaction_id text not null, /* this is the coin payment tx in the stellar network */
  processed_transaction_id int null, /* FK to the processed transactions */

  payment_qr_image bytea null, /* qr-image for the payment transaction */

  payment_usage varchar(255) not null, /* used only for fiat and stellar payment payments. For stellar, the data will be read from the momo */

  /* this field is used to save any error message that happened during the client payment */
  payment_error_message text not null,

  /* ============== */
  /* default fields */
  created_at timestamp with time zone NOT NULL default current_timestamp,
  updated_at timestamp with time zone NOT NULL default current_timestamp,
  updated_by character varying not null,

  CONSTRAINT valid_address_index CHECK (address_index >= 0)
);
create index idx_user_order_user_profile on user_order(user_id);
create unique index idx_user_order_ix1 on user_order(exchange_currency_id, payment_address) where payment_network <> 'stellar' and payment_network <> 'fiat';
create index idx_user_order_ix2 on user_order(updated_at); /* need this for fast filtering */
create index idx_user_order_ix3 on user_order(order_status);
create index idx_user_order_ix4 on user_order(stellar_user_public_key);
create index idx_user_order_ix5 on user_order(ico_phase_id);

CREATE TYPE transaction_status AS ENUM ('new', 'processed');
CREATE TABLE processed_transaction (
  id SERIAL PRIMARY KEY NOT null,
  status transaction_status not null,

  payment_network payment_network NOT NULL,

  /* Ethereum: "0x"+hash (so 64+2) */
  transaction_id varchar(66) NOT NULL,
  /* bitcoin 34 characters */
  /* ethereum 42 characters */
   /* stellar 56 characters */
  refund_tx_id text not null, /* refund payment hash/id from the PaymentNetwork */

  receiving_address varchar(56) NOT NULL,
  payment_network_amount_denomination varchar(64) not null, /* max one billion */

  order_id integer not null REFERENCES user_order(id),

  created_at timestamp with time zone NOT NULL default current_timestamp,
  updated_at timestamp with time zone NOT NULL default current_timestamp
);
create unique index idx_processed_transaction_ix1 on processed_transaction(order_id);
create unique index idx_processed_transaction_ix2 on processed_transaction(payment_network, transaction_id);
alter table user_order add FOREIGN KEY(processed_transaction_id) REFERENCES processed_transaction(id);

CREATE TABLE multiple_transaction (
  id SERIAL PRIMARY KEY NOT null,
  payment_network payment_network NOT NULL,
  transaction_id varchar(66) NOT NULL,
  refund_tx_id text not null,
  receiving_address varchar(56) NOT NULL,
  payment_network_amount_denom varchar(64) not null,
  order_id integer not null REFERENCES user_order(id),
  created_at timestamp with time zone NOT NULL default current_timestamp,
  updated_at timestamp with time zone NOT NULL default current_timestamp
);
create unique index idx_multiple_transaction_ix1 on multiple_transaction(order_id, transaction_id);
create index idx_multiple_transaction_ix2 on multiple_transaction(payment_network, transaction_id);
create index idx_multiple_transaction_ix3 on multiple_transaction(payment_network, receiving_address);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
drop table IF EXISTS key_value_store;
ALTER TABLE user_order DROP CONSTRAINT IF EXISTS "user_order_processed_transaction_id_fkey";
ALTER TABLE processed_transaction DROP CONSTRAINT IF EXISTS "processed_transaction_order_id_fkey";
ALTER TABLE multiple_transaction DROP CONSTRAINT IF EXISTS "multiple_transaction_order_id_fkey";
drop table if exists processed_transaction;
drop table if exists multiple_transaction;
drop table IF EXISTS user_order;
drop type if exists order_status;
drop type if exists transaction_status;
drop type if exists denomination_amount;
