# secretenv
This is a Golang package for reading secrets from files and environment variables.

# Usage

Inject secrets into your container.

You can optionally specify an environment variable suffixed with `_FILE` to indicate the location of the secret. For example, `DB_PASSWORD_FILE=/run/secrets/db_password`.

To read its content: 

```go
import "github.com/ndthuan/secretenv"

func main() {
    dbPassword := secretenv.Get("DB_PASSWORD")
}
```

If `DB_PASSWORD_FILE` does not exist or there is an error reading it, the value of the environment variable `DB_PASSWORD` will be returned.
