<p align="center">
<a href="https://saweria.co/ardihikaru">
<img src="https://mfardiansyah.id/assets/images/saweria.png" width="30%" />
</a>
<br>
Love my work? Drop me a coffee here. :)
</p>

<p align="center">
<a href="#">
<img src="https://img.shields.io/badge/%20Platforms-Windows%20/%20Linux-blue.svg?style=flat-square" alt="Platforms" />
</a>
<img src="https://img.shields.io/badge/%20Licence-MIT-green.svg?style=flat-square" alt="license" />
</p>
<p align="center">
<a href="https://github.com/ardihikaru/go-chi-example-part-1/blob/master/CODE_OF_CONDUCT.md">
<img src="https://img.shields.io/badge/Community-Code%20of%20Conduct-orange.svg?style=flat-squre" alt="Code of Conduct" />
</a>
<a href="https://github.com/ardihikaru/go-chi-example-part-1/blob/master/SUPPORT.md">
<img src="https://img.shields.io/badge/Community-Support-red.svg?style=flat-square" alt="Support" />
</a>
<a href="https://github.com/ardihikaru/go-chi-example-part-1/blob/master/CONTRIBUTING.md">
<img src="https://img.shields.io/badge/%20Community-Contribution-yellow.svg?style=flat-square" alt="Contribution" />
</a>
</p>
<hr>

# Api Service with Swagger built with Go-Chi framework

Global Template Repository for Development and Operations Of Your Projects.

| Key               | Values                                                                                  |
|-------------------|-----------------------------------------------------------------------------------------|
| Author            | Muhammad Febrian Ardiansyah                                                             |
| Email             | mfardiansyah.id@gmail.com                                                               |
| LinkedIn          | [Muhammad Febrian Ardiansyah](https://www.linkedin.com/in/muhammad-febrian-ardiansyah/) |
| Personal Homepage | [https://mfardiansyah.id](https://mfardiansyah.id/)                                     |

## Table of Contents

* [Dependencies](#dependencies)
* [Prerequisites](#prerequisites)
* [Installation](#installation)
* [Development](#development)
* [Usage](#usage)
* [Swag Usage](#Swag-Usage)
* [Contributing](#contributing)
* [License](#license)

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See deployment for notes on how to deploy the project on a live system.

### Prerequisites

What things you need to install the software and how to install them

```shell
apt-get -y install git
```

Or

```shell
yum -y install git
```

### Installation

A step by step series of examples that tell you how to get a development env running

Say what the step will be clone this repository.

```shell
git clone git@github.com:ardihikaru/go-chi-example-part-1.git
```

## Development

- N/A

## Usage

Reference and programming instructional materials.

## Swag Usage

- Add comments on your API handlers using the declarative syntax explained [here](https://swaggo.github.io/swaggo.io/declarative_comments_format/)

-  Install [swaggo/swag](https://github.com/swaggo/swag)?

```
        go install github.com/swaggo/swag/cmd/swag@latest
```

- Run the Swag in your Go project root folder which contains `main.go` file.

``` 
        swag init -g cmd/main/main.go
        swag init --parseDependency -g cmd/api/main.go
```

If your `main.go` file is not in root but uses the models defined in root, you can provide the path of `main.go` file.

```
        swag init -d "./" -g "$FOLDER_NAME/main.go"
```

Swag will parse comments and generate required files(docs folder and docs/doc.go).

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

Looking to contribute to our code but need some help? There's a few ways to get information:

* Connect with me on [Twitter](https://twitter.com/ardikucing)
* Connect with me on [Facebook](https://facebook.com/ardihikaru)
* Connect with me on [LinkedIn](https://linkedin.com/in/muhammad-febrian-ardiansyah)
* Log an issue here on github

## Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available, see the [tags on this repository](https://github.com/ardihikaru/go-chi-example-part-1/tags).

## Authors

* **[Muhammad Febrian Ardiansyah](https://github.com/ardihikaru)** - *Initial work*

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details

<p> Copyright &copy; 2024 Public Use. All Rights Reserved.

## Acknowledgments

* Hat tip to anyone whose code was used
* Inspiration
* etc

## MISC

* Validates CORS
```shell
curl -v --request OPTIONS 'http://localhost:8080/public/service-id' -H 'Origin: http://other-domain.com' -H 'Access-Control-Request-Method: GET'
```
```shell
curl -v --request OPTIONS 'http://localhost:8080/auth/login' -H 'Origin: http://other-domain.com' -H 'Access-Control-Request-Method: POST'
```
```shell
curl -v -X OPTIONS \
  http://localhost:8080/public/service-id \
  -H 'cache-control: no-cache' \
  -F Origin=http://www.google.com
```
  * **Allowed** CORS result (please set `cors.Debug: true`)
    ```shell
    [cors] 2024/06/16 23:53:13 Handler: Preflight request
    [cors] 2024/06/16 23:53:13 Preflight response headers: map[Access-Control-Allow-Methods:[GET] Access-Control-Allow-Origin:[http://other-domain.com] Access-Control-Max-Age:[6000] Vary:[Origin Access-Control-Request-Method Access-Control-Request-Headers]]
    ```
  * **NOT Allowed** CORS result (please set `cors.Debug: true`)
    ```shell
    [cors] 2024/06/16 23:52:13 Handler: Preflight request
    [cors] 2024/06/16 23:52:13 Preflight aborted: origin 'http://other-domain.com' not allowed
    ```
