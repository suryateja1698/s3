# Steps to run AWS S3

1. Create account in [AWS console](https://aws.amazon.com/console/)
2. Open [IAM Console](https://us-east-1.console.aws.amazon.com/iamv2/home?region=us-east-1#/users)
3. On the navigation menu, choose Users.
4. Choose your IAM user name (not the check box).
5. Open the Security credentials tab, and then choose Create access key.
6. To see the new access key, choose Show. Your credentials resemble the following:
````
   Access key ID: AKIAIOSFODNN7EXAMPLE
   Secret access key: wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY
````

Store these keys in `~/.aws/credentials` file in following format:
````
[default]
aws_access_key_id = your-access-key
aws_secret_access_key = your-secret-key
````

# To run 
````
go run main.go
````
