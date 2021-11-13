# My water is down again!

This is a silly project that hopes to prove some regions in Auckland get more water outages than others (I swear I don't live that rurally!)

Unfortunately there is no public API for previous water outages, so the data collected by this is largely incomplete.

But one day, I'll prove it! Maybe.

## APIs

There is only one API, available at the root of the server.

It comes with the following (query) parameters:
- outage_type
- start_date (after)
- end_date (before)
- suburb
- street
- location

Data is collected hourly.

## Installation instructions

1. Copy the following files & make changes as needed:
    - docker-compose.yml
    - .env-example: Rename to .env when done
    - docker_postgres_init.sql
2. Pull prepared image from DockerHub and start: ```docker-compose up -d```
3. Navigate to localhost:APP_PORT (whatever you set up in the .env file)

## Live version

None at the moment
