server: 
	go run main.go

mock:
	mockgen -package mockdb -destination test/mock/store.go  github.com/fxmbx/go_practice_pr/repository UserRepository