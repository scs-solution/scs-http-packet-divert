# SCS Http Packet Divert Tool

Http Header에 Uuid를 삽입한다.

삽입 대상

- 서버 어플리케이션 Inbound Packet 중 Uuid가 없는 Packet들
- 서버 어플리케이션의 Outbound Packet

- 모든 서버 어플리케이션의 Inbound Packet은 Tracing 서버로 전송됨

## init

```sh
sudo apt update
sudo apt install libnetfilter-queue-dev
sudo apt install pkg-config

sudo yum makecache
sudo yum install golang
sudo yum install libnetfilter_queue
sudo yum install -y libnetfilter_queue-devel

sudo iptables -L
sudo iptables -A INPUT -p tcp -j NFQUEUE --queue-num 0 --dport 80
setcap 'cap_net_admin=+ep' ./scs-http-packet-divert
```

## filter

```
This is a TCP packet!
From src port 57096 to dst port 80
PACKET LAYER: IPv4
IPv4    {Contents=[..20..] Payload=[..502..] Version=4 IHL=5 TOS=0 Length=522 Id=24117 Flags=DF FragOffset=0 TTL=106 Protocol=TCP Checksum=62393 SrcIP=~~~~ DstIP=~~~~ Options=[] Padding=[]}
00000000  45 00 02 0a 5e 35 40 00  6a 06 f3 b9 70 97 96 5e  |E...^5@.j...p..^|
00000010  ac 1f 09 ea                                       |....|

PACKET LAYER: TCP
TCP     {Contents=[..20..] Payload=[..482..] SrcPort=57096 DstPort=80(http) Seq=2332025059 Ack=1747083112 DataOffset=5 FIN=false SYN=false RST=false PSH=true ACK=true URG=false ECE=false CWR=false NS=false Window=1026 Checksum=3961 Urgent=0 Options=[] Padding=[]}
00000000  df 08 00 50 8a ff e0 e3  68 22 5f 68 50 18 04 02  |...P....h"_hP...|
00000010  0f 79 00 00                                       |.y..|

PACKET LAYER: Payload
Payload 482 byte(s)
00000000  47 45 54 20 2f 66 61 76  69 63 6f 6e 2e 69 63 6f  |GET /favicon.ico|
00000010  20 48 54 54 50 2f 31 2e  31 0d 0a 48 6f 73 74 3a  | HTTP/1.1..Host:|
00000020  20 77 77 77 2e 72 6f 6c  6c 72 61 74 2e 63 6f 6d  | www.rollrat.com|
00000030  0d 0a 43 6f 6e 6e 65 63  74 69 6f 6e 3a 20 6b 65  |..Connection: ke|
00000040  65 70 2d 61 6c 69 76 65  0d 0a 55 73 65 72 2d 41  |ep-alive..User-A|
00000050  67 65 6e 74 3a 20 4d 6f  7a 69 6c 6c 61 2f 35 2e  |gent: Mozilla/5.|
00000060  30 20 28 57 69 6e 64 6f  77 73 20 4e 54 20 31 30  |0 (Windows NT 10|
00000070  2e 30 3b 20 57 69 6e 36  34 3b 20 78 36 34 29 20  |.0; Win64; x64) |
00000080  41 70 70 6c 65 57 65 62  4b 69 74 2f 35 33 37 2e  |AppleWebKit/537.|
00000090  33 36 20 28 4b 48 54 4d  4c 2c 20 6c 69 6b 65 20  |36 (KHTML, like |
000000a0  47 65 63 6b 6f 29 20 43  68 72 6f 6d 65 2f 31 30  |Gecko) Chrome/10|
000000b0  36 2e 30 2e 30 2e 30 20  53 61 66 61 72 69 2f 35  |6.0.0.0 Safari/5|
000000c0  33 37 2e 33 36 0d 0a 41  63 63 65 70 74 3a 20 69  |37.36..Accept: i|
000000d0  6d 61 67 65 2f 61 76 69  66 2c 69 6d 61 67 65 2f  |mage/avif,image/|
000000e0  77 65 62 70 2c 69 6d 61  67 65 2f 61 70 6e 67 2c  |webp,image/apng,|
000000f0  69 6d 61 67 65 2f 73 76  67 2b 78 6d 6c 2c 69 6d  |image/svg+xml,im|
00000100  61 67 65 2f 2a 2c 2a 2f  2a 3b 71 3d 30 2e 38 0d  |age/*,*/*;q=0.8.|
00000110  0a 52 65 66 65 72 65 72  3a 20 68 74 74 70 3a 2f  |.Referer: http:/|
00000120  2f 77 77 77 2e 72 6f 6c  6c 72 61 74 2e 63 6f 6d  |/www.rollrat.com|
00000130  2f 0d 0a 41 63 63 65 70  74 2d 45 6e 63 6f 64 69  |/..Accept-Encodi|
00000140  6e 67 3a 20 67 7a 69 70  2c 20 64 65 66 6c 61 74  |ng: gzip, deflat|
00000150  65 0d 0a 41 63 63 65 70  74 2d 4c 61 6e 67 75 61  |e..Accept-Langua|
00000160  67 65 3a 20 6b 6f 2d 4b  52 2c 6b 6f 3b 71 3d 30  |ge: ko-KR,ko;q=0|
00000170  2e 39 2c 65 6e 2d 55 53  3b 71 3d 30 2e 38 2c 65  |.9,en-US;q=0.8,e|
00000180  6e 3b 71 3d 30 2e 37 0d  0a 49 66 2d 4e 6f 6e 65  |n;q=0.7..If-None|
00000190  2d 4d 61 74 63 68 3a 20  57 2f 22 36 35 34 62 2d  |-Match: W/"654b-|
000001a0  31 38 33 34 35 31 39 35  36 38 30 22 0d 0a 49 66  |18345195680"..If|
000001b0  2d 4d 6f 64 69 66 69 65  64 2d 53 69 6e 63 65 3a  |-Modified-Since:|
000001c0  20 46 72 69 2c 20 31 36  20 53 65 70 20 32 30 32  | Fri, 16 Sep 202|
000001d0  32 20 30 37 3a 30 30 3a  33 32 20 47 4d 54 0d 0a  |2 07:00:32 GMT..|
000001e0  0d 0a                                             |..|
```

## ref

nfqueue

- https://colorme.tistory.com/4
- https://github.com/chifflier/nfqueue-go
