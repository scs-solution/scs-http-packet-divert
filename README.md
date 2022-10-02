# SCS Http Packet Divert Tool

## init

```sh
sudo apt update
sudo apt install libnetfilter-queue-dev
sudo apt install pkg-config

sudo iptables -A OUTPUT -p icmp -j NFQUEUE --queue-num 0
setcap 'cap_net_admin=+ep' ./scs-http-packet-divert
```

## iptables

nfqueue

- https://colorme.tistory.com/4
- https://github.com/chifflier/nfqueue-go
