## Golang Base Project
This app can be launched either locally or on virtual machine, it is heavily based on [uberswe/golang-base-project](https://github.com/uberswe/golang-base-project)

Following functionalities are implemented
- list users (after registration and login)
- register account
- login into account
- activate account
- reset password
- resend activation email

## Demo
https://filmbox-test.pl/

| username | password |
|:--------:|:--------:|
| testuser | testuser |

## Deployment
1) copy file config/example.env into .env in root directory
2) verify that all credentials in .env and deploy.sh are correct
3) place ssh password to your server in file named "pass_file" in root directory
4) run deploy.sh script with `bash deploy.sh` or `./deploy.sh` (requires rsync installed both where you launch script and where you want to copy files into)
5) log via ssh into your server `ssh -p 22 <username>@<hostname>`
6) run `docker compose up --build --force-recreate`
7) if you want also https to work you need to put valid cert and key files into certbot/live/__yourdomain__ folder

## Notable Features
+ implemented overlay above standard gorm api allowing to build custom multistage operations like paging
+ experimented with golang standard templates, defined custom funcmaps that allow to recursively render templates and pass easily arguments from top call
+ loading html template files recursively, now allowing to have tree structure and better organize files by modules
+ custom logger over gorm api that allows to gather all messages in database table
+ frontend uses bootstrap v4 and works also on mobile devices
+ separate module for all validations
+ in-app switching between (available) languages
+ configuration through .env file
