# protomq

* install protoc
## php


* https://github.com/protocolbuffers/protobuf/tree/master/php

```bash
composer require google/protobuf
composer require spiral/roadrunner 
```
## go

```bash
go get github.com/spiral/roadrunner
go get -u github.com/golang/protobuf/protoc-gen-go
```

## notes

```bash
panic: [5] Leader Not Available: the cluster is in the middle of a leadership election and there is currently no leader for this partition and hence it is unavailable for writes
```