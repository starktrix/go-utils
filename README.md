# config-go


# io.Copy vs io.CopyBuffer


In Go, both `io.Copy` and `io.CopyBuffer` are used to copy data from an `io.Reader` to an `io.Writer`. However, there are important differences in terms of **control over buffer allocation** and **performance optimization**, which can be critical in high-performance or memory-sensitive applications.

---

### 🔧 `io.Copy(dst io.Writer, src io.Reader) (written int64, err error)`

This function is a **convenience wrapper** that internally allocates its **own buffer** (currently 32KB as of Go 1.22) to perform the copy.

#### ✅ Pros:
- Simple and easy to use.
- Ideal for general-purpose use where buffer reuse isn't a concern.
- Less code to manage.

#### ❌ Cons:
- **Allocates a new buffer every time** it is called — not ideal in high-frequency or high-throughput code paths.
- You **cannot control the buffer size**, which might be suboptimal depending on the I/O characteristics (e.g., network latency, disk block size).

---

### 🧠 `io.CopyBuffer(dst io.Writer, src io.Reader, buf []byte) (written int64, err error)`

This version lets you **pass in your own buffer**, giving you **control over memory allocation** and **buffer reuse**.

#### ✅ Pros:
- **Avoids repeated allocations**, which can reduce GC pressure and improve performance.
- Allows tuning the buffer size for **specific workloads** (e.g., aligning with network MTU, disk block size).
- Can be used in pooled systems where buffers are reused from a sync.Pool or similar.

#### ❌ Cons:
- Slightly more complex to use — you have to manage the buffer lifecycle.
- Incorrect buffer sizing can reduce performance if not tuned properly.

---

### 💡 Performance Example

In high-performance servers (e.g., reverse proxies, file servers), using `io.CopyBuffer` with a shared buffer (e.g., from a `sync.Pool`) can significantly reduce **memory allocations** and **GC pauses**, especially under heavy load.

```go
var bufPool = sync.Pool{
    New: func() interface{} {
        return make([]byte, 64*1024) // e.g., 64KB
    },
}

func copyHandler(dst io.Writer, src io.Reader) (int64, error) {
    buf := bufPool.Get().([]byte)
    defer bufPool.Put(buf)
    return io.CopyBuffer(dst, src, buf)
}
```

This reuse pattern avoids unnecessary allocations while still giving you the full power of buffered I/O.

---

### ✅ When to Use What?

| Use Case | Use `io.Copy` | Use `io.CopyBuffer` |
|----------|----------------|---------------------|
| Quick scripts, one-off tools | ✅ | ❌ |
| Performance-sensitive services | ❌ | ✅ |
| Custom buffer size needed | ❌ | ✅ |
| Memory reuse/zero allocation paths | ❌ | ✅ |
| Simpler code and you don’t care about performance | ✅ | ❌ |

---

### Final Thoughts

As a rule of thumb:
- Use `io.Copy` when **simplicity** is more important than **performance**.
- Use `io.CopyBuffer` when you care about **performance, memory management, or need a specific buffer size**.

If you're writing code where **scalability** and **predictable memory use** matter (e.g., concurrent file transfers, HTTP proxying, etc.), go with `io.CopyBuffer` and manage the buffer properly (often with a `sync.Pool`).

Would you like a benchmark comparison or real-world performance data between the two?


# io.Reader: consuming data

[how to compose interface in golang](https://www.youtube.com/watch?v=-gW7oSFxT2I&t=4s)

Yes, reading from an io.Reader in Go consumes the data, meaning that once you read from it, those bytes are no longer available for subsequent reads—unless the underlying data source supports re-reading (like bytes.Buffer or strings.Reader), or you explicitly store the read content.

🔍 Understanding io.Reader Behavior
An io.Reader is an interface:

go
Copy
Edit
type Reader interface {
    Read(p []byte) (n int, err error)
}
Each call to Read() pulls bytes from the data stream and advances the internal read cursor. Once the bytes are read, they’re not available again—unless:

✅ The Reader is:
A buffer-based reader like bytes.Buffer or strings.Reader, where the underlying data remains intact and can be re-read (or Seek can reposition).

A custom reader that internally retains or resets its data.

❌ The Reader is:
A streaming reader, like a file or a network socket, where data is processed in chunks and not retained unless you store it yourself.

An io.PipeReader, or the output of os.Stdin, or even a compressed stream (gzip.Reader)—these do not allow re-reading without buffering.

🧪 Example: bytes.Buffer (Reusable Reader)

```go
buf := bytes.NewBufferString("hello")
b := make([]byte, 5)

n, _ := buf.Read(b)
fmt.Println(string(b[:n]))  // Output: hello

n, _ = buf.Read(b)
fmt.Println(n)              // Output: 0 (buffer is now empty)
```
🛠️ If You Need to Read Multiple Times
If you want to read from an io.Reader more than once, you need to either:

Wrap it in a buffer and retain the content, like this:

```go
var buf bytes.Buffer
tee := io.TeeReader(originalReader, &buf)
// read from `tee`, and reuse `buf.Bytes()` later
```
Read all at once, store the data, and create new readers:

```go
data, _ := io.ReadAll(originalReader)
r1 := bytes.NewReader(data)
r2 := bytes.NewReader(data)
```
Summary
✅ Reading from io.Reader consumes the stream.

❌ You can’t "re-read" unless you explicitly store the data.

✅ Use TeeReader, ReadAll, or buffering strategies to retain data.

Let me know if you want a working code demo for any of these patterns.


# Golang cnditional build

<!-- foo_dev.go -->
```go
//go:build dev
// +build dev

package main

...
```

This hints the compiler that files ending in _dev should be compiled.


# Tunnelling SSH to HTTP

[Tunnel SSH Directly to HTTP with Golang](https://www.youtube.com/watch?v=RK5xh7xH1Jw)

Instead of using locks in golang, channels can be uitilized.

[runtime.GoSched](https://www.youtube.com/watch?v=INcV8kNlDPM&t=588s)

this is used to yield the control and execute another go rountine. Execution returns automatically.