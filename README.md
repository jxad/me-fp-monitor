# Magic Eden Floor Price Monitor
## Contents
- [Contents](#contents)
- [Features](#features)
- [Installation](#installation)
- [Quick Start Single Collection Monitor](#quick-start-single-collection-monitor)
- [Quick Start Multiple Collection Monitor](#quick-start-multiple-collection-monitor) 
- [Sneak Peek](#sneak-peek)
- [Licence](#licence)

## Features
- Single collection monitoring
- Multiple collection monitoring

## Installation
```
git clone github.com/jxad/me-fp-monitor
```

```
go run .
```

```
go build
```

## Quick Start Single Collection Monitor
- Input collection name
- Select UP or DOWN (UP is going to check if the price go higher than the price inserted, DOWN is going to check if the price go lower)
- Input the price

## Quick Start Multiple Collection Monitor
- Edit config.json
```json
{
    "errorWebhook": "https://discord.com/api/webhooks/YOUR_WEBHOOK_HERE",
    "priceAlertWebhook": "https://discord.com/api/webhooks/YOUR_WEBHOOK_HERE",
    "magicEdenApiEndpoint": "https://api-mainnet.magiceden.io/v2"
}
```
- Edit collections.json
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

- Run mefpmonitor.exe
- Select Start Multiple Collection Monitor
- Select Start Monitor
- Input a delay
 
The delay its per collection. If you input 60 seconds. All collections are going to be checked every 60 seconds.
If you have 3 collections, the monitor is going to send a request every 20 seconds. [MagicEden Public API](https://api.magiceden.dev/) are ratelimited: default limit is 120 TPM or 2 QPS.

I recommend high delays so you can monitor many collections with no limitation.

If is going to be needed, proxy support can be added.

## Sneak Peek
### Single Collection Monitor
![single_collection_sneak_peek](https://user-images.githubusercontent.com/35408842/171857198-f302da22-fa29-4940-8fd3-f4091c5e1a39.gif)


### Discord Webhook Price Alert
![webhook_price_alert](https://user-images.githubusercontent.com/35408842/171857301-4a9a5a82-e7ee-4250-9a6e-2012e0c0ba6e.PNG)


## Licence
