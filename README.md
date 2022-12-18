# GoBlog

## About 

This repo is a pure practice for go principlies to improve my skills.

## Packages Used

- viper
- mux
- gorm
- docker-compose

## API Usage

1. Remote Instance

The application is deployed on heroku. You can access the API at [GoBlog Remote Server](http://37.46.128.188/goblog/docs/index.html)

2. Local Instance

To run the local instance, you first have to set up the env vars, You can use the default ones in the `.*.env.example`.

```bash
cp .env.example .env
cp .db.env.example .db.env
```

```bash
docker-compose up 
```

Then navigate to [Local Instance](http://localhost:9000/docs/index.html) to see the API in action.
