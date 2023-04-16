# AWS Lambda - S3 Logger (Go)

This is a simple AWS Lambda function written in Go that logs the S3 event to CloudWatch Logs.


## Build
```bash
$ GOOS=linux go build -o main
$ zip deployment.zip main
```


## Usage and Configuration

1. Create a new Lambda function
   - Open the AWS Management Console and navigate to the Lambda service.
   - Click "Create function".
   - Choose "Author from scratch".
   - Enter a function name, for example, "s3-data-transfer-logger".
   - Choose "Go 1.x" as the runtime.
   - Under "Function code", click "Upload from" and select ".zip file". Upload the deployment.zip file created earlier.
   - In the "Execution role" section, choose "Create a new role with basic Lambda permissions". This creates an IAM role with basic permissions for your Lambda function. 
1. Configure an S3 trigger for your Lambda function:
   - In the "Function overview" section, click "Add trigger".
   - Choose "S3" as the trigger type.
   - Select the S3 bucket for which you want to track file downloads.
   - Set the "Event type" to "ObjectCreated (All)".
   - Optionally, set a prefix or suffix to filter events based on the object key (file path).
   - Click "Add" to create the trigger.
2. Configure permissions for your Lambda function:
   - Make sure that the IAM role associated with your Lambda function has the necessary permissions to access S3 and interact with CloudWatch Logs. To do this, open the IAM service in the AWS Management Console, find the role you just created, and add the following managed policies:
     - `AmazonS3ReadOnlyAccess`
     - `AWSLambdaBasicExecutionRole`
   - Customize the role as needed, based on your specific requirements.



## License

This project is licensed under the MIT License - see the [MIT license](LICENSE) for details.
