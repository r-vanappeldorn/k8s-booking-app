
-- Referentiedata
CREATE TABLE continents (
  id        TINYINT UNSIGNED PRIMARY KEY,
  code      CHAR(2) NOT NULL UNIQUE,    -- "EU", "NA", ...
  name      VARCHAR(64) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE countries (
  code            CHAR(2) PRIMARY KEY,  -- ISO 3166-1 alpha-2 (NL, DE, ...)
  name            VARCHAR(128) NOT NULL,
  continent_id    TINYINT UNSIGNED NOT NULL,
  CONSTRAINT fk_countries_continent
    FOREIGN KEY (continent_id) REFERENCES continents(id)
      ON UPDATE RESTRICT ON DELETE RESTRICT
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE locations (
  id            BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
  name          VARCHAR(128) NOT NULL,           -- bv. "Amsterdam"
  country_code  CHAR(2) NOT NULL,
  admin_area    VARCHAR(128) NULL,               -- provincie/regio
  latitude      DECIMAL(9,6) NULL,
  longitude     DECIMAL(9,6) NULL,
  CONSTRAINT fk_locations_country
    FOREIGN KEY (country_code) REFERENCES countries(code)
      ON UPDATE RESTRICT ON DELETE RESTRICT
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- Domein
CREATE TABLE trips (
  id                BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
  slug              VARCHAR(128) NOT NULL UNIQUE,         -- voor prettige URLs
  title             VARCHAR(200) NOT NULL,
  description       TEXT NULL,
  location_id       BIGINT UNSIGNED NOT NULL,
  starts_at         DATETIME NOT NULL,
  ends_at           DATETIME NOT NULL,
  capacity          INT UNSIGNED NOT NULL,                -- totale plekken
  base_price_cents  INT NOT NULL,                         -- prijs per seat
  currency          CHAR(3) NOT NULL,                     -- "EUR", "USD"
  status            ENUM('draft','published','archived') NOT NULL DEFAULT 'draft',
  created_at        DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at        DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  CONSTRAINT fk_trips_location
    FOREIGN KEY (location_id) REFERENCES locations(id)
      ON UPDATE RESTRICT ON DELETE RESTRICT,
  CONSTRAINT chk_dates CHECK (ends_at > starts_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE bookings (
  id                 BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
  trip_id            BIGINT UNSIGNED NOT NULL,
  account_id         VARCHAR(64) NOT NULL,                -- extern (UUID/ULID)
  status             ENUM('pending','confirmed','cancelled','refunded')
                      NOT NULL DEFAULT 'pending',
  seats              INT UNSIGNED NOT NULL DEFAULT 1,
  price_cents        INT NOT NULL,                        -- per booking (incl. seats, kortingen)
  currency           CHAR(3) NOT NULL,
  payment_ref        VARCHAR(100) NULL,                   -- PSP id
  idempotency_key    VARCHAR(100) NULL,                   -- voor duplicated submits
  created_at         DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at         DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

  CONSTRAINT fk_bookings_trip
    FOREIGN KEY (trip_id) REFERENCES trips(id)
      ON UPDATE RESTRICT ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- Indexen
CREATE INDEX idx_locations_country ON locations(country_code);
CREATE INDEX idx_trips_loc_time ON trips(location_id, starts_at);
CREATE INDEX idx_trips_status ON trips(status);
CREATE INDEX idx_bookings_trip ON bookings(trip_id);
CREATE INDEX idx_bookings_account ON bookings(account_id);
CREATE UNIQUE INDEX uq_bookings_idem ON bookings(idempotency_key);
