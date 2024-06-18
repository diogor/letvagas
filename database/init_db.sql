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
CREATE TYPE question_type AS ENUM (
    'computing',
    'language'
);
CREATE TYPE position_type AS ENUM (
    'temporary',
    'contract',
    'long_term'
);
CREATE TYPE allocation AS ENUM (
    'remote',
    'on_site',
    'hybrid'
);
CREATE TYPE level AS ENUM (
    'internship',
    'junior',
    'mid',
    'senior',
    'specialist'
);
