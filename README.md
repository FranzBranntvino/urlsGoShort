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

- A machine running [Docker Daemon](https://docs.docker.com/engine/install/) and [Docker Compose](https://docs.docker.com/compose/install/)

### Start up the Services

```bash
# build and run services:
docker-compose up --build
```

```bash
# stop and remove services:
docker-compose down
```

or for already existing containers

```bash
# start services:
docker-compose start
```

```bash
# stop services:
docker-compose stop
```

### Try out the Demo

Visit [http://localhost:8080](http://localhost:8080) for a Demo-Frontend.

The statistics for a short-link can be requested via [http://localhost:8080/shortUrlStats/"put_url-Code_here"](http://localhost:8080/shortUrlStats/"put_url-Code_here")

### Look at the ```godoc``` Documentation

The documentation for the Go-App is available under [http://localhost:6060/pkg/urlShortener/](http://localhost:6060/pkg/urlShortener/).

Specifically for the REST-API follow the link [http://localhost:6060/pkg/urlShortener/urlHandler/#pkg-constants](http://localhost:6060/pkg/urlShortener/urlHandler/#pkg-constants).

### Dev-Setup

Prerequisite:

- A cloned Repo:

    ```bash
    git clone https://github.com/FranzBranntvino/urlsGoShort.git
    ```

- A machine with [Docker Daemon](https://docs.docker.com/engine/install/) running and [VSCode](https://code.visualstudio.com/) including [Remote Plugins](https://code.visualstudio.com/docs/remote/remote-overview) installed.

    See the [.devcontainer/README.md](.devcontainer/README.md) for more detailed information on this subject.

***NOTE:***
It might be that in the current Dev-Setup some of the Tests are failing if the DB-Service are not running, hence the DB-client is unable to connect to the DB.

## Areas for Improvement

- [x] Some Status Codes of the response are not used right: [https://en.wikipedia.org/wiki/List_of_HTTP_status_codes](https://en.wikipedia.org/wiki/List_of_HTTP_status_codes) and [https://cheatography.com/jp153/cheat-sheets/http-status-codes-golang/](https://cheatography.com/jp153/cheat-sheets/http-status-codes-golang/)
  - Post [http://localhost:8080/createShortUrl](http://localhost:8080/createShortUrl): ```200``` response, should become: ```201``` (Created)
  - Post [http://localhost:8080/shortUrlDelete](http://localhost:8080/shortUrlDelete): ```302``` response -> missing Location header, should become: ```200``` (OK) or ```204``` (No Content)

- [x] Improve on REST API, e.g. Naming, Documentation, Functionality, ```GET```, ```POST``` and ```DELETE``` usage, return JSON

- [ ] Separate concerns for Unit and Integration tests: provide mocks and stubs for the unit tests to work sufficiently and independent from other modules (DB Service)

- [ ] Default Handling for undefined routes

- [ ] Implement graceful stopping of http-Server

- [ ] Implement graceful stopping of DB-client

- [ ] Build in scalability for the services

- [x] Make Mega-Links work: e.g. long (>2000 chars) full version of [https://godbolt.org/z/9TeK9hYh9](https://godbolt.org/#z:OYLghAFBqd5QCxAYwPYBMCmBRdBLAF1QCcAaPECAM1QDsCBlZAQwBtMQBGAFlJvoCqAZ0wAFAB4gA5AAYppAFZdSrZrVDIApACYAQjt2kR7ZATx1KmWugDCqVgFcAtrS6dSV9ABk8tTADlnACNMYhBuADZSAAdUIUJzWjtHFzcYuIS6Hz9ApxCwyKNMEzM6BgJmYgJk51dOd2NMU0TyyoJsgODQ8KihCqqa1PqjfvbfTrzuyIBKI1QHYmQOKR0AZl9kRywAak1Vm0riZgBPPexNGQBBNY2tzF398z7iTGYnM4vr7XXaTYcdvY2ZDPXzAD5XT6fAD0ACptuViA5TNssFRfJlaEJtjCoZ8CMdophUdtnkiCNsAPpCVBOTAU5gOIgU4BWULMAhEiloAEAdn0V22gu2QVQ9m20WIqEJVWOMj2/MuQuFotY4sl0vxnHlnyVIrFEqloXx2m1AqFetVBo1x1WpsV5pVasNMu4dt1jqtRuOAFY7ZoeQAREk0ukMpksvxHDnoLkYTB%2Bq5QpPbADirKjRO2bDwzCxNGI2wAKgT41cHPF1NtROqjXhMEIAErFV4iB5B6m0%2BmM1DM9Psznc%2BOrBWfPrsvDIKs1qp1xvN3P3ABubAcmF0DioVFCbd2fJ1QqobBEpH3gsPrGPp%2B258vZrPR8wJ7v14fT/t94vj8hgYTlw5TmiVQOUBTZcyxQtwUuRdUDwdBtlUAAvY47AAypMARMkIDQTFyQg7QIm2ZdWGmXcFWTIUxzMSdC0Ilc1w3LdiDdIU8CobYIHYvp0BAEBaScZAAIgF5fA5YgJUwAguVzAgQIQSptmwvoYTOCAdAiIjV3XTdQlmbZhPoUJxMklg%2Blk%2BTFIIZTVmwVT8KI3T4gQzBUCoCBC2mEiwDAPYgxkEiSP9BUlSVLieLQRkHhsQFdm0bRtn8AB5GjMAARwcNgYriwFotCkBPFYZjgtoxx6O0gsfOKwrdyDYpW0Cq8isFXLwvJbL9ky7ZUvS1UdCy/YcoIbi8usArhyvf0A2/SaIRmy5KInBS6D6TBxAlYNOzDHsIzZaNYx2ABratnTMesm3YBdOB3ernxvL8btfK9brfJUnseh77s/N8Jt/UcKioxacJWtaO1Dbte0jfsY0HbZDunE653OkQ%2BqDa73xfT63oxj7bzRghETutHXuxgnvrG2bPhE7YnGYXwIACvdn2a%2BZWv69rHOc1yNNKxiAtZmwSUGnj8t/JVEOQmlojQjDTAgWHjtnM6W0wThpiqsWUMll5pYIWWjo1BX5yR1WyeuZ8XgIBZaG2OUTe%2B2apFmVhpG9eRXFkeRUGkKK9AMYMFiWGLVk4eQCGkOQPNIfbwhkAA6VYZAADm0HkE55bgeE4b0AE58JUaRuHkJwuBkGRSDduRSE9qR5CEEBS9D93ZjgWAYEQEBmeiRlyEoNAALwdgwncIlCBIZR%2BEYFh2C4Xhx%2BEMRJHd0gAHcjmiaRg8d53XbDj3pASxlO/JFztkuBsAFltmAZBJyzmPtBjy6IFoOhMBIiANf77c1hVkOd4jqPtAJxjjIVYPJVgJyDjyHk3oeQlxznnKQBdSBF04LfEunBE4wJ5BEVY8cZDYLLjvSu0ga511IA3cOpBm5txWk0bsZAKAQEqMAIQog1DFAYAgVAS83bB1IL3aI/dxx0DYX4VgnDuHl3kAIz%2BYRgAYNWPwiWsj96MC4TwohtDkCXGICwkhHhxBNHKPgN28hx5MDYBwHgfA6AEDnhIIhSh3CqHUCgAwBgVB4CCLXSAswpSlExPo6k/srEjBMaIjh6jeHyBXswNeUgN4IJdoQxeVdxAJwiAAWgiNwS%2B19tgYLjtsbAhjkD0PYrgEe5Vvgq22N7fQehf6N1mAgV4WAwh0wQUgouCdY5ZyzgnLOPBYqrAiBEfpiipHEOrkYMhFCHazAAdwGO3oZCrNiunDOqdOAJwQasbeqT9HzKbq3CASAZED27lhZRA8QCVCcDndwaJWCiVrhAIIRCgi%2BEqMcde0iaS0noAlWgrBfmLywNTdQU9wV4BeM0Rc9ZNGlMZMsPhIlihENYF4o4xBxZYCIXjPARcEkOxsfQCxU9rGzxEA4xeTiVBqA0O4vQnjvHwD8dEAJtcpCZISvsooJREiWGsIMOoHhrAdFyPkZQsR4gBNFTKjIATJVdEHgKpoATWgDHsLUZQjRmhlFGCqyYaqxzapSGKs1Ywciqq4LMYJiwrGbykMkqZVdT4XyvjfO%2BD92LPz8G/SpRBqlB1IHUm5X8akkXqQYJp4cWltO6J0qOqxvQPwyeMhOCduA8k4BEb0qxChO0QQciuVdSH1z/s67Qpbd4zOObMBFxB4gWG4EAA%3D%3D%3D)
  - reference: [https://bitly.com/](https://bitly.com/) does not work with it either and it is supposing this seems to be an invalid URL

- [x] Handling of local links (e.g. local files), despite they lead to security violations: [file:///home/vagrant/urlsGoShort/.devcontainer/README.md](file:///home/vagrant/urlsGoShort/.devcontainer/README.md)
  - Second besides the formatting issue is, that this is a Security issue in some Browsers: ```Security Error: Content at http://localhost:8080/ may not load or link to file:///home/vagrant/urlsGoShort/.devcontainer/README.md.```
  - reference: [https://bitly.com/](https://bitly.com/) does not work with it either and it is supposing this seems to be an invalid URL

- [ ] Improve on Frontend formatting e.g. representation of very long links, Error Responses/Handling/Messaging
