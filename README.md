# GeoFence - A TCP authorization Proxy

GeoFence is a fast TCP reverse-proxy to block countries from connecting to your service.  Written to whitelist players in my community accessing my minecraft server while keeping out the bots.

## Prerequisites
To determine the country of an address, GeoFence uses the GeoLite2 IP Database.  Sign over your firstborn to Maxmind to receive a copy of their binary database.

## Configuration
```yaml
db: /var/db/geolite2-country.mmdb

report:
  token: "<discord-bot-token>" 
  channel: "<discord-channel-id>" 

services:
- name: Minecraft
  listen: 0.0.0.0:25565
  upstream: hypixel.net:25565
  rules:
  - if: ip != "127.0.0.1"
  - if: city(ip) == "Rotterdam"
  - if: country(ip) == "nl"
```
