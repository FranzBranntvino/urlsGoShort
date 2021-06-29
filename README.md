# URLs GO Short

## Task

Create and document a small web service exposing URL shortening functions.
One should be able to create, read, and delete shortened URLs.
The API functions will be exposed under the ```/api``` path while accessing a shortened URL at the ```root``` level will cause redirection to the shortened URL.

### Rules of the game

- Code in Golang (unless differently agreed)
- It's ok to forget about permissions (everyone can do anything) for the sake of the exercise.
- Documentation will be publicly exposed via a specific service URL.
- Code will be tested to a reasonable extent
- You're free to choose any storage mechanism you wish

We expect to be able to run the application locally just following the project README documentation.
Do not assume the host running your service will meet any requirement (no storage engine is pre-installed).

### Bonus

- Implement a counter of the shortened URL redirects
- Add an API endpoint to read shortened URL redirects count

## Philosophy of Planning and Design

- Assume you always start wrong (being ill informed on the beginning of any project is the only thing that is right), makes starting assumptions always questionable
- Assume there will always be change (make it easy to be able to change things and revise decisions taken in the beginning)

## How it works

In short it has as a prerequisite:

- A machine running Docker Daemon and Docker Compose

### Start up the Services

```bash
# build and run services:
docker-compose up --build
# stop and remove services:
docker-compose down
```

or for already existing containers

```bash
# start services:
docker-compose start
# stop services:
docker-compose stop
```

### Try out the Demo

Visit [http://localhost:8080](http://localhost:8080) for a Demo-Frontend.

The statistics for a short-link can be requested via [http://localhost:6060/pkg/urlShortener/"put_url-Code_here"](http://localhost:6060/pkg/urlShortener/"put_url-Code_here")

### Look at the ```godoc``` Documentation

The documentation for the Go-App is available under [http://localhost:6060/pkg/urlShortener/](http://localhost:6060/pkg/urlShortener/).

Specifically for the REST-API follow the link [http://localhost:6060/pkg/urlShortener/urlHandler/#pkg-constants](http://localhost:6060/pkg/urlShortener/urlHandler/#pkg-constants).

### Dev-Setup

See the [.devcontainer/README.md](.devcontainer/README.md) for more information on this.

Prerequisite: a machine with Docker running and VSCode (including Remote Plugins installed).
