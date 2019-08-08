## Framework
github.com/ziyeziye/framework
是一个Golang+Gin+Gorm的脚手架
 
## Framework-gen
github.com/ziyeziye/framework-gen
是一个可以通过数据库生成对应framework的models,struct以及相应的restful api的工具。

## Supported Databases
- MariaDB
- MySQL
- PostgreSQL
- Microsoft SQL Server

## Usage

```BASH
--connstr database connection string
--database Database to for connection
--table Table to build struct from
--prefix Table prefix
--package name to set for package

go run main.go --connstr "root:pass@tcp(127.0.0.1:3306)/dbname?&parseTime=True" --package github.com/ziyeziye/framework --prefix pf_ --json --gorm --g
uregu --rest

framework-gen --connstr "root:pass@tcp(127.0.0.1:3306)/dbname?&parseTime=True" --package github.com/ziyeziye/framework --prefix pf_ --json --gorm --g
uregu --rest
```
#### Supported Datatypes

Currently only a limited number of datatypes are supported. Initial support includes:
-  tinyint (sql.NullInt64 or null.Int)
-  int      (sql.NullInt64 or null.Int)
-  smallint      (sql.NullInt64 or null.Int)
-  mediumint      (sql.NullInt64 or null.Int)
-  bigint (sql.NullInt64 or null.Int)
-  decimal (sql.NullFloat64 or null.Float)
-  float (sql.NullFloat64 or null.Float)
-  double (sql.NullFloat64 or null.Float)
-  datetime (null.Time)
-  time  (null.Time)
-  date (null.Time)
-  timestamp (null.Time)
-  var (sql.String or null.String)
-  enum (sql.String or null.String)
-  varchar (sql.String or null.String)
-  longtext (sql.String or null.String)
-  mediumtext (sql.String or null.String)
-  text (sql.String or null.String)
-  tinytext (sql.String or null.String)
-  binary
-  blob
-  longblob
-  mediumblob
-  varbinary
