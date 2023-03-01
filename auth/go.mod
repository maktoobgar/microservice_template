module service/auth

go 1.20

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/rubenv/sql-migrate v1.3.1
	golang.org/x/crypto v0.5.0
	golang.org/x/text v0.7.0
	google.golang.org/grpc v1.53.0
	google.golang.org/protobuf v1.28.1
	service/build v0.0.0-00010101000000-000000000000
	service/config v0.0.0-00010101000000-000000000000
	service/pkg v0.0.0-00010101000000-000000000000
)

require (
	github.com/alecthomas/units v0.0.0-20211218093645-b94a6e3cc137 // indirect
	github.com/denisenkom/go-mssqldb v0.12.3 // indirect
	github.com/go-gorp/gorp/v3 v3.0.5 // indirect
	github.com/go-sql-driver/mysql v1.7.0 // indirect
	github.com/golang-sql/civil v0.0.0-20190719163853-cb61b32ac6fe // indirect
	github.com/golang-sql/sqlexp v0.1.0 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/lestrrat-go/file-rotatelogs v2.4.0+incompatible // indirect
	github.com/lestrrat-go/strftime v1.0.6 // indirect
	github.com/lib/pq v1.10.7 // indirect
	github.com/mattn/go-sqlite3 v1.14.16 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/sirupsen/logrus v1.9.0 // indirect
	github.com/xhit/go-str2duration/v2 v2.1.0 // indirect
	golang.org/x/net v0.5.0 // indirect
	golang.org/x/sys v0.4.0 // indirect
	google.golang.org/genproto v0.0.0-20230110181048-76db0878b65f // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace service/pkg => ../pkg

replace service/build => ../build

replace service/config => ../config
