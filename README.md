# nscr - CLI running npm scripts with interactive filter

## Usage

```sh
$ nscr
> 
  dev    webpack-dev-server --inline
  build  webpack --config prod.config.js
```

## Requirement

- An interactive filter
  - [fzf](https://github.com/junegunn/fzf)
  - [peco](https://github.com/peco/peco)
  - [fzy](https://github.com/jhawthorn/fzy)
  - [percol](https://github.com/mooz/percol)
  - ...

## Installation

```sh
$ go get github.com/micnncim/nscr
```

## Note

Maybe `npm run` or `yarn run` are more useful...😢

## LICENSE

[MIT](./LICENSE)
