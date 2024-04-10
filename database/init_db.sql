CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TYPE user_role AS ENUM (
    'admin',
    'client',
    'applicant'
);
CREATE TYPE education_type AS ENUM (
    'fundamental',
    'medio',
    'superior',
    'pos',
    'mba',
    'outros'
);
