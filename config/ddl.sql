
#Master Kurir
CREATE TABLE couriers (
    user_id VARCHAR(255) PRIMARY KEY,
    name VARCHAR(150) NOT NULL,
    email VARCHAR(150) UNIQUE,
    phone VARCHAR(20) NOT NULL,
    address TEXT NOT NULL,
    level VARCHAR(150) NOT NULL,
    status VARCHAR(20) NOT NULL DEFAULT 'active',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


#Master Level kurir
CREATE TABLE levels (
	level_id VARCHAR(255) PRIMARY KEY,
	name VARCHAR(150) NOT NULL,
	description TEXT
);

#Master Status Kurir
CREATE TABLE status (
	type VARCHAR(150) PRIMARY KEY
);

#korelasi
ALTER TABLE couriers
ADD CONSTRAINT fk_level_couriers
FOREIGN KEY (level)
REFERENCES levels(level_id);

ALTER TABLE couriers
ADD CONSTRAINT fk_status_couriers
FOREIGN KEY (status)
REFERENCES status(type);
