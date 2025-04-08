
# TODO


## Telemetry
- Setup in project
- Add to all layers
- Setup LGTM stack
- Collect

## Domain

- Handle ErrDataNotFound

- Upgrade models.Ping.TimeCreated:
```
TimeRxed  time.Time
TimeTxed  time.Time
```

## Domain Tests / Validation
- Check:
  - Phones From and tT are different
  - Phone From is valid and exists
  - Phone To is valid and exists

## Persistence

- Transaction for repos:
  - passed as optional arg
  - embeded transaction ?

- Remove uuid.MustParse
- Add context
- Timezone not handled by Time.Scan
- Repo: refactor: Use fonction to create model from pers types
- Add ñock for test with https://github.com/objectbox/objectbox-go
