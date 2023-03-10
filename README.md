# cloaki


*A lightweight, fast and secure web-based secrets store*

<p align="center">
<img style="text-align: center;" display="block" width="400" height="400" title="Cloaki" src="https://github.com/zalgonoise/cloaki/blob/media/cloaki.png">
</p>


<p align="right">
<a href="https://www.instagram.com/notpiupiuken">
<span style="font-style: italic;">Cloaki artwork by @notpiupiuken</span>
</a>
</p>

___________

## Concept

Cloaki is a secrets store served as a web app. It allows storing secrets (passwords and any type of confidential values) in a secure way, as a locally-deployed service (at home, in your cloud server, etc).

It supports multiple users (in accounts secured by a password defined by the user) that can share secrets among themselves, for a certain period of time. Shares will always have an expiry date, either defined by the user as a time or duration value, or one month by default.

Regarding security, all secrets are encrypted with a key unique to each user, and decrypted when read. User's private keys are generated and stored in Cloaki, where the server full manages both key access and secrets encryption / decryption. User passwords are not stored in the databases, instead they are salted and hashed with a secure algorithm, where the database will store the salt and hash alone. Cloaki is not vulnerable to rainbow-table attacks.

User sessions are managed with JWT, with a signing key that can either be defined by the admin or generated by Cloaki on the first run. Issued JWT expire within one hour.

Cloaki uses Bolt as a key-value store for all confidential data (secrets values, JWT), and SQLite to store user, secrets and shares metadata. Bolt and SQLite data can be persisted to a file.

____________

## Installation

To install `cloaki` you can use Go to run it as a binary:

```
go install github.com/zalgonoise/cloaki@latest
```

...or Docker:

```
docker pull zalgonoise/cloaki:latest
```

____________

## Configuration

Configuring the app's runtime can be done with:

- CLI flags
- OS environment variables
- (Docker) compose file

### CLI flags

Below is the list of options that can be configured as CLI flags:

Flag | Type | Default value | Description
:---:|:----:|:-------------:|:-----------:
`-port` | `int` | `8080` | port to use for the HTTP server
`-bolt-path` | `string` | `"/cloaki/keys.db"` | path to the Bolt database file
`-sqlite-path` |  `string` | `"/cloaki/sqlite.db"` | path to the SQLite database file
`-jwt-key` |  `string` | `"/cloaki/server/key"`  | path to the JWT signing key file
`-logfile-path` |   `string` | `"/cloaki/error.log"`  | path to the logfile stored in the service
`-tracefile-path` |  `string` | `"/cloaki/trace.json"`  | path to the tracefile stored in the service

### OS environment variables

Cloaki can also be configured with OS environment variables, which will shadow the CLI flag configuration of the same kind (OS env takes precedence over CLI flags).


Var | Type | Description
:---:|:----:|:-----------:
`CLOAKI_PORT` | `int` | port to use for the HTTP server
`CLOAKI_BOLT_PATH` | `string` | path to the Bolt database file
`CLOAKI_SQLITE_PATH` |  `string` | path to the SQLite database file
`CLOAKI_JWT_KEY_PATH` |  `string` | path to the JWT signing key file
`CLOAKI_LOGFILE_PATH` | `string` | path to the logfile stored in the service
`CLOAKI_TRACEFILE_PATH` | `string` | path to the tracefile stored in the service

### Docker compose

The `docker-compose.yaml` allows for the admin to configure the service's exposed HTTP port and to define a volume (for example, a Docker volume or a folder) to persist the container's data:

```yaml
    ports: 
      - 8080:8080
    volumes:
      - /tmp/cloaki:/cloaki
```

________


## Endpoints

To interact with Cloaki, you can use the HTTP API directly. Plans for a CLI client and a Flutter app in the future.

### Users 

Auth | Path | Method | Description | Post data
:---:|:---:|:------:|:-----------:|:---------:
N | `/users/` | `POST` | Creates a user | `{"username":"myUser","name":"User","password":"mySecretPassword"}`
Y | `/users/` | `GET` | Lists all users |
Y | `/users/{username}` | `GET` | Fetches the user |
Y | `/users/{username}` | `PUT` | Updates the user('s name) | `{"name":"NewName"}`
Y | `/users/{username}` | `DELETE` | Deletes the user |

### Secrets

Auth | Path | Method | Description | Post data
:---:|:---:|:------:|:-----------:|:---------:
Y | `/secrets/` | `GET` | Lists all secrets owned and shared with the user |
Y | `/secrets/` | `POST` | Creates a new secret | `{"key":"mySecret","value":"myValue"}`
Y | `/secrets/{key}` | `GET` | Fetches the secret |
Y | `/secrets/{key}` | `DELETE` | Deletes the secret |

### Shares


Auth | Path | Method | Description | Post data
:---:|:---:|:------:|:-----------:|:---------:
Y | `/secrets/{key}/share` | `POST` | Shares the secret with other users. May include a duration (`dur` field) or a time (`time` field) | `{"targets":["otherUser","myFriend"]}`
Y | `/shares/` | `GET` | Lists all secrets that the user has shared with other users |
Y | `/shares/{key}` | `GET` | Fetches the secret's share metadata |
Y | `/shares/{key}` | `DELETE` | Unshares the secret, optionally specifying the users to remove (otherwise unsharing with all users) | `{"targets":["otherUser"]}`

### Sessions


Auth | Path | Method | Description | Post data
:---:|:---:|:------:|:-----------:|:---------:
N | `/login` | `POST` | Sign in action which returns a new JWT if successful | `{"username":"myUser","password":"mySecretPassword"}`
Y | `/logout` | `POST` | Terminate the user's active sesion (invalidates the JWT) | 
Y | `/recover` | `POST` | Updates a user's password. Requires the previous password. | `{"password":"mySecretPassword","new_password":"myNewSecretPassword"}`
Y | `/refresh` | `POST` | Refreshes the user's JWT | `{"username":"myUser"}`


___________

## How it was made

See the entire process of [designing, writing and building Cloaki](https://github.com/zalgonoise/research/blob/master/ddd/index.md)!

___________


## Contributing

There are a ton of features (and flaws, surely!) that deserve the attention, despite this being a cozy, home-focused, local deployment type of project. 

You can suggest changes by opening an issue or to contribute directly with a PR. All input is welcome!