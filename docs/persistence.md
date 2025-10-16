# Persistence in a KV Store

In the other documentation ([kvstore.md](kvstore.md)), there are two types of key-value stores mentioned: in-memory and persistent. I decided to implement persistence in my key-value store, and allowed users to choose between in-memory and persistent storage when creating a new store (by using flags, like a CLI).

## Why Persistence?
The main reason is that I thought it would be cool and obviously educational. I also saw that [Redis](https://redis.io/docs/latest/operate/oss_and_stack/management/persistence/) supports both in-memory and persistent storage, so I wanted to try implementing it myself.

I realized that Redis does persistence in three ways:
1. RDB (Redis Database)
2. AOF (Append Only File)
3. Both RDB and AOF (hybrid approach)


## How RDB Works
The RDB method works by taking snapshots of the dataset at specified intervals. These snapshots are saved to disk in a binary format. When the server restarts, it can load the most recent snapshot to restore the dataset to the state it was in at the time of the snapshot. However, any changes made after the last snapshot will be lost.

## How AOF Works
The AOF method works by logging every write operation received by the server. This log is then written to a file in an append-only manner. When the server restarts, it can replay the log to reconstruct the dataset, essentially restoring the state of the key-value store to the point of the last write operation.

## Next Steps
I plan to implement the RDB method too in the future, as I already made the AOF method.