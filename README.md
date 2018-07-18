# container-scan

Scans docker images and containers for security purposes. 

## Reasoning

Could not find a decent scanning tool for malicous software on container images. Existing systems assumed the container was in a running state where a malicious service could use various methods to hide itself. Other scan tools were simply for CVEs and assume the state of the system itself is proper.

## Technology Decisions 

* Golang for the language as it allows for easy access to docker client calls without process calling
* [urfave/cli](https://github.com/urfave/cli) for the CLI based on readability and performance

## Setup

Currently this is at the CLI build phase so there's nothing backend at the moment. That said the standard process applies here of just doing a `go build` and running the resulting `container-scan` binary.
