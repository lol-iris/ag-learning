FROM postgres:14
COPY setup.sql /docker-entrypoint-initdb.d
