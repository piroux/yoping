#!/bin/bash

${SHELL} --version
set -x -u -o pipefail
# set -e
pwd

/app/restish --rsh-retry=0 post "http://api-backend:8855/ping/ex/+33600000001/+33600000003"

/app/restish --rsh-retry=0 get -H "Content-Type: application/json" "http://api-backend:8855/ping/in/+33600000001/+33600000003"


## -- CloudEvent JSON Payload
#cat <<EOF >input.json
#{
#    "specversion" : "1.0",
#    "type" : "dev.piroux.yoping.v1.ping.in",
#    "source" : null,
#    "subject": null,
#    "id" : "???",
#    "time" : "2025-04-06T17:31:00Z",
#    "datacontenttype" : "application/json",
#    "data" : {
#        "phone_from": "+33600000001",
#        "phone_to":   "+33600000003"
#    }
#}
#EOF
#
#/app/restish --rsh-retry=0 get -H "Content-Type: application/cloudevents+json" "http://api-backend:8855/ping/in/+33600000001/+33600000003" <input.json;


## -- Simple JSON Payload
cat <<EOF >input.json
{
    "phone_from": "+33600000001",
    "phone_to":   "+33600000003"
}
EOF
cat input.json

/app/restish --rsh-retry=0 get -H "Content-Type: application/json" "http://api-backend:8855/ping/in/+33600000001/+33600000003" <input.json;
