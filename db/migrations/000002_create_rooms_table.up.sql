CREATE TABLE rooms 
(
    id VARCHAR(64) NOT NULL,
    display_name VARCHAR(64) NOT NULL,
    description VARCHAR(255),
    created_by VARCHAR(32) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    
    PRIMARY KEY (id),
    FOREIGN KEY (created_by) REFERENCES users(id)
);