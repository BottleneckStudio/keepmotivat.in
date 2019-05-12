# keepmotivat.in
[![Build Status](https://travis-ci.org/BottleneckStudio/keepmotivat.in.svg?branch=master)](https://travis-ci.org/BottleneckStudio/keepmotivat.in) [![Maintainability](https://api.codeclimate.com/v1/badges/df3b7ed02f0cc01f6291/maintainability)](https://codeclimate.com/github/BottleneckStudio/keepmotivat.in/maintainability) [![Test Coverage](https://api.codeclimate.com/v1/badges/df3b7ed02f0cc01f6291/test_coverage)](https://codeclimate.com/github/BottleneckStudio/keepmotivat.in/test_coverage) [![Go Report Card](https://goreportcard.com/badge/github.com/BottleneckStudio/keepmotivat.in)](https://goreportcard.com/report/github.com/BottleneckStudio/keepmotivat.in)


> An app that keeps you motivated. 🏋 💪


# Installation:
```sh
$ cd $GOPATH/src/github.com && mkdir BottleNeckStudio
$ git clone git@github.com:BottleneckStudio/keepmotivat.in.git
$ dep ensure
$ ./scripts/pre-commit-setup.sh
```

# npm

After cloning run the following commands.

```
$ npm install
$ npm run build
```

This will ensure that uikit and all future `stylesheet` assets will be bundled by **Webpack**.

## Adding new stylesheets

If you want custom stylesheet files you must add them under `./app/data/assets/stylesheets/main/`
and reference them in `.app/data/assets/stylesheets/main.scss` by adding the following.

```
*main.scss*
@import './main/*custom-scss-file.scss*'
```
You have to run `$npm run build` below everytime you will add a new `scss` file or change the current stylesheet files to include them with the current webpack bundle.
