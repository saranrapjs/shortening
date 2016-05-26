*"A rudimentary link shortener"*

This is a tiny go app which stores/updates/redirects a set of slugs to links.

## Local Setup

### Docker & docker-compose

#### Docker for Mac beta!

The simplest way to spin up the app is via docker + docker-compose.

To *truly* make the future now, consider signing up for the [Docker for Mac beta](https://beta.docker.com/), which enables Mac volume sharing and hot reloadable code inside the Docker container. More info [here](https://blog.docker.com/2016/03/docker-for-mac-windows-beta/) and [here](http://jdlm.info/articles/2016/03/06/lessons-building-node-app-docker.html). 

Without the Beta, static files aren't pushed into the container without rebuilding the container image.

#### Regular docker-machine + docker on a Mac

For a regular `docker-machine` setup, try the following:

1. Install docker and docker-compose [here](https://docs.docker.com/v1.8/installation/mac/)
2. Create a docker machine: `docker-machine create --driver virtualbox default`
3. Connect to your newly created machine: `eval "$(docker-machine env default)"`
4. Run docker-compose: `docker-compose up` (this will build the docker image locally, and spin up the image + its dependencies)
5. ...alternately, you can run it in the background: `docker-compose up -d`

### Rebuilding the container image

To rebuild the container image:
```
docker-compose build shorteningapp
```

### Go development

Requires a valid $GOPATH, & the app is configurable via environment variables.

e.g.:
```
export DBHOST="root:@tcp(database-hostname-here):3306)"
go run main.go
```

### JS development

Source code lives in `src/`, and is packaged by webpack + babel into `dist/`, which is where the go app reads its static files. These static files are rewritten to end up at `/manage/` in the browser.

Setup:

1. `npm install`
2. `npm run build`
3. If you haven't spun up the go app, there's an alias at: `npm run run-server`

*Why are you inlining the webpack config as command line arguments to the build commands?*

Great question; I find it insane and wild that every little javascript tool expects to be given its own configuration file; in the case of Webpack, it seems equally insane and wild that many simple setups (just transpiling es6, or es6+react transpilation) can be satisfied with a somewhat-terse command line one-liner.

I'm sure that I will come to disagree with myself as this webpack config grows more baroque :)