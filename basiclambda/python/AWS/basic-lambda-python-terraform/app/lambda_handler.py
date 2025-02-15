import json

print("Loading function")

def lambda_handler(event, context):
    print(f"Received event: {event}")
    return {
        'statusCode': 200,
        'body': json.dumps('Hello world from basic-lambda-python-terraform Lambda!')
    }