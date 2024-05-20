<br />
  <div align="center">
    <img src="https://github.com/yohansky/Fe-Blanja-React/assets/69236028/d17e913b-a889-40c1-b745-46193b16a9ba"/>
  <br />
  <h1>Blanja</h1>
    <a href="https://github.com/yohansky/Fe-Blanja-React">View Demo</a>
    .
    <a href="https://github.com/yohansky/Be-Blanja-fiber">Api Demo</a>
  </div>

  ## Table of Contents

- [Table of Contents](#table-of-contents)
- [About The Project](#about-the-project)
  - [Built With](#built-with)
- [Installation](#installation)
  - [Documentation](#documentation)
  - [Related Project](#related-project)
 
## About The Project

The Blanja e-commerce website project is an online platform designed to make it easier for users to make buying and selling transactions online. This website was developed using React.js and Redux frontend technologies to optimize display performance and efficiency as well as state management. On the backend, this website uses Golang with Fiber Framework technology which functions as a server to manage data taken from Cloudinary. Cloudinary functions as an image data store so that it allows users to easily view images of each product being sold.

### Built With
These are the libraries and service used for building this backend API

- [Golang](https://go.dev/)
- [Fiber](https://gofiber.io/)
- [PostgreSQL](https://www.postgresql.org/)
- [Json Web Token](https://jwt.io/)
- [Gorm.io](https://gorm.io/index.html)

## Installation
1. Clone this repository

```sh
git clone https://github.com/yohansky/Be-Blanja-fiber
```

2. Change directory to Be_Blanja

```sh
cd Be-Blanja-fiber
```

3. Install all of the required modules

```sh
go mod tidy
```

4. Create PostgreSQL database, query are provided in [query.sql](./query.sql)

5. Create and configure `.env` file in the root directory, example credentials are provided in [.env.example](./.env.example)

6. Run this command to run the server

```sh
air
```

- Run this command for debugging and finding errors

```sh
golangci-lint run
```

### Documentation

- [Postman API colletion]()
- [PostgreSQL database query](./query.sql)

API endpoint list are also available as published postman documentation

### Related Project
:rocket: [`Backend Blanja`](https://github.com/yohansky/Be-Blanja-fiber)

:rocket: [`Frontend Blanja`](https://github.com/yohansky/Fe-Blanja-React)
