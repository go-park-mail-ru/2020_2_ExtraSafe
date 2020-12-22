FROM postgres:alpine

COPY ../../init/ /docker-entrypoint-initdb.d/
RUN chmod +x /docker-entrypoint-initdb.d/init.sh
