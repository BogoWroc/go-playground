How to Write Go Code 
====================
	https://golang.org/doc/code.html

All in one
==========
https://dev.to/digitalocean/how-to-code-in-go-32p0

Tips & Tricks
=============
	https://peter.bourgon.org/go-best-practices-2016/#repository-structure
	https://medium.com/@benbjohnson/standard-package-layout-7cdbc8391fc1
	
	https://yourbasic.org/golang/
		https://yourbasic.org/golang/conversions/

	https://www.ardanlabs.com/all-posts/
	https://blog.golang.org/context
	
	https://blog.rubylearning.com/best-practices-for-a-new-go-developer-8660384302fc
	https://github.com/go-martini/martini
	
	http://devs.cloudimmunity.com/gotchas-and-common-mistakes-in-go-golang/
	
	https://dmitri.shuralyov.com/idiomatic-go

Effective go
============
	https://golang.org/doc/effective_go.html

Go Patterns
===========
    http://tmrts.com/go-patterns/
    
FAQ
===
	https://golang.org/doc/faq

PROJECT TEMPLATE
================
    https://github.com/ardanlabs/service
    
Packaging/Modules
=================
    https://github.com/ardanlabs/gotraining/tree/master/topics/go/design/packaging
    https://github.com/golang/go/wiki/Modules#how-do-i-use-vendoring-with-modules-is-vendoring-going-away
    https://www.ardanlabs.com/blog/2019/10/modules-01-why-and-what.html

Testing and Mocking
===================
	https://blog.codecentric.de/2019/07/gomock-vs-testify/
	https://blog.alexellis.io/golang-writing-unit-tests/

	BDD testing library
	-------------------
	https://github.com/onsi/ginkgo
	
	Commands
	--------
	Execute all tests: go test ./...
	Execute tests in module: go test
	Execute benchmark tests in module: go test -bench .
	Execute code coverage in module: go test -cover
	Generate reports form coverage:
	    go test -coverprofile c.out
	    
	    or
	     
	    go tool cover -html=c.out 
	    <next> 
	    go tool cover -html=c.out -o c.html


Error handling
==============
	https://blog.golang.org/error-handling-and-go
	https://blog.golang.org/errors-are-values
	https://blog.golang.org/defer-panic-and-recover

Examples
========
	https://gobyexample.com/
	https://github.com/GoesToEleven/go-programming

Golang Web Server, and unit tests 
=======================================================================
    https://medium.com/myntra-engineering/my-journey-with-golang-web-services-4d922a8c9897
    https://github.com/adnaan/learngo


Screen recording
================
    https://obsproject.com/

Interesting libraries
======================
    Linter
    ------
    go get -u golang.org/x/lint/golint
    
    Local document generator
    ------------------------
    go get golang.org/x/tools/cmd/godoc
 
    Starting local documentation server   
    godoc -http=:8080
    
    Go Env
    ------
    https://github.com/syndbg/goenv/blob/master/INSTALL.md


A Golang BDD Testing framework
==============================

    https://github.com/onsi/ginkgo

# Installation 

    go get -u github.com/onsi/ginkgo/ginkgo  # installs the ginkgo CLI
    go get -u github.com/onsi/gomega/...     # fetches the matcher library

# IntelliJ configuration

![image](IntelliJConfiguration.png)
