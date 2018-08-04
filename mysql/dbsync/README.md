# Dbsync server

HTTP server which provides access to a MySQL database in the form of a dump file (tarred and zipped). The dump is created via `select into outfile` statements and contains one file per table, and a `statements.sql` for recreating the table structure.  This was created for dumping/moving InnoDB tables, and while it will no doubt work with other table types there are probably faster ways to get the data.

*No authentication* is provided by the server, so be careful where you put this (bind to a private IP address or localhost).

This code is probably buggy and untested.

# Installation

No service start-up scripts are provided, so the server needs to be started manually (but look into crontab's `@reboot`).  It should also have access to the same file system as MySQL.

```
$ go get github.com/allermedia/dbsync
$ go build -o dbsync compress.go dump.go server.go
$ mv dbsync /usr/local/bin
```

Run as unprivildged user, and edit the server's required environment variables (see `varexport.sh`).  There is currently *hacked in* support for dumping from two different database servers - for example if a particular database resides on a different server.  If still additional servers are needed, this needs to be added - and perhaps cleaned up :)

If you use a second server, then you also need to specify which databases this server should serve.  Everything else will be served by the first server.

```
$ vim varexport.sh
$ source varexport.sh
```

## Start the server

```
$ dbsync &> /tmp/dbsync.log &
```

# Accessing the service

```
GET /database/:databasename
```

## Arguments (via query string)

* Restrict each table to this amount of rows.

```
limit=NUMBER
```

* Ignore tables (relevant create statements will still be generated, but no data sent)

```
ignore=cache*,voting*
```

# Loading the dump into MySQL

This requires you to write your own program/bash script.  Unzip and tar the file; process the `statements.sql` file first, then you can process all the other files which are named according to the table name.  You can use this SQL statement to load the data back in:

```
LOAD DATA LOCAL INFILE 'table_name.txt' INTO TABLE table_name CHARACTER SET UTF8 FIELDS TERMINATED by ',' ENCLOSED BY '"'
```

You will probably also need these MySQL settings (as a `my.cfg` file):

```
[mysqld]
max_allowed_packet=2047000000
innodb_autoinc_lock_mode = 2
foreign_key_checks = 0
unique_checks = 0
tx_isolation='READ-UNCOMMITTED'
sql_log_bin = 0
innodb_lock_wait_timeout = 5000

[mysql]
local-infile=1
```

# TODO

* Provide server as Docker image
* Where are the tests
* Get rid of all the fatal errors in dump (need to send errors back to client)
