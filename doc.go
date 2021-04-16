//Steps
//gmod go mod init github.com/basant93/gmod
//go: creating new go.mod: module github.com/basant93/gmod
//➜  gmod go build .
//➜  gmod ./testgomod
package main

//Indirect tag means go package has been added to modules but not getting used in code
//if we use it Indirect is removed

//to get the dependencies use go get
//gmod  go get -u github.com/gorilla/mux
//go: downloading github.com/gorilla/mux v1.8.0
//go: github.com/gorilla/mux upgrade => v1.8.0

//to check the dependency module is relying on
//➜  gmod go list
//github.com/basant93/gmod
//➜  gmod go list all
//➜  gmod go list -m all
//github.com/basant93/gmod
//github.com/gorilla/mux v1.8.0

//verifting the go modules
//➜  gmod go mod verify
//all modules verified

//go mod tidy
//remove dependencies that are no longer required.
//creates a dependency graph for the entire modules and compares it with modules in go mod file
//and remove the no longer required dependency

//module version rules

//major, minor and patch
//rsc.io/quote
//quote2 rsc.io/quote/v2
//provide package alias, v2 is a syntactic folder

//for unversioned commits, we do use systematic versioning
//prerelease identifier are timestamp and commit hash
require golang.org/x/tools v0.0.0-2021041511121223-9023213213

//to download the specific module, go get -u golang.org/x/tools@v1.2.3
//use module queries

//subcommands
//go mod why github.com/gurilla/mux
//go mod graph
//go mod vendor
//go run -mod=vendor
//go mod edit -module "module path"
//go build -mod=readonly


