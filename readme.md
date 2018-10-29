# darkraiden/switch

**NOTE**: This project is still **work in progress**!

Switch allows you to create a small, lightweight proxy by declaring a list of rules in `YAML` format.

- [darkraiden/switch](#darkraidenswitch)
  - [Install](#install)
  - [Examples](#examples)
  - [TODO's](#todos)
  - [License](#license)

## Install

With a correctly configured Go toolchain:

```bash
go get -u github.com/darkraiden/switch
```

## Examples

Create your own `routes.yaml` file and make sure it follows the structure of the one below:

```yaml
paths:
    - source: /google
      destination: https://google.com
    - source: /facebook
      destination: https://facebook.com
    - source: /
      destination: https://www.darkraiden.com/
```

Load the routes file using the `filename` flag:

```bash
switch -filename yourCustomRoutesFile.yaml
```

A GET request to `https://yourdomain.com:3000/google` will cause a redirect to `https://google.com` with a 302 Status Code.

## TODO's

-   [ ] Add Tests
-   [ ] Add Redirection type to YAML structure

## License

BSD licensed. See the LICENSE file for details.
