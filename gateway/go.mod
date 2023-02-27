module service/gateway

go 1.18

require (
	github.com/golodash/galidator v1.3.0
	golang.org/x/text v0.7.0
	service/build v0.0.0-00010101000000-000000000000
	service/config v0.0.0-00010101000000-000000000000
	service/pkg v0.0.0-00010101000000-000000000000
)

require (
	github.com/alecthomas/units v0.0.0-20211218093645-b94a6e3cc137 // indirect
	github.com/dlclark/regexp2 v1.7.0 // indirect
	github.com/go-playground/locales v0.14.0 // indirect
	github.com/go-playground/universal-translator v0.18.0 // indirect
	github.com/go-playground/validator/v10 v10.11.1 // indirect
	github.com/golodash/godash v1.2.0 // indirect
	github.com/jinzhu/copier v0.3.5 // indirect
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/lestrrat-go/file-rotatelogs v2.4.0+incompatible // indirect
	github.com/lestrrat-go/strftime v1.0.6 // indirect
	github.com/nicksnyder/go-i18n/v2 v2.2.1 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/sirupsen/logrus v1.9.0 // indirect
	github.com/xhit/go-str2duration/v2 v2.1.0 // indirect
	golang.org/x/crypto v0.6.0 // indirect
	golang.org/x/sys v0.5.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace service/pkg => ../pkg

replace service/build => ../build

replace service/config => ../config
