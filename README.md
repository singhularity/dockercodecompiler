<h1>About</h1>
<p>
The "Docker Code Compiler" can be used to run adhoc code in a "sandboxed" mode.
Basically, it brings up a temporary docker instance, runs your code and then returns any output/error generated after running the code.
The library can run as a service or for a "single" run 
</p>

<h1>Uses</h1>
<p>You can use it to run any code in a sandbox so you don't risk untrusted sources messing up your server contents.
A good example is an online code compiler application!
</p>

<h1>Set Up</h1>

* Install go
* Checkout code in $GOPATH/src/github.com

**Run Docker setup**
* `cd docker_setup`
* `./BuildCompilerDockerImage.sh`

**Build Library**
* `go build start.go`

If you see errors like "undefined: client.NewClientWithOpts", run the following
* `go get github.com/docker/docker@master`

<h1>Usage</h1>

**Can be run in "service" OR "Local" mode**

<b>To run as service</b>

    `./start svc`

You can then Post requests to http://localhost:8090/api/compile

* Java Example<br> 

    `curl -X POST \
       http://localhost:8090/api/compile \
       -d '{
       "language": "java",
       "code": "class HelloWorld { public static void main(String[] args) { System.out.println(\"Hello \" + args[0]);}}",
       "stdin": "World!"
     }'
     `

* Python Example

    `curl -X POST \
      http://localhost:8090/api/compile \
      -d '{
      "language": "python",
      "code": "print('\''Hello World!'\'')"
    }'
    `

<b>To run single instance</b>

* Java Example<br> 
    `./start local java "class HelloWorld { public static void main(String[] args) { System.out.println(\"Hello \" + args[0]);}}" world!`
* Python Example<br>
`./start local python "print('Hello')"`

<h1>Supported Languages</h1><br>
Java<br>
Python3<br>
Python<br>
PHP<br>
Ruby<br>
Scala<br>
GoLang<br>


