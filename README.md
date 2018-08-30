# Nomenclature Utilities

This project contains utilities for checking device name availability and selection.

## Device Name Format

The name of the device will follow this format:

```{SERVICE}-{FUNCTION}-{ID}.{DOMAIN}```

where

| Main key             | Index    |
|:-------------------- |:--------:|
|`{SERVICE}-{FUNCTION}`| `{ID}`   |

Each element is determined as described in the ensuing sections. The main key is used to look up corresponding devices of the same group to determine the next device index to attribute to a new device in that group.

_Note that the device name with the domain name removed must be unique!_

    Example: inf-as-10.domain.local

## Choice of Service

 * Client identified by account ID (integer)
 * Internal: Office (`off`)
 * Internal: Infrastructure (`inf`)

The selection will form the initial part of the hostname. Internal use-cases are identified by these keywords: `web`, `mail`, `voip`, `dia`, `amb`

## Choice of Function

Physical Device: `hv`, `sd`, `pw`, `mm`

Virtual Device:

| Service       | Description       |
|:------------- |:------------------|
|ml             | Mail              |
|se             | Security          |
|db             | Database          |
|px             | Proxy             |
|ad             | Authentication    |
|fw             | Firewall          |
|fs             | Fileserver        |
|as             | Application Server|
|vo             | VOIP              |
|mn             | Monitoring        |
|bk             | Backup            |
|ns             | Nameserver        |
|cf             | Configuration     |
|mq             | Messaging         |
|wb             | Web               |

## Device ID

The ID of the device is a numerical ID which should increment by one for each new device where the main key is the same.