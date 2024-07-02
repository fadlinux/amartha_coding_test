CREATE TABLE IF NOT EXISTS public.customer
(
    cif_id serial NOT NULL,
    fullname "char",
    ktp_no "char",
    address text COLLATE pg_catalog."default",
    create_date date,
    update_date date,
    CONSTRAINT customer_pkey PRIMARY KEY (cif_id)
);


ALTER TABLE IF EXISTS public.customer
    OWNER to postgres;

ALTER TABLE public.customer
ADD COLUMN new_cif_id serial;

ALTER TABLE public.customer
DROP CONSTRAINT customer_pkey;

ALTER TABLE public.customer
DROP COLUMN cif_id;

ALTER TABLE public.customer
RENAME COLUMN new_cif_id TO cif_id;

ALTER TABLE public.customer
ADD CONSTRAINT customer_pkey PRIMARY KEY (cif_id);



CREATE TABLE IF NOT EXISTS public.loan 
(
    loan_id serial NOT NULL,
    cif_id integer,
    total_amount integer,
    delinquent integer,
    create_time date,
    total_weeks integer,
    interest_rate integer,
    status "char",
    weekly_payment integer
);


ALTER TABLE ONLY public.loan 
ADD CONSTRAINT loan_pkey PRIMARY KEY (loan_id);


CREATE TABLE IF NOT EXISTS public.payment
(
    payment_id serial NOT NULL,
    loan_id "char",
    payment_date date,
    amount integer
);


ALTER TABLE public.payment
ADD CONSTRAINT payment_pkey PRIMARY KEY (payment_id);
