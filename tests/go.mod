module github.com/bbernhard/imagemonkey-core-tests

go 1.12

require (
	github.com/bbernhard/imagemonkey-core v0.0.0-00010101000000-000000000000
	github.com/jackc/pgx/v4 v4.1.2
	github.com/sirupsen/logrus v1.4.2
	gopkg.in/resty.v1 v1.12.0
)

replace github.com/bbernhard/imagemonkey-core => ../src/
