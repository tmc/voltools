BEGIN;
ALTER TABLE zip_codes ADD CONSTRAINT zip_codes_pkey PRIMARY KEY (geoid10);
ALTER TABLE congressional_districts ADD CONSTRAINT congressional_districts_pkey PRIMARY KEY (geoid);
ALTER TABLE states ADD CONSTRAINT states_fkey PRIMARY KEY (statefp);
ALTER TABLE congressional_districts ADD CONSTRAINT congressional_districts_states_fkey FOREIGN KEY (statefp) REFERENCES states(statefp);
COMMIT;
