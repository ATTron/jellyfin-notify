# Jellyfin-Notify

Send messages out via [AWS Simple Notification Service(SNS)](https://aws.amazon.com/sns/) in a timed fashion whenever new items are added to your [Jellyfin Server](https://jellyfin.org/)

## Prerequisites
Jellyfin-Notify assumes you have the [AWS CLI](https://aws.amazon.com/cli/) installed and configured. Also that you have created a valid SNS topic in AWS.

## Usage
Jellyfin-Notify can be used in 2 different ways

### 1. CLI w/ flags
```bash
./jellyfin-notify -endpoint <ENDPOINT> -api-key <API-KEY> -user-key <USER-KEY> -aws-region <AWS-REGION> -sns-arn <SNS-ARN> -wait-time 72
```

### 2. CLI w/ .env file
* create an .env file using the template in this repo
```bash
./jellyfin-notify -env-file
```

