FROM mysql:5.7.24

ADD ./init/*.sql /docker-entrypoint-initdb.d/
#ADD ./init/csv/* /var/lib/mysql/sql-files/
ADD ./conf/my.cnf /etc/mysql/conf.d/my.cnf

RUN chmod 644 /etc/mysql/conf.d/my.cnf
