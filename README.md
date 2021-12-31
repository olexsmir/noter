# noter [backend]

![Go version](https://img.shields.io/github/go-mod/go-version/Smirnov-O/noter?style=flat-square)
![Repo size](https://img.shields.io/github/repo-size/Smirnov-O/noter?style=flat-square)

### Setup

Create .env file in root directory and add following values:

```shell
POSTGRES_HOST="localhost"
POSTGRES_PORT="5432"
POSTGRES_USERNAME="postgres"
POSTGRES_PASSWORD="postgres"
PASSWORD_SALT="random string"
JWT_SIGNING_KEY="random string"
```

### Make targets
- `run` - Build and start project
- `build` - Only build project
- `migrate.new name=migrateName` - Create new migration
- `migrate.up` - Setup migrations
- `migrate.down` - Remove last migration
- `mifrate.drop` - Remove all migrations
