# Iris

[![License Apache 2][badge-license]](LICENSE)
[![GitHub version](https://badge.fury.io/gh/nlamirault%2Firis.svg)](https://badge.fury.io/gh/nlamirault%2Firis)

Master :
* [![Circle CI](https://circleci.com/gh/nlamirault/iris/tree/master.svg?style=svg)](https://circleci.com/gh/nlamirault/iris/tree/master)

Develop :
* [![Circle CI](https://circleci.com/gh/nlamirault/iris/tree/develop.svg?style=svg)](https://circleci.com/gh/nlamirault/iris/tree/develop)


Iris is a [cadvisor][] client to monitor [Docker][] containers.

## Usage

Launch the web service :

	$ iris -d


## Development

* Initialize environment

        $ make init

* Build tool :

        $ make build

* Launch unit tests :

        $ make test

## Contributing

See [CONTRIBUTING](CONTRIBUTING.md).


## License

See [LICENSE](LICENSE) for the complete license.


## Changelog

A [changelog](ChangeLog.md) is available


## Contact

Nicolas Lamirault <nicolas.lamirault@gmail.com>



[badge-license]: https://img.shields.io/badge/license-Apache2-green.svg?style=flat


[Docker]: https://docker.com
[cadvisor]: https://github.com/google/cadvisor
