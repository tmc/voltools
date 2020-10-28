ALTER TABLE zip_codes DROP CONSTRAINT zip_codes_pkey;
ALTER TABLE congressional_districts DROP CONSTRAINT congressional_districts_pkey;
ALTER TABLE states DROP CONSTRAINT states_pkey;
ALTER TABLE congressional_districts DROP CONSTRAINT congressional_districts_states_fkey;
