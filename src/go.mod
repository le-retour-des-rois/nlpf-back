module test/main

go 1.17

replace test/transaction v1.0.0-unpublished => ./transaction

replace test/realEstate v1.0.0-unpublished => ./realEstateProject

require (
	github.com/gorilla/mux v1.8.0
	go.mongodb.org/mongo-driver v1.7.3
)

require (
	github.com/fogleman/gg v1.3.0 // indirect
	github.com/fogleman/primitive v0.0.0-20200504002142-0373c216458b // indirect
	github.com/go-stack/stack v1.8.0 // indirect
	github.com/golang/freetype v0.0.0-20170609003504-e2365dfdc4a0 // indirect
	github.com/golang/snappy v0.0.1 // indirect
	github.com/klauspost/compress v1.13.6 // indirect
	github.com/nfnt/resize v0.0.0-20180221191011-83c6a9932646 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/xdg-go/pbkdf2 v1.0.0 // indirect
	github.com/xdg-go/scram v1.0.2 // indirect
	github.com/xdg-go/stringprep v1.0.2 // indirect
	github.com/youmark/pkcs8 v0.0.0-20181117223130-1be2e3e5546d // indirect
	golang.org/x/crypto v0.0.0-20200302210943-78000ba7a073 // indirect
	golang.org/x/image v0.0.0-20210628002857-a66eb6448b8d // indirect
	golang.org/x/sync v0.0.0-20190911185100-cd5d95a43a6e // indirect
	golang.org/x/text v0.3.6 // indirect
	test/domain v1.0.0-unpublished // indirect
	test/realEstate v1.0.0-unpublished // indirect
	test/transaction v1.0.0-unpublished // indirect
)

replace test/domain v1.0.0-unpublished => ./domain
