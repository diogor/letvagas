CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TYPE user_role AS ENUM (
    'admin',
    'client',
    'applicant'
);
