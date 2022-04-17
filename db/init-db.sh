#!/bin/bash
set -e

psql -v ON_ERROR_STOP=1 -U postgres <<-EOSQL
CREATE USER docker WITH ENCRYPTED PASSWORD 'qwerty';
CREATE DATABASE shmot_shprot_db WITH OWNER docker;
\c shmot_shprot_db docker
\i /scripts/schema.sql
\i /scripts/init.sql
EOSQL


