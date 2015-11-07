##Instalation
- Install Node.js * [Node.js](https://nodejs.org)

##Commands
Start Application ``DEBUG=signedoff npm start``

Rebuild React js files ``gulp``

If you install a new package, make sure you run ``npm install --save-dev request``

##Running
The web app will run on ``http://localhost:3000``

##Endpoints
- /releases/[username]/[repository]
- /release/[username]/[repository]/[releasename]
- /users
- /user/[username]
- /settings
- /builds/[username]/[repository]
- /action/signoff/[releasename]/[build]/[commit]
- /action/comment/[releasename]/[build]/[commit]
- /action/add/[build]
