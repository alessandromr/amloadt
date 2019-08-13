module github.com/alessandromr/amloadt

require (
	github.com/alessandromr/amloadt/sharedq v0.0.1
	github.com/alessandromr/amloadt/sualyzer v0.0.1
	github.com/gorilla/mux v1.7.3
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
)

replace (
	github.com/alessandromr/amloadt/sharedq => ./sharedq
	github.com/alessandromr/amloadt/sualyzer => ./ualyzer
)
