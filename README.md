<!-- togo-brand -->
<p align="center"><img src=".github/assets/togo-mark.svg" width="96" alt="togo" /></p>
<h1 align="center">compute-spark</h1>
<p align="center"><sub>part of the <a href="https://github.com/togo-framework">togo-framework</a></sub></p>

A togo **compute** backend that submits jobs to **Apache Spark** via `spark-submit`.
Registers into [compute](https://github.com/togo-framework/compute)'s slot.

```bash
togo install togo-framework/compute-spark
togo provider:use compute spark
togo config:set SPARK_MASTER yarn
```

| Key | Default | Purpose |
|---|---|---|
| `SPARK_SUBMIT_BIN` | `spark-submit` | submitter binary |
| `SPARK_MASTER` | — | `--master` value |
| `SPARK_OPTS` | — | extra opts |

MIT © fadymondy
