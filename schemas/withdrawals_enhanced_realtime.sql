CREATE TABLE IF NOT EXISTS beacon_mainnet.withdrawals_enhanced_realtime
(
    slot_number     BIGINT       NOT NULL,
    index           INTEGER      NOT NULL,
    validator_index BIGINT       NOT NULL,
    index_position  INTEGER      NOT NULL,
    address         VARCHAR(100),
    amount          NUMERIC(38, 0),
    block_time      TIMESTAMP    NOT NULL,
    block_number    BIGINT       NOT NULL,
    block_hash      VARCHAR(100) NOT NULL,
    block_date      DATE         NOT NULL,
    PRIMARY KEY (index)
);

CREATE INDEX ON beacon_mainnet.withdrawals_enhanced_realtime (block_date);
CREATE INDEX ON beacon_mainnet.withdrawals_enhanced_realtime (validator_index);


CREATE TABLE IF NOT EXISTS beacon_holesky.withdrawals_enhanced_realtime
(
    slot_number     BIGINT       NOT NULL,
    index           INTEGER      NOT NULL,
    validator_index BIGINT       NOT NULL,
    index_position  INTEGER      NOT NULL,
    address         VARCHAR(100),
    amount          NUMERIC(38, 0),
    block_time      TIMESTAMP    NOT NULL,
    block_number    BIGINT       NOT NULL,
    block_hash      VARCHAR(100) NOT NULL,
    block_date      DATE         NOT NULL,
    PRIMARY KEY (index)
);

CREATE INDEX ON beacon_holesky.withdrawals_enhanced_realtime (block_date);
CREATE INDEX ON beacon_holesky.withdrawals_enhanced_realtime (validator_index);
