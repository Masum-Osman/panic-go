# run test on the current directory:
go test

#
go test -v

#
go test -v controllers/truckimage_test.go

#go benchmark testing:
#to run all benchmark function:
go test -bench=. 

#to run specific benchmark function:
go test -bench=Add

#to check how much of code is under test:
go test -coverprofile=coverage.out