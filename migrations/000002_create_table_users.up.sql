CREATE TABLE IF NOT EXISTS todoapp.users(
    id           SERIAL       PRIMARY KEY,
    version      BIGINT       NOT NULL DEFAULT 1,
    full_name    VARCHAR(100) NOT NULL CHECK (char_length(full_name) BETWEEN 3 AND 100),
    phone_number VARCHAR(15)  NULL     CHECK (phone_number IS NULL
                                                OR (char_length(phone_number) BETWEEN 10 AND 15
                                                    AND phone_number ~'^\+[0-9]+$'))
);
