ALTER TABLE users
ADD CONSTRAINT users_password_check CHECK (LEN(password) >= 8);
