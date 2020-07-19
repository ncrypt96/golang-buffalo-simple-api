
# Welcome

> The app was created while learning buffalo and golang

This application is developed using golang, buffalo and boltDB
> You can learn more about golang [here](https://golang.org/)
> You can learn more about buffalo [here](http://gobuffalo.io)  
> You can learn more about boltdb [here](https://github.com/boltdb/bolt)  


## Starting the Application

  

Buffalo ships with a command that will watch your application and automatically rebuild the Go binary and any assets for you. To do that run the ```buffalo dev``` command.
You can also run the same command using the command ``make run``

 
  

If you point your browser to [http://127.0.0.1:3000](http://127.0.0.1:3000) you should see a "Welcome to Buffalo!" page.

**Congratulations!** You now have your Buffalo application up and running.

### Overview
This is a relatively very simple app made during exploring buffalo
The app was created using buffalo cli as an api service without a database
The app lets you add name and quotes on the route u/add
and lets you view the quote on the route u/get
The app uses boltDB for storing the data

*** The API documentation can be found [here](https://documenter.getpostman.com/view/8732425/T1DjkKba) ***

### Action generation
 buffalo comes with various configurations for generating action 
 you can read more about it [here](https://gobuffalo.io/en/docs/actions)
 However you can also generate a subset of actions using the command
 ```shell
 make action r=route s=subroute m=METHOD
  ```
  example
  ```
  make action r=auth s=signup m=POST
  ```
This will generate /auth/signup 

### Creating a build
You can use the command ```buffalo build``` or ```make build```
to build the server.


## What Next?

  

I recommend you heading over to [http://gobuffalo.io](http://gobuffalo.io) and reviewing all of the great documentation there.
  

Good luck!

  

[Powered by Buffalo](http://gobuffalo.io)

  
###  API documentation
https://documenter.getpostman.com/view/8732425/T1DjkKba