FROM docker.io/ipfs/kubo:latest

ENV IPFS_PATH=/data/ipfs

# 5001 API
# 8080 Gateway
# 4001 P2P Swarm
EXPOSE 4001 5001 8080

# Initialize the repository on first run
# If /data/ipfs has no "config", initialize it
ENTRYPOINT ["/usr/local/bin/start_ipfs"]
