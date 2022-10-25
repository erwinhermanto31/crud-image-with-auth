FROM mysql:5.7.22

COPY ./custom.cnf /etc/mysqld/conf.d/custom.cnf
