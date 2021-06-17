This is a very simple script programmed in Go that pretty-prints any json output piped to it, for instance from calling a REST API with ``curl``.

```bash
$ curl https://reqres.in/api/users/2 | prettyjson --indent 4
```