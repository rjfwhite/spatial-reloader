# spatial-reloader
Simple tool for [SpatialOS](spatialos.com) to aid in local development. Scans the `build/assembly/workers` folder for changes, then restarts all connected workers within a local deployment when any file changes.

Especially for simple workers, this can drastically reduce iteration times.

## Dependencies
https://github.com/Jeffail/gabs

## Usage
This assumes you have your golang `bin` folder on your `PATH`. Runs from the base of your SpatialOS project, and runs indefinitely.

```bash
go install github.com/rjfwhite/spatial-reloader
cd your-spatial-project
spatial-reloader
```

## Potential Extensions

* Invoke `spatial local` too, restarting it when any snapshot or configuration changes
