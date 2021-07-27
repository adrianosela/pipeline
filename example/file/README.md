### Sample Output

```
11:05 $ go run main.go
2021/07/27 11:05:18 [SINK:<writeRepoName>] Starting...
2021/07/27 11:05:18 [SOURCE:<readRepoURL>] Starting...
2021/07/27 11:05:18 [STAGE:<trimURLPrefix>] Starting...
2021/07/27 11:05:18 [STAGE:<trimRepoOrgPrefix>] Starting...
2021/07/27 11:05:18 [SOURCE:<readRepoURL-0>] Source finished. Thread terminating...
2021/07/27 11:05:18 [SOURCE:<readRepoURL>] All threads terminated. Closing output channel...
2021/07/27 11:05:18 [STAGE:<trimURLPrefix-4>] Input channel closed. Thread terminating...
2021/07/27 11:05:18 [STAGE:<trimURLPrefix-3>] Input channel closed. Thread terminating...
2021/07/27 11:05:18 [STAGE:<trimURLPrefix-0>] Input channel closed. Thread terminating...
2021/07/27 11:05:18 [STAGE:<trimURLPrefix-1>] Input channel closed. Thread terminating...
2021/07/27 11:05:18 [STAGE:<trimURLPrefix-2>] Input channel closed. Thread terminating...
2021/07/27 11:05:18 [STAGE:<trimURLPrefix>] All threads terminated. Closing output channel...
2021/07/27 11:05:18 [STAGE:<trimRepoOrgPrefix-4>] Input channel closed. Thread terminating...
2021/07/27 11:05:18 [STAGE:<trimRepoOrgPrefix-3>] Input channel closed. Thread terminating...
2021/07/27 11:05:18 [STAGE:<trimRepoOrgPrefix-0>] Input channel closed. Thread terminating...
2021/07/27 11:05:18 [STAGE:<trimRepoOrgPrefix-1>] Input channel closed. Thread terminating...
2021/07/27 11:05:18 [STAGE:<trimRepoOrgPrefix-2>] Input channel closed. Thread terminating...
2021/07/27 11:05:18 [STAGE:<trimRepoOrgPrefix>] All threads terminated. Closing output channel...
2021/07/27 11:05:18 [SINK:<writeRepoName-1>] Input channel closed. Quitting...
2021/07/27 11:05:18 [SINK:<writeRepoName-3>] Input channel closed. Quitting...
2021/07/27 11:05:18 [SINK:<writeRepoName-0>] Input channel closed. Quitting...
2021/07/27 11:05:18 [SINK:<writeRepoName-2>] Input channel closed. Quitting...
2021/07/27 11:05:18 [SINK:<writeRepoName-4>] Input channel closed. Quitting...
2021/07/27 11:05:18 [SINK:<writeRepoName>] All threads terminated. Closing output channel...
```

```
11:05 $ cat out.txt
spoof
resume
multikey
interview
iprepd-firewall
ecdh
vpn
sslmgr
middleman
rdtp
padl
```