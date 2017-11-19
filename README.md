# Zipkin Agent

[![Build Status](https://travis-ci.org/jcchavezs/zipkin-agent.svg?branch=master)](https://travis-ci.org/jcchavezs/zipkin-agent)

This is an agent that should run in the same host of the application in order to reduce the load
on the zipkin server.

## Installation

### Golang
```bash
go get github.com/jcchavezs/zipkin-agent
go install github.com/jcchavezs/zipkin-agent/cmd/zipkin-agent
```

## Binary distribution
Each release provides pre-built binaries for different architectures, you can download them here:
https://github.com/jcchavezs/zipkin-agent/releases/latest

## Getting started

You can run the agent by:
```bash
zipkin-agent
```

the default transport is the standard stdout but you can also define the transport:
```bash
TRANSPORT=http zipkin-agent
```