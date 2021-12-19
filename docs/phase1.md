# Set up CI and test frameworks

The first order of business for any project should be to set up 

1. An acceptance test
1.  A CI pipeline to see the test fail.
1.  minimal code to make the above do something even if the acceptance test still fails.

While this initial code copies liberally from Petr Jahodas article [Create Go Service The Easy Way](https://medium.com/swlh/create-go-service-the-easy-way-de827d7f07cf), it does not set up a system service.

