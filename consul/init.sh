#!/bin/sh
#sh /usr/local/bin/docker-entrypoint.sh "agent" "-dev" "-client" "0.0.0.0" &
#sleep 10;
echo "Initialize rms-api config environtment"
curl --output /dev/null -sX PUT --data-binary @/consul/config/rms-api.dev.yaml consul:8500/v1/kv/rms-api/dev &
curl --output /dev/null -sX PUT --data-binary @/consul/config/rms-api.prod.yaml consul:8500/v1/kv/rms-api/prod