# GoFish, The Package Manager

## Original author

The original author of this project is [Matt Fisher](mailto:matt.fisher@fishworks.io).
Due to the massive amount of time and money he had to spend on this project he decided to stop working on it and [archived](https://github.com/fishworks/gofish) it.

In the past I had already contributed some 'receipes' to the companion project [Fish Food](https://github.com/fishworks/fish-food) (which is, BTW, also archived), so I decided that I would try to keep the `GoFish` and its companion project still alive as long as possible.

[![Build Status](https://github.com/piccobit/gofish/actions/workflows/main.yaml/badge.svg)](https://github.com/piccobit/gofish/actions/workflows/main.yaml)
[![Release Build Status](https://github.com/piccobit/gofish/actions/workflows/release.yaml/badge.svg)](https://github.com/piccobit/gofish/actions/workflows/release.yaml)

## What does GoFish do?

GoFish is a cross-platform systems package manager, bringing the ease of use of Homebrew to
Linux and Windows.

```
$ gofish install go
==> Installing go...
🐠  go 1.10.1: installed in 2.307602197s
```

GoFish works across all three major operating systems (Windows, MacOS, and Linux). It installs
packages into its own directory and symlinks their files into `/usr/local` (or `C:\ProgramData` for Windows).
You can think of it as the cross-platform Homebrew.

## Want to add your project to the list of installable thingies?

Make a PR at [piccobit/fish-food](https://github.com/piccobit/fish-food)! Just make sure to follow the [Contributing Guide](https://gofi.sh#contributing) first.

## Troubleshooting

TBD

## Security

Please email security issues to [HD Stich](mailto:hd@stich.io).

## License

GoFish is licensed under the [Apache v2 License](LICENSE).
