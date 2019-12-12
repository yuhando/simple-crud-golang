The Installation
 
1. Open your Command Prompt
2. Run this command to pull or download the repo to your local gopath directory : go get -u -d github.com/yuhando/tniscodingtest
3. Restore the database on your local, the scheme there are on db folder
3. set db connection on dbconn.go file
4. run this command on your command prompt: go build
5. run this command on your command prompt: ./foldername-of-this-apps
6. access this app through http://localhost:8080/ this run on port 8080, you can change it on main.go file if you want.