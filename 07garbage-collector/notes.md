# Understanding GOGC in Go

`GOGC` is an environment variable in Go that controls **Garbage Collection (GC)** behavior. It determines **how often the garbage collector runs** by setting the **growth target of the heap**.

---

## 🔍 What is `GOGC`?

`GOGC` stands for **Go Garbage Collection percentage**.

It defines the **percentage by which the heap must grow** before the GC runs again.

---

## 🧠 How it Works

The logic behind `GOGC`:

> **GOGC = (HeapGrowth / LiveHeap) * 100**

- **LiveHeap**: The size of live (in-use) objects after the last GC.
- **HeapGrowth**: How much additional memory the program can allocate before the next GC.

So:

> **HeapGrowth = (GOGC / 100) * LiveHeap**


---

## 📈 Example

Assume:
- After GC, the heap has 100 MB of live data.
- `GOGC=100` (default value).

Then:
- GC will run again when heap size reaches **100 MB + 100 MB = 200 MB**.

---

## 🔧 Tuning `GOGC`

| GOGC Value      | Effect                                                               |
|-----------------|----------------------------------------------------------------------|
| Higher (e.g., 200) | GC runs **less often** → Better throughput, more memory usage       |
| Lower (e.g., 50)   | GC runs **more often** → Less memory usage, more CPU overhead       |
| 0                 | Disables GC entirely (not recommended for production)              |
| < 0               | Invalid, Go falls back to the default (100)                        |

---

## ⚙️ How to Set `GOGC`

### In the terminal:

```bash
GOGC=150 ./your-app
