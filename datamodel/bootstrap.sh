dropdb telemetry
createdb telemetry
find . -name '*.sql' -exec psql telemetry -f {} \;
