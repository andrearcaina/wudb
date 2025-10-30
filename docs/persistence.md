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

## Example Usage

To run the server with AOF persistence enabled, you can use the following command:

```bash
> ./bin/main -a -p test.aof
2025/10/30 18:58:52 AOF file is empty, starting with empty database
Starting server on :8080
2025/10/30 18:58:58 "POST http://localhost:8080/set HTTP/1.1" from [::1]:34906 - 201 30B in 3.709444ms
2025/10/30 18:59:01 "GET http://localhost:8080/get/name HTTP/1.1" from [::1]:33502 - 200 17B in 30.802µs
^C2025/10/30 18:59:03 Shutting down server...
```

You can see that the AOF file is created (because it was empty/nonexistent), and then the user sends a POST request to set a key-value pair, and then a GET request to retrieve the value.

The AOF file (`test.aof`) will contain the logged operations:

```
{"operation":"SET","key":"name","value":"John"}
```

When running the server again with the same AOF file, the data is loaded from the file, and the user can retrieve the previously set value, and even delete it.

```bash
> ./bin/main -a -p test.aof
2025/10/30 19:02:36 Loaded 1 operations from AOF file
Starting server on :8080
2025/10/30 19:02:38 "POST http://localhost:8080/set HTTP/1.1" from [::1]:49596 - 201 30B in 3.15874ms
2025/10/30 19:02:44 "DELETE http://localhost:8080/del/name HTTP/1.1" from [::1]:49598 - 204 0B in 3.002684ms
2025/10/30 19:02:47 "GET http://localhost:8080/get/name HTTP/1.1" from [::1]:41848 - 404 14B in 26.681µs
^C2025/10/30 19:02:51 Shutting down server...
```

Then, the AOF file (`test.aof`) will be updated to reflect the delete operation:

```
{"operation":"SET","key":"name","value":"John"}
{"operation":"DEL","key":"name"}
```