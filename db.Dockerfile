FROM mysql:8.0.23

COPY ./config/*.sql /docker-entrypoint-initdb.d/