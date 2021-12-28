# Redis Publisher for General Events

- Publish to Redis General events database

## Environment variables

### Access to General Events Redis Container

- GENERAL_REDIS_HOST=general-events-redis
- GENERAL_REDIS_PORT=6379
- GENERAL_REDIS_DB=0
- GENERAL_REDIS_PASSWORD=glblEvntRds2021Dktg_acsIntrm1
- GENERAL_REDIS_CHANNEL=main

### Publishing to Redis
    
- PUBLISH_INTERVAL_MLSC=1
  - interval in milliseconds
- PUBLISH_AMOUNT=0 
  - 0 = infinite
- PUBLISHER_ID=Golang-Publisher

### Log file path

- LOG_FILE_PATH=/log/app.log
