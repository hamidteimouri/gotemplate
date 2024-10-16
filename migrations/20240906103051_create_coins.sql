-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS coins
(
    id                     SERIAL PRIMARY KEY,
    symbol                 VARCHAR(128)  NULL DEFAULT NULL,
    name                   VARCHAR(256)  NULL DEFAULT NULL,
    name_locale            VARCHAR(256)  NULL DEFAULT NULL,
    coin_type              VARCHAR(128)  NULL DEFAULT FALSE, -- crypto , toman
    is_active              BOOLEAN            DEFAULT FALSE, -- if false , does not return for users
    is_suspended           BOOLEAN            DEFAULT FALSE, -- if false , does not accept anything
    need_kyc               BOOLEAN            DEFAULT FALSE,
    is_coin                BOOLEAN            DEFAULT FALSE,
    is_fiat                BOOLEAN            DEFAULT FALSE,
    is_popular             BOOLEAN            DEFAULT FALSE,
    is_new                 BOOLEAN            DEFAULT FALSE,
    high_risk              BOOLEAN            DEFAULT FALSE,
    can_sell               BOOLEAN            DEFAULT FALSE,
    can_buy                BOOLEAN            DEFAULT FALSE,
    can_trade              BOOLEAN            DEFAULT FALSE,
    can_deposit            BOOLEAN            DEFAULT FALSE,
    can_withdraw           BOOLEAN            DEFAULT FALSE,
    can_stake              BOOLEAN            DEFAULT FALSE,
    can_loan               BOOLEAN            DEFAULT FALSE,
    precision              INTEGER       NULL DEFAULT 1,
    position               INTEGER       NULL DEFAULT 1000,
    suspension_description VARCHAR(1024) NULL DEFAULT NULL,
    website                VARCHAR(1024) NULL DEFAULT NULL,
    created_at             TIMESTAMP     NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at             TIMESTAMP     NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_bc_symbol on coins (symbol);
CREATE INDEX IF NOT EXISTS idx_bc_isActive on coins (is_active);
CREATE INDEX IF NOT EXISTS idx_bc_coinType on coins (coin_type);
CREATE INDEX IF NOT EXISTS idx_bc_canTrade on coins (can_trade);
CREATE INDEX IF NOT EXISTS idx_bc_canDeposit on coins (can_deposit);
CREATE INDEX IF NOT EXISTS idx_bc_canWithdraw on coins (can_withdraw);
CREATE INDEX IF NOT EXISTS idx_bc_position on coins (position ASC);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS idx_bc_symbol;
DROP INDEX IF EXISTS idx_bc_isActive;
DROP INDEX IF EXISTS idx_bc_coinType;
DROP INDEX IF EXISTS idx_bc_canTrade;
DROP INDEX IF EXISTS idx_bc_canDeposit;
DROP INDEX IF EXISTS idx_bc_canWithdraw;
DROP INDEX IF EXISTS idx_bc_position;


DROP TABLE IF EXISTS coins;
-- +goose StatementEnd
