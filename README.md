# Telegraf Input Plugin P2KSP


## Description

Reads the content of UltimaULG.txt along with its modification timestamp and tags with the retail store number.

## Build

Clone the Telegraf project and include the address at plugins/inputs/all/all.go.

Follow the instructions to build a Telegraf binary.

## Issues

It should work on Windows and Linux but I hadnt tested on Linux and Windows in case there are multiple sp_lj* instances or in Linux install where there is an sp_lj9999 symlink.
