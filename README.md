# go-necos
### Simple Nekos API wrapper written in Go

![Go](https://img.shields.io/badge/Go-00ADD8?style=plastic&logo=go&logoColor=white) ![tests](https://github.com/rinnothing/go-necos/actions/workflows/go.yml/badge.svg)

Documentation for the API can be found by the links below:

[nekosapi.com/docs](https://nekosapi.com/docs) | [api.nekosapi.com/v3/docs](https://api.nekosapi.com/v3/docs)

## Usage
The wrapper provides you with a [Client](api.go#L17) structure that has DefaultQuery field for setting the default query,
which will be [merged](api.go#L48) with every request from that Client, and Domain for setting base domain for all API calls.

All methods to interact with API (except 'Get a random image file redirect', which is not implemented) have the same names
as in the documentation (with adding Get or Post prefixes here and there). 

Mandatory arguments for calls are in the signature
of methods and optional arguments are given through Request structure (simply a url.Values alias). 
What fields can be used by a call is stated in API docs or in the method comment.

To simplify the creation of Requests the wrapper gives you [SetFields and AddFields](sugar.go#L34) functions that can
set and add new fields to already existing Request or make a new one. Odd arguments are responsible for keys and even for values.

Also, since the API goal is to provide you with pictures, wrapper simplifies it here as well.
You can use [Save, SaveTemp and SaveSlice](sugar.go#L106) functions to create writers to file by given path, file in temp directory or to slice respectfully.
The latter writers can be used by methods like [DownloadImage, DownloadSample and Download](sugar.go#L171) to download images into them.

Examples of usage can be found in tests and in [examples](examples)
