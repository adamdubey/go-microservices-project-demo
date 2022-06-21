# go-microservices-project-demo

## Quick Start

_Building & Starting Workloads:_
```sh
# Build & Start container services
make up_build

# Build front-end
make build_front

# Start front-end
make start

# in browser, visit http://localhost
```

_Stopping Workloads:_
```sh
# Stop all container services
make down

# Stop front-end
make stop
```

---

## Technologies & Frameworks

- [bcrypt](https://pkg.go.dev/golang.org/x/crypto/bcrypt)
- [chi](github.com/go-chi)
- [Docker](https://www.docker.com/)
- [Go](https://go.dev/)
- [make](https://www.gnu.org/software/make/manual/make.html)
- [pgconn](https://github.com/jackc/pgconn)
- [pgx](https://github.com/jackc/pgx/v4)
- [PostgreSQL](https://www.postgresql.org/)
