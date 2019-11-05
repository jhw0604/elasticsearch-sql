# A Pure Go Elasticsearch driver for Go's database/sql package #
It is writen **Pure Go**
run via elasticsearch rest api /_xpack/sql
So... only support **SELECT** at this time

## Install ##
Install with 
```bash
go get github.com/jhw0604/elasticsearch-sql
```


## Connection Parameters and DSN ##
```http[s]://[<encoding url base64 id:password>@]address:port```

ID and Password connect with ":" and encoding base64 URL.

```go
import (
	"database/sql"
	_ "github.com/jhw0604/elasticsearch-sql"
)


db, err := sql.Open("elasticsearch", "http://localhost:9200")
```

```go
import (
	"database/sql"
	_ "github.com/jhw0604/elasticsearch-sql"
)


db, err := sql.Open("elasticsearch", "http://ZWxhc3RpYzpwYXNzd29yZA==@localhost:9200")
```

