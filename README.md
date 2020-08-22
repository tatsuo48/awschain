# awschain

awschain is set current AWS* environment variables to envchain namespace.  

envchain is here.  
https://github.com/sorah/envchain

If you are using AWS SSO, envchain configuration updates will occur frequently.  
This tool will help you update your envchain configuration.

## Installing
```
go get -u github.com/tatsuo48/awschain
```

## Usage
```bash
# Set AWS Credentials in your terminal
$ export AWS_ACCESS_KEY_ID="XXXXXXXXXXXXXXX"
$ export AWS_SECRET_ACCESS_KEY="XXXXXXXXXXXXXXX"
$ export AWS_SESSION_TOKEN="XXXXXXXXXXXXXXX"

$ prinetenv | grep AWS
AWS_ACCESS_KEY_ID=XXXXXXXXXXXXXXX
AWS_SECRET_ACCESS_KEY=XXXXXXXXXXXXXXX
AWS_SESSION_TOKEN=XXXXXXXXXXXXXXX

$ aws s3 ls 
2020-05-16 21:00:52 test-bucket1
2020-05-22 23:45:11 test-bucket2
2020-05-23 00:01:45 test-bucket3

$ awschain test-account
```

Run in another terminal
```bash
$ envchain test-account aws s3 ls
2020-05-16 21:00:52 test-bucket1
2020-05-22 23:45:11 test-bucket2
2020-05-23 00:01:45 test-bucket3
```
