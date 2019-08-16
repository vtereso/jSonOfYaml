# yson
`yson` is a utility wrapper around [`github.com/ghodss/yaml`](https://github.com/ghodss/yaml) to easily convert yaml files into json. By default, `yson` prints to stdout.

```
yson <INPUT_FILE> [-o <OUTPUT_FILE>]
`INPUT_FILE`    - Specify `-` as `<INPUT_FILE>` for stdin
`-o`            - File to store conversion. Will create/overwrite.
```