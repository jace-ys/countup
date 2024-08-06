# Count Up

A little web app game, consolidating the best practices from my years of experience building production-ready software and infrastructure.

Probably way overengineered for incrementing a counter, but thought this will be a good exercise to codify my learnings!

## Features

- Async, transactional worker system for processing background jobs using [River](https://riverqueue.com/).
- Compiled SQL queries into type-safe application code using [sqlc](https://sqlc.dev/).
- Declarative database schema and migrations using a combination of [Atlas](https://atlasgo.io/) + [Goose](https://pressly.github.io/goose/).
- DDD-lite, interface-driven approach to writing decoupled and testable business logic.