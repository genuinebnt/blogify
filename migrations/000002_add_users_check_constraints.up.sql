ALTER TABLE users
ADD CONSTRAINT users_password_check CHECK (LENGTH(password) >= 8);

