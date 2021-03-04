module github.com/rexwsd/blog

go 1.15

replace (
	github.com/rexwsd/blog/api => ../api
	github.com/rexwsd/blog/bootstrap => ../bootstarp
	github.com/rexwsd/blog/conf => ../pkg/conf
	github.com/rexwsd/blog/config => ../config
	github.com/rexwsd/blog/middleware => ../middleware
	github.com/rexwsd/blog/models => ../models
	github.com/rexwsd/blog/pkg/e => ../pkg/e
	github.com/rexwsd/blog/pkg/setting => ../pkg/setting
	github.com/rexwsd/blog/pkg/util => ../pkg/util
	github.com/rexwsd/blog/routers => ../routers
	github.com/rexwsd/blog/tools => ../tools
)

require (
	github.com/gin-contrib/cors v1.3.1
	github.com/gin-gonic/gin v1.6.3
	github.com/go-ini/ini v1.62.0
	github.com/go-playground/validator/v10 v10.4.1 // indirect
	github.com/go-sql-driver/mysql v1.5.0
	github.com/gohouse/golib v0.0.0-20201013092821-ca87ab19b554 // indirect
	github.com/gohouse/gorose/v2 v2.1.10
	github.com/golang/protobuf v1.4.3 // indirect
	github.com/joho/godotenv v1.3.0
	github.com/json-iterator/go v1.1.10 // indirect
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/lestrrat-go/file-rotatelogs v2.4.0+incompatible
	github.com/lestrrat-go/strftime v1.0.4 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.1 // indirect
	github.com/ugorji/go v1.2.4 // indirect
	github.com/unknwon/com v1.0.1 // indirect
	go.uber.org/multierr v1.6.0 // indirect
	go.uber.org/zap v1.16.0
	golang.org/x/crypto v0.0.0-20210220033148-5ea612d1eb83 // indirect
	golang.org/x/sys v0.0.0-20210223212115-eede4237b368 // indirect
	google.golang.org/protobuf v1.25.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)
