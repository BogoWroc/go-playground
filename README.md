How to Write Go Code 
====================
	https://golang.org/doc/code.html

Tips & Tricks
=============
	https://peter.bourgon.org/go-best-practices-2016/#repository-structure
	https://medium.com/@benbjohnson/standard-package-layout-7cdbc8391fc1
	https://yourbasic.org/golang/
	https://www.ardanlabs.com/all-posts/
	https://blog.golang.org/context
	
	https://blog.rubylearning.com/best-practices-for-a-new-go-developer-8660384302fc
	https://github.com/go-martini/martini
	
	http://devs.cloudimmunity.com/gotchas-and-common-mistakes-in-go-golang/

Effective go
============
	https://golang.org/doc/effective_go.html

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


A Golang BDD Testing framework
==============================

    https://github.com/onsi/ginkgo

# Installation 

    go get -u github.com/onsi/ginkgo/ginkgo  # installs the ginkgo CLI
    go get -u github.com/onsi/gomega/...     # fetches the matcher library

# IntelliJ configuration

![image](IntelliJConfiguration.png)
