-- USERS TABLE
CREATE TABLE public.users (
    id integer NOT NULL,
    name text NOT NULL,
    age integer NOT NULL,
    sex text NOT NULL
);

ALTER TABLE public.users
ADD CONSTRAINT users_pkey PRIMARY KEY (id);


-- ADMIN TABLE
CREATE TABLE public.admin (
    login text NOT NULL,
    pass text NOT NULL
);

ALTER TABLE public.admin
ADD CONSTRAINT admin_pkey PRIMARY KEY (login);