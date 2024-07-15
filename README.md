# UDP app project

## Installation
There are two ways of setting this up locally. One is via docker-compose and the other is building the binaries directly on your machine.
The client app will be available on `localhost:8081` and the server will be using `localhost:8080` by default.

---

### Docker-Compose
Pretty self-explanatory

Build & run:
```sh
docker-compose up -d
```
Force rebuild:
```sh
docker-compose up --force-recreate --build
```

---

### Executable files
Use utils.sh to build (via make) and run the app. Refer to the utils.sh script for more info.

Allowed arguments:
- `build` - build binaries
- `exec` - runs the binaries in the background (need to be built first obviously)
- `clean` - cleans the contents of bin/ in services' respective directories
- `stop` - kills the processes associated with the services
```sh
./utils.sh <command>
```

