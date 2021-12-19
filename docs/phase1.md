# Set up CI and test frameworks

The first order of business for any project should be to set up 

1. An acceptance test
1. A CI pipeline to see the test fail.
1. minimal code to make the above do something even if the acceptance test still fails.

While this initial code copies liberally from Petr Jahodas article [Create Go Service The Easy Way](https://medium.com/swlh/create-go-service-the-easy-way-de827d7f07cf), it does not set up a system service.

# Steps

1. copy homepage.go and main.go files from <https://github.com/petrjahoda/medium_service
1. set up go modules: `go mod init github.com/cypherfox/tree-publish` and extract a current list of used modules: `go mod tidy`
1. set up GitHub CI pipeline to compile go code
