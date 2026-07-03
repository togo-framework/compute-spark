<!-- togo-header -->
<div align="center">
  <img src=".github/assets/togo-mark.svg" alt="togo" height="64" />
  <h1>togo-framework/compute-spark</h1>
  <p>
    <a href="https://to-go.dev/marketplace"><img src="https://img.shields.io/badge/marketplace-to--go.dev-1FC7DC" alt="marketplace" /></a>
    <a href="https://pkg.go.dev/github.com/togo-framework/compute-spark"><img src="https://pkg.go.dev/badge/github.com/togo-framework/compute-spark.svg" alt="pkg.go.dev" /></a>
    <img src="https://img.shields.io/badge/license-MIT-blue" alt="MIT" />
  </p>
  <p><strong>Part of the <a href="https://to-go.dev">togo</a> framework.</strong></p>
</div>

## Install

```bash
togo install togo-framework/compute-spark
```

<!-- /togo-header -->

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
<!-- togo-sponsors -->
---

<div align="center">
  <h3>💎 Premium sponsors</h3>
  <p>
    <a href="https://id8media.com"><img src=".github/assets/id8media.svg" height="44" alt="ID8 Media" /></a>
    &nbsp;&nbsp;&nbsp;&nbsp;
    <a href="https://one-studio.co"><img src=".github/assets/one-studio.jpeg" height="44" alt="One Studio" /></a>
  </p>
  <p><sub>Support togo — <a href="https://github.com/sponsors/fadymondy">become a sponsor</a>.</sub></p>
</div>
<!-- /togo-sponsors -->
