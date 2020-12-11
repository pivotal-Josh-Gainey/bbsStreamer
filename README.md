# bbs_stream

This binary can be used to subscribe to bbs events and print to screen for debugging. 

BBS uses mutual TLS so the CA/cert/key is required as well as bbs.service.cf.internal be resolvable on the machine it is ran on.

```Usage: ./bbs_stream <ca-crt> <client-crt> <client-key>```
