# Magic Eden Floor Price Monitor
## Contents
- [Contents](#contents)
- [Features](#features)
- [Installation](#installation)
- [Quick Start Single Collection Monitor](#quick-start-single-collection-monitor)
- [Quick Start Multiple Collection Monitor](#quick-start-multiple-collection-monitor)

## Features
- Single collection monitoring
- Multiple collection monitoring

## Installation
TODO

## Quick Start Single Collection Monitor
TODO

## Quick Start Multiple Collection Monitor
- run mefpmonitor.exe
- edit config.json
```json
{
    "errorWebhook": "https://discord.com/api/webhooks/YOUR_WEBHOOK_HERE",
    "priceAlertWebhook": "https://discord.com/api/webhooks/YOUR_WEBHOOK_HERE",
    "magicEdenApiEndpoint": "https://api-mainnet.magiceden.io/v2"
}
```
- edit collections.json
```json
[
    {
        "symbol": "enviro",
        "price": "7",
        "upordown": "UP"
    },
    {
        "symbol": "okay_bears",
        "price": "150",
        "upordown": "DOWN"
    }
]
```
Collections can be also edited from the CLI
