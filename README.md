# Rate Limit Service

This project is a minimal service written in Go that is designed to track compliance with predefined rules. At the current stage, the service implements only time-based rules to provide basic functionality for monitoring user connections.

## Overview

The primary goal of this service is to evaluate whether user actions comply with a set of rules. For now, the scope is intentionally limited to time-based rules, serving as a foundation for future expansion to additional rule types.

## Features

- Implemented in Go
- Minimal and lightweight design
- Tracks compliance with rules
- Supports time-based rules for user connections
- Extensible architecture for future rule types

## Current Limitations

- Only time-based rules are supported
- No advanced rule composition or persistence mechanisms implemented yet
- Intended primarily as a proof of concept / minimal viable implementation

## Use Cases

- Tracking how frequently a user connects within a defined time window
- Enforcing simple rate or time-based access rules
- Serving as a base for more complex rule-compliance or rate-limiting systems

## Future Improvements

- Add support for additional rule types (e.g. quota-based, event-based)
- Introduce persistent storage (e.g. Redis, PostgreSQL)
- Improve configuration and rule definition
- Add observability (logging, metrics, tracing)

## Requirements

- Go 1.20 or newer
