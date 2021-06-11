# URLs GO Short

## Task

Create and document a small web service exposing URL shortening functions.
One should be able to create, read, and delete shortened URLs.
The API functions will be exposed under the ```/api``` path while accessing a shortened URL at the ```root``` level will cause redirection to the shortened URL.

## Rules of the game

- Code in Golang (unless differently agreed)
- It's ok to forget about permissions (everyone can do anything) for the sake of the exercise.
- Documentation will be publicly exposed via a specific service URL.
- Code will be tested to a reasonable extent
- You're free to choose any storage mechanism you wish

We expect to be able to run the application locally just following the project README documentation.
Do not assume the host running your service will meet any requirement (no storage engine is pre-installed).

## Bonus

- Implement a counter of the shortened url redirections
- Add an API endpoint to read shortened url redirections count
