# ipfs
`export ipfs_data="/Users/aaron/data_ipfs"`

`export ipfs_staging="/Users/aaron/data"`

`docker run -d --name ipfs_host -v $ipfs_staging:/export -v $ipfs_data:/data/ipfs -p 4001:4001 -p 8080:8080 -p 5001:5001 ipfs/go-ipfs:latest`

通过`docker logs -f ipfs_host` 查看日志

```shell
Changing user to ipfs
ipfs version 0.6.0-rc7
initializing IPFS node at /data/ipfs
generating 2048-bit RSA keypair...done
peer identity: QmXsxFFwDZYTXNMHLR5cH2AY1esLLSBm8N8CdEwgYKmgUk
to get started, enter:

	ipfs cat /ipfs/QmQPeNsJPyVWPFDVHb77w8G42Fvo15z4bG2X8D2GhfbSXc/readme

Initializing daemon...
go-ipfs version: 0.6.0-rc7-63ff04c
Repo version: 10
System version: amd64/linux
Golang version: go1.14.2
Swarm listening on /ip4/127.0.0.1/tcp/4001
Swarm listening on /ip4/127.0.0.1/udp/4001/quic
Swarm listening on /ip4/172.17.0.2/tcp/4001
Swarm listening on /ip4/172.17.0.2/udp/4001/quic
Swarm listening on /p2p-circuit
Swarm announcing /ip4/127.0.0.1/tcp/4001
Swarm announcing /ip4/127.0.0.1/udp/4001/quic
Swarm announcing /ip4/172.17.0.2/tcp/4001
Swarm announcing /ip4/172.17.0.2/udp/4001/quic
Swarm announcing /ip6/::1/tcp/4001
Swarm announcing /ip6/::1/udp/4001/quic
API server listening on /ip4/0.0.0.0/tcp/5001
WebUI: http://0.0.0.0:5001/webui
Gateway (readonly) server listening on /ip4/0.0.0.0/tcp/8080
Daemon is ready
```

执行上面的命令

```
docker exec ipfs_host ipfs cat /ipfs/QmQPeNsJPyVWPFDVHb77w8G42Fvo15z4bG2X8D2GhfbSXc/readme
```

查看对等节点

```
docker exec ipfs_host ipfs swarm peers
```

将文件放入关联到docker目录并添加到ipfs中

```
# 将一个文件添加到宿主机目录  -r  recursion
cp docker-compose.yml $ipfs_staging
# 添加到ipfs中
docker exec ipfs_host ipfs add /export/docker-compose.yml
```

通过 `ipfs cat <CID>` 查看文件内容

```
docker exec ipfs_host ipfs cat QmbJniEGEvCF54REqzCguU8MA4QUa3LAAN8VLkzpLQTjp8
```

可以验证修改文件名等操作，并不会改变存储内容的的文件名（标识信息），如果我们想获取追踪这些信息，通过增加Warpped信息来标识

（以下操作将`docker exec ipfs_host ipfs` alias 成 ipfs` 方便在宿主机操作）

```
ipfs add -w /export/docker-compose.yml
```

这里会返回两个CID，第一个可以看到与之前的相同，当我们通过`ipfs ls -v <cid>`来查看时，显示`Hash` 和 `Size Name` 均为空，但当我们通过返回的第二个CID来查看以后，则提供了更多的信息：

```
$ ipfs ls -v QmeuBEv4Nt1mxbCb3wqW6G4gnoopSg1JdnDurNxe4Td19z
$	Hash                                           Size Name
	QmbJniEGEvCF54REqzCguU8MA4QUa3LAAN8VLkzpLQTjp8 662  docker-compose.yml
```

注意到这里我们将其标识为两组信息，并不能简单用`cat` 来查看文件内容

```
$ ipfs cat QmeuBEv4Nt1mxbCb3wqW6G4gnoopSg1JdnDurNxe4Td19z
$ Error: this dag node is a directory
```

而是要使用如下方式来获取文件内容

```shell
 ipfs cat QmeuBEv4Nt1mxbCb3wqW6G4gnoopSg1JdnDurNxe4Td19z/docker-compose.yml
```

将文件发布到网络

```
$ ipfs name publish QmbJniEGEvCF54REqzCguU8MA4QUa3LAAN8VLkzpLQTjp8
$ Published to QmXsxFFwDZYTXNMHLR5cH2AY1esLLSBm8N8CdEwgYKmgUk: /ipfs/QmbJniEGEvCF54REqzCguU8MA4QUa3LAAN8VLkzpLQTjp8
```

