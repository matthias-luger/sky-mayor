## Sky Mayor Service

Service to store the voting and result of the mayor elections.

## Start

-   docker compose up
-   go get
-   go run ./cmd/sky-mayor

## REST

-   POST /electionPeriod
-   GET /electionPeriod/:year
-   GET /electionPeriod/range/:from/:to

### Configuration

See .env.dist
