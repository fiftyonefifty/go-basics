# go-basics

# Obvious stuff.

```
mkdir hello-rest
cd hello-rest
go mod init hello-rest
```
This puts things in context, where we don't need the GOPATH.  You will notice a **go.mod** file inside the **./hello-rest/** folder.  

While still inside the **./hello-rest/** folder call the following;  

```
go get -u github.com/gorilla/mux
```
this will create a **go.sum** file that looks like this;  

```
github.com/gorilla/mux v1.7.3 h1:gnP5JzjVOuiZD07fKKToCAOjS0yOpj/qPETTXCCS6hw=
github.com/gorilla/mux v1.7.3/go.mod h1:1lud6UwP+6orDFRuTfBEV8e9/aOM/c4fVVCaMa2zaAs=

```

The next thing to know is that most your code will be in another file and the following shows you how to reference that;  
While still inside the **./hello-rest/** folder call the following;  
```
mkdir pkg
cd pkg
mkdir models
mkdir controllers
```
The final layout is located [here](src/hello-rest)  

While still inside the **./hello-rest/** folder call the following;  

```
go run ./main.go 
```
To see it work browse to the following url;
[http://localhost:8000/books](http://localhost:8000/books)  


