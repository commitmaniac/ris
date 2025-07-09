# ris

Rename (files) in sequence

## Quick start

```console
$ make
$ ./ris -help
```

Static binaries are available on the [Releases] page

## Examples

> [!TIP]
> Use the `-simulate` flag to see how files will be renamed.
> Useful when deciding how files should be renamed.

Rename all text files inside the working directory

```console
$ ./ris -select "*.txt" .
```

Rename all text files to Markdown files while keeping the original filename

```console
$ ./ris -select "*.txt" -rename-to ".md" -keep-name .
```

## License

BSD-3-Clause

[Releases]: https://github.com/commitmaniac/ris/releases
