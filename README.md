# backend
Backend service of AWS Billing Automation

# Build Status

[![Build Status](https://github.com/javed-tft/aws-billing-app.git)

# Development setup

To run the script, it takes nothing more than:

  - Bash
  - Go
  - AWS CLI
  


## HowTo
### Configure the AWS CLI 

 ```bash
   -- open a terminal and hit command: aws configure
   -- set credentials like access key, secret key, region, output format etc.
   -- choose region ap-south-1 because the service charged is less for this region.

### Run the Application
- Open the terminal from root folder
- hit command `cd cmd` then `go run main.go`



### Run Tests
```bash
- hit command `go test`



for this app...

You need to install golang in your system
need code editer (i'm prefer VS code)

install AWS CLI
configure credential
            -- open terminal and hit commond: aws configure
            -- set credential like, access key, secret key, region, output format etc.
            -- choose region ap-south-1 because service charged is less for this region.

in the AWS Billing app...
we(peers) have used several function

The function getDate(), it returns range of date in the format YY-MM-DD

The function GetAwsCost() is where we request the information to AWS, this function gives aws servises cost.
we set the level of granularity (daily) or Monthly or yearly as per our requirement.

create a file and store in CSV.
CSV was chosen for ease of use, almost every spreadsheet program can open those and if not the we processed with a simple text file

The main function is where we use the AWS CLI library, for executing commands, and calling functions

after all done.....
hit the command go run main.go
it will genrate a csv file and you simply open it in spreadsheet
