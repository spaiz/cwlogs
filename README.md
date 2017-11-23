# CWLOGS - Cloud Watch Logs Downloader

Tiny utility for downloading log files from Amazon Cloud Watch.

# Installation using Docker
```sh
docker build -t cwlogs .
```

### Usage with default AWS credentials
```sh
docker build -t cwlogs .
docker run --rm -v $(pwd):/logs -e "HOME=/home" -v $HOME/.aws:/home/.aws cwlogs --group LOGS_GROUP --stream LOGS_STREAM
```

### Usage with custom AWS profiles

```
AWS_ACCESS_KEY_ID=$(aws --profile default configure get aws_access_key_id)
AWS_SECRET_ACCESS_KEY=$(aws --profile default configure get aws_secret_access_key)
docker run --rm -v $(pwd):/logs -e AWS_ACCESS_KEY_ID=$AWS_ACCESS_KEY_ID -e AWS_SECRET_ACCESS_KEY=$AWS_SECRET_ACCESS_KEY cwlogs --group LOGS_GROUP --stream LOGS_STREAM
```

### Native binary usage
```sh
go install
cwlogs --group group_name --stream stream_name
```

### AWS Config
Be sure you have `~/.aws/credentials` file with you AWS credentials:

```
[default]
aws_access_key_id = YOUR_KEY_ID
aws_secret_access_key = YOUR_SECRET
```

P.S.
CWLOGS will download logs until an empty results returned from AWS.
It means you shouldn't use it for logs that someone still writes to it.