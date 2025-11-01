# game-assistant-server

## Prerequisites
Install both docker and docker compose.

## How to start
- Prepare **.env** file containing env variables defined in *docker-compose.yml* file. Values for variables don't matter at this stage.
- Choose which game database to use:
  - `docker compose up --build` (default)
  - `GAME_DB=hannibal docker compose up --build`
  - `GAME_DB=ato docker compose up --build`
- With debug mode (includes adminer):
  - `docker compose --profile debug up --build`
  - `GAME_DB=hannibal docker compose --profile debug up --build`
- Run localhost:8080 in your browser to access **adminer** - lightweight database viewer. Note: only for debug mode.
- And obviously `docker compose down -v` for teardown, to remove containers and volumes

## Notes
- Database schemas are loaded only once when the volume is empty. To switch between different games or reload schemas, call **docker compose down -v** first.
- Available games: `hannibal`, `ato`