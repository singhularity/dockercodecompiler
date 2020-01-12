**Checkout code**
* Install go
* Checkout code in $GOPATH/src/github.com

**Run Docker setup**
* `cd docker_setup`
* `./BuildCompilerDockerImage.sh`

**Build Library**
* `go build start.go`

If you see errors like "undefined: client.NewClientWithOpts", run the following
* `go get github.com/docker/docker@master`

**Run as a service OR Local**

<b>To run as service</b>

`./start svc`

You can then Post requests to http://localhost:8090/api/compile

Java Example

`curl -X POST \
   http://localhost:8090/api/compile \
   -d '{
   "language": "java",
   "code": "class HelloWorld { public static void main(String[] args) { System.out.println(\"Hello \" + args[0]);}}",
   "stdin": "World!"
 }'
 `

Python Example

`curl -X POST \
  http://localhost:8090/api/compile \
  -d '{
  "language": "python",
  "code": "print('\''Hello World!'\'')"
}'
`

To run stand alone
* `./start local java "class HelloWorld { public static void main(String[] args) { System.out.println(\"Hello \" + args[0]);}}" world!`
* `./start local python "print('Hello')"`

**Supported Languages**
* Java
* Python3
* Python
* PHP
* Ruby
* Scala
* GoLang

