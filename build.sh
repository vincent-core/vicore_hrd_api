#!/bin/bash
echo "pulling latest update from git"
git pull

echo "setup golang ..."
export PATH=$PATH:/usr/local/go/bin

ceho "Clear systectl log ... "
journalctl --vacuum-time=1d

echo "building..."
build -o hrd-api

echo "restarting the service"
systemctl restart rme.service

echo "service status"
systemctl status rme.service

echo "done."