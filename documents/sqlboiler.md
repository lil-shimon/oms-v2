
# sqlboiler mysql

## error 1
```sh
Error: unable to initialize tables: unable to fetch table data: driver (/go/bin/sqlboiler-mysql) exited non-zero: exit status 
```

### driver install
```sh
go get -u -t github.com/volatiletech/sqlboiler/drivers/sqlboiler-mysql
```

### v4 driver
```sh
go get github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-mysql
```

### 解決方法
```
[mysql]
  dbname="oms"
  host="mysql_host" -> docker mysql container nameに変更（元々はlocalhost）
  port=3306
  user="admin"
  pass="admin"  
  sslmode="false"
```

## error 2
```
Error: unable to initialize tables: unable to fetch table data: something totally unexpected happened when running the binary driver /go/src/server/sqlboiler-mysql: fork/exec /go/src/server/sqlboiler-mysql: no such file or directory
```
再度インストールで治った