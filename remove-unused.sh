#!/usr/bin/env bash

git rm -rf runsc shim tools webhook
git rm -rf pkg/shim
git rm -rf pkg/eventchannel
git rm -rf pkg/coverage
git rm -rf pkg/sentry
git rm -rf pkg/metric
git rm -rf pkg/hostos

git rm -rf pkg/ring0
git rm -rf pkg/prometheus

go mod tidy
go get -u all
