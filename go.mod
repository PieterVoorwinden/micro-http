module github.com/PieterVoorwinden/micro-http

go 1.13

// Pin to commit https://github.com/micro/go-micro/commit/7cad77bfc060e03bd6d0232424790a427f69297e
replace github.com/micro/go-micro/v2 => github.com/micro/go-micro/v2 v2.2.1-0.20200302161726-7cad77bfc060

require (
	github.com/golang/protobuf v1.3.4
	github.com/micro/go-micro/v2 v2.0.0-00010101000000-000000000000
)
