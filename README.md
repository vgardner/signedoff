signedoff.io API
================
##Installation
- [golang.org]

##Commands
Start Application ``go run main.go``

##Running
The API will run on ``http://localhost:3002``

##Updating dependencies using Godep

        go get github.com/gorilla/mux
        godep save
        godep go build

For more info, here's a good step-by-step on Godep:

* [Godep Tutorial](http://www.goinggo.net/2013/10/manage-dependencies-with-godep.html)

##API Endpoints
- /releases/[username]/[repository]
- /release/[username]/[repository]/[releasename]
- /users
- /user/[username]
- /settings
- /builds/[username]/[repository]
- /action/signoff/[releasename]/[build]/[commit]
- /action/comment/[releasename]/[build]/[commit]
- /action/add/[build]




