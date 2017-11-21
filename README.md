# Zipkin Agent

[![Build Status](https://travis-ci.org/jcchavezs/zipkin-agent.svg?branch=master)](https://travis-ci.org/jcchavezs/zipkin-agent)

The aim of this agent is to do bulk reporting of spans over the transport.

The main use cases for this agent is:
- To reduce the load on the server on situations where there is direct reporting. For example
on PHP where the reporting should be done right after the request is served.
- To reduce the load on transports that act as middleware between the client and the server, for example Kafka.
- To act as a buffer when having limited or fragile connectivity to the server.

## Installation

### Binary

You can download the binary for every release on the [releases page](https://github.com/jcchavezs/zipkin-agent/releases)

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

the default transport is http but you can also define the transport:
```bash
TRANSPORT=logger zipkin-agent
```