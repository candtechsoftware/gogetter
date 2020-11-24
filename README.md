# GoGetter
Tool to quickly create go projects and install multiple go packages

## Install
Tested only on Mac for now. Will be testing on Linux and windows. 
`GOBIN=/usr/local/bin go install`

Future Installs will hopefully be throuhg homebrew (mac), pacman/apt (linux) and chocolatey (windows)


## Basic Usage
```
  gogetter mk <project-name>  -u <username>  
  cd <project-name>
  gogetter <enter multiple packages space seperated>....
 ```
## Default settings 
 - By default it will create a git repo you can set this to false by using the git flag `-git=false`.
 - It will create a main.go by default. I plan on adding basic templates. Inspired by express generator. So eventually it would be able to have a type flag and then create the desired folder structures with base files. Examples `-t=rest-web-service` `-t=proto-web-service` `-t=cli-tool`. Open to prs of differnt templates when that feature is implemented. 
 - Gogetter uses go modules so it will by default it will use the module naming convention "github.com/<username>/<package-name>" if supplied a username. If no username is supplies it will just uset the <project name> has the package name. 
 - If you would like to specify a different domain you can use the `-d` flag and set the domain. Ex `gogetter mk coolproject -u goku -d google.com` the package name will be `pacakge google.com/goku/coolproject` 
  
 
