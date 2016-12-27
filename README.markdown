pqinterval
==========

Scan targets for PostgreSQL's interval type.

The `Interval` type supports the full range of PostgreSQL intervals. Any
database interval should successfully `Scan()` into a `pqinterval.Interval`.

Example:

```golang
var ival pqinterval.Interval

err := conn.QueryRow("SELECT '4 days'::INTERVAL").Scan(&ival)
if err != nil {
    log.Fatal(err)
}

fmt.Println(ival.Hours())
```

The `Duration` type is an alias of `time.Duration`, but which supports
scanning from PostgreSQL intervals (potentially failing with `ErrTooBig`).

Example:

```golang
var since time.Duration
d := (*pqinterval.Duration)(&since)

err = conn.QueryRow("SELECT '2 days'::INTERVAL").Scan(d)
if err != nil {
    log.Fatal(err)
}

fmt.Println(since)
```
