# PingSQL
MySQL死活監視ツール。Dockerで活用できます。

### Using with Parameters.
```
pingsql -u user -p password -h hostname -c port(int) -n database_name
pingsql -u minako -p m1lli0n -h localhost -c 3306 -n satake
```

### Using with System Environment. (ShellScript Sample)
```
export DB_USER=user
export DB_PASS=password
export DB_HOST=hostname
export DB_PORT=3306
export DB_NAME=database_name

echo -n "Waiting MySQL."
until pingsql; do
    sleep 1
    echo -n "."
done
```
