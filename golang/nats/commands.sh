#!/usr/bin/env bash

nats-server

nats pub msg.test "NATS MESSAGE 1"
nats pub msg.test.new "NATS MESSAGE 2"

nats sub "msg.test"

nats sub "msg.*.*"

nats sub "msg.>"

nats sub "msg.test" --dump "./received-messages"

nats subscribe --count 10 -s another-address:4222 topic

nats sub -s localhost:4222 --user=USER --password=PASSWORD

nats reply help.please 'OK, I CAN HELP!!!'
nats request help.please 'I need help!'
