server: 
	go run main.go

mock:
	mockgen --build_flags=--mod=mod -package mockdb -destination test/mock/userRepo.mock.go  github.com/fxmbx/go_practice_pr/repository UserRepository	

test:
	go test -v -cover ./...
	
# mockgen -source=/go_practice_pr/repository/user.repository.go  -destination test/mock/store.go -package mockdb