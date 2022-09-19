module github.com/laurentino14/teste

go 1.19

replace github.com/laurentino14/teste/teste => ./teste

replace github.com/laurentino14/teste/db => ./db

replace github.com/laurentino14/teste/http => ./http

replace github.com/laurentino14/teste/http/routes => ./http/routes

require (
	github.com/iancoleman/strcase v0.0.0-20190422225806-e506e3ef7365
	github.com/joho/godotenv v1.4.0
	github.com/labstack/echo/v4 v4.9.0
	github.com/prisma/prisma-client-go v0.16.2
	github.com/shopspring/decimal v1.3.1
	github.com/takuoki/gocase v1.0.0
)

require (
	github.com/labstack/gommon v0.3.1 // indirect
	github.com/mattn/go-colorable v0.1.11 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasttemplate v1.2.1 // indirect
	golang.org/x/crypto v0.0.0-20210817164053-32db794688a5 // indirect
	golang.org/x/net v0.0.0-20211015210444-4f30a5c0130f // indirect
	golang.org/x/sys v0.0.0-20211103235746-7861aae1554b // indirect
	golang.org/x/text v0.3.7 // indirect
)
