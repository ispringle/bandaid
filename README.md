# Bandaid

A tool for triggering events based of Prometheus exporters on localhost.
It will also scrape remote, which can be configured with flags, but I
intend for this only to be for testing purposes.

## Usage (intended)

Bandaid is intended to be a stopgap for bad alerts that require brainless
reactions in order to resolve. Bandaid will scrape `localhost` and if an
alert matches the bandaid config (tbd) it will trigger an event. The event
will be standalone from bandaid.

Like all bandaids, this one is gross and I hope you never have to use it.

### bandaid.conf (wip)

```yaml
name: bad_alert
  trigger: >
   node_systemd_unit_state{name="trash_tier_service.service",state="active"} == 0
  event: >
   systemctl restart trash_tier_service.service
name: anoter_one
  trigger: >
   node_systemd_unit_state{name="technical_debt.service",state="active"} == 0
  event: >
   sh /path/to/further_technical_debt.sh
```

## TODO
- [ ] get parsing working
- [ ] determine a format for configs
