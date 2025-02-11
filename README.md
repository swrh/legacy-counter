# Legacy Counter

A nostalgic recreation of the classic counter.digits.com service from 2004, which served as a simple webpage hit counter.

## About

Legacy Counter is a modern implementation inspired by the beloved counter.digits.com service that was widely used in the early 2000s to track website visitor counts. This project aims to bring back that classic functionality while maintaining the retro aesthetic that made the original service so memorable.

## Features

- Simple hit counting functionality
- Classic digit display styles
- Easy integration
- Lightweight and fast

## Current Limitations

The current implementation supports only a single global counter stored in a plain text file. Future versions will likely migrate to a database solution to support multiple independent counters for different websites.

## Technical Details

The counter value is persisted in a simple text file (`counter.txt`) that gets incremented with each request. While this approach works for demonstration purposes, it's not suitable for production use with multiple counters. A future database implementation will allow for:

- Multiple independent counters
- Counter ownership and management
- Better concurrency handling
- Improved scalability
