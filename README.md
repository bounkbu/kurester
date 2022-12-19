# KU Rester 🍱

## Project Overview 🔎

On everyday, there will be someone come up with the question "วันนี้กินไรดี". so we want to collect the data and analyze the appropriate menu for user.

## Features 🚀

- Get recommend menu based on user preferation by using data analysis

## Requirements 🔧

- go1.17+
- Node.js 14.18.1+
- npm 9.2.0
- Mobile Simulator (ex. Xcode for macOS user, Android Studio for Windows user)

## Build Steps (backend) 📦

### 1.Install depedencies

```bash
$ make deps
```

or

```bash
$ go mod download
```

### 2. Create the required config file

create **.env** file in the root directory

_example:_

```
# App config
ENV=development
GIN_MODE=release
PORT=8888

# Database config
MYSQL_HOSTNAME=
MYSQL_PORT=
MYSQL_USERNAME=
MYSQL_PASSWORD=
MYSQL_DATABASE=
```

### 3. Start backend server

```bash
$ make server
```

or

```bash
$ go run main.go
```

## Build Steps (frontend) 📦

### 1. Install depedencies

```bash
$ yarn install
```

or npm using

```bash
$ npm install
```

### 2. start frontend server

> iOS

```bash
yarn ios
```

or npm using

```bash
npm run ios
```

> Andriod

```bash
yarn andriod
```

or npm using

```bash
npm run android
```

## Swagger UI Documentation 📄

[Swagger](http://localhost:8888/swagger/index.html)

## Data Visualization 📊

[Data Visualization](https://kurester-visualization.vercel.app/)
