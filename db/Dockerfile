FROM postgres:alpine

COPY . /scripts 
COPY ./init-db.sh /docker-entrypoint-initdb.d/init-db.sh
WORKDIR /

ENV POSTGRES_PASSWORD=qwerty

RUN chmod u+x ./docker-entrypoint-initdb.d/init-db.sh

#docker build -t shmot-shprot-db:1.0 .
#docker run --name docker-experiment [-e POSTGRES_PASSWORD=qwerty] [-v ./db/:/scripts] -d --rm postgres:alpine
#docker exec -it $id /bin/bash

#socket /var/run/postgresql/.s.PGSQL.5432
