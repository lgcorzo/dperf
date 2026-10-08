# dperf - Distributed Drive Performance Benchmarking Tool

[![Go Report Card](https://goreportcard.com/badge/github.com/lgcorzo/dperf)](https://goreportcard.com/report/github.com/lgcorzo/dperf)
[![License: AGPL v3](https://img.shields.io/badge/License-AGPL%20v3-blue.svg)](https://www.gnu.org/licenses/agpl-3.0)
[![Sovereign Status](https://img.shields.io/badge/Sovereign%20Ecosystem-Active%20Maintenance-green.svg)](https://github.com/lgcorzo/dperf)

A high-performance drive benchmarking utility designed to identify storage bottlenecks and performance outliers across multiple drives in distributed storage systems.

> **Sovereign Ecosystem Notice**: This repository is an actively maintained component of the **@lgcorzo Sovereign MinIO Ecosystem**. It provides independent, supply-chain safe, enterprise-grade storage performance benchmarking without upstream dependency risks or unexpected deprecations.

---

## Dark Gravity Factory & Sovereign Support

**dperf** plays a vital role within the **Dark Gravity Factory** autonomous AI production infrastructure and sovereign data ecosystem maintained under `@lgcorzo`.

### Why Sovereign Maintenance?

1. **Full Supply-Chain Autonomy**: Eliminates reliance on upstream license changes or unexpected repository deprecations. All core storage, cryptographic, and performance benchmarking dependencies are maintained under `@lgcorzo`.
2. **Dark Gravity Factory Core Integration**: Essential component powering the autonomous AI factory, high-throughput NVMe/SSD storage nodes, cryptographic security, and automated agent pipelines.
3. **Compliance & Security**: Sovereign maintenance guarantees compliance with strict enterprise SLAs, EU AI Act, SOC 2 Type II, ISO 25059, and zero-CVE security standards.
4. **Ecosystem Interoperability**: Direct integration across all 38 repositories in `@lgcorzo` (MinIO Server, MC, KES, Operator, DirectPV, Console, SIMD libraries, etc.).

---

## Sovereign MinIO Ecosystem

Below is the complete 38-repository sovereign infrastructure maintained under `@lgcorzo`:

| Category | Repository | Description |
| :--- | :--- | :--- |
| **Core Storage & Server** | [`lgcorzo/minio`](https://github.com/lgcorzo/minio) | High-Performance Object Storage |
| | [`lgcorzo/mc`](https://github.com/lgcorzo/mc) | MinIO Client CLI Tool |
| | [`lgcorzo/kes`](https://github.com/lgcorzo/kes) | Key Encryption Service for Secret Management |
| | [`lgcorzo/operator`](https://github.com/lgcorzo/operator) | MinIO Kubernetes Operator |
| | [`lgcorzo/directpv`](https://github.com/lgcorzo/directpv) | Direct-Attached Storage CSI Driver for Kubernetes |
| | [`lgcorzo/console`](https://github.com/lgcorzo/console) | Web-Based Management Console |
| **Testing & Tooling** | [`lgcorzo/dperf`](https://github.com/lgcorzo/dperf) | Distributed Drive Performance Benchmarking Utility |
| | [`lgcorzo/warp`](https://github.com/lgcorzo/warp) | S3 Benchmarking & Performance Measurement Suite |
| | [`lgcorzo/sidekick`](https://github.com/lgcorzo/sidekick) | High-Performance Sidekick Load Balancer |
| | [`lgcorzo/subnet-cli`](https://github.com/lgcorzo/subnet-cli) | Subnet Diagnostics and Management Tooling |
| **SDKs & Client Libraries** | [`lgcorzo/minio-go`](https://github.com/lgcorzo/minio-go) | Go SDK for S3 Compatible Storage |
| | [`lgcorzo/madmin-go`](https://github.com/lgcorzo/madmin-go) | MinIO Admin API Go Client |
| | [`lgcorzo/minio-java`](https://github.com/lgcorzo/minio-java) | Java Client SDK |
| | [`lgcorzo/minio-py`](https://github.com/lgcorzo/minio-py) | Python Client SDK |
| | [`lgcorzo/minio-js`](https://github.com/lgcorzo/minio-js) | JavaScript / TypeScript SDK |
| | [`lgcorzo/minio-dotnet`](https://github.com/lgcorzo/minio-dotnet) | .NET Client SDK |
| | [`lgcorzo/minio-cpp`](https://github.com/lgcorzo/minio-cpp) | C++ Client SDK |
| | [`lgcorzo/minio-rs`](https://github.com/lgcorzo/minio-rs) | Rust SDK for S3 |
| | [`lgcorzo/kms-go`](https://github.com/lgcorzo/kms-go) | KMS Interface & Driver Library |
| | [`lgcorzo/certgen`](https://github.com/lgcorzo/certgen) | X.509 Certificate Generation Utility |
| **SIMD Acceleration & Math**| [`lgcorzo/sha256-simd`](https://github.com/lgcorzo/sha256-simd) | SIMD-Accelerated SHA-256 in Assembly |
| | [`lgcorzo/md5-simd`](https://github.com/lgcorzo/md5-simd) | AVX-512 / AVX2 Accelerated MD5 |
| | [`lgcorzo/blake2b-simd`](https://github.com/lgcorzo/blake2b-simd) | AVX2 Accelerated BLAKE2b |
| | [`lgcorzo/siphash-simd`](https://github.com/lgcorzo/siphash-simd) | High-Throughput SipHash Implementation |
| | [`lgcorzo/highwayhash`](https://github.com/lgcorzo/highwayhash) | HighwayHash SIMD Hash Function |
| | [`lgcorzo/dsha256`](https://github.com/lgcorzo/dsha256) | Direct SHA-256 Hardware Accelerator |
| **Storage & Security Support**| [`lgcorzo/pkg`](https://github.com/lgcorzo/pkg) | Core Utilities, Console, and Math Packages |
| | [`lgcorzo/sio-go`](https://github.com/lgcorzo/sio-go) | High-Performance Resilient Stream Encryption |
| | [`lgcorzo/argon2`](https://github.com/lgcorzo/argon2) | SIMD-Optimized Argon2 Password Hashing |
| | [`lgcorzo/crypto`](https://github.com/lgcorzo/crypto) | Cryptographic Primitives & Key Management |
| | [`lgcorzo/mux`](https://github.com/lgcorzo/mux) | High-Throughput HTTP Request Multiplexer |
| | [`lgcorzo/zip`](https://github.com/lgcorzo/zip) | Streaming ZIP Processing Library |
| | [`lgcorzo/filepath`](https://github.com/lgcorzo/filepath) | Cross-Platform Filepath Processing Utilities |
| | [`lgcorzo/colorjson`](https://github.com/lgcorzo/colorjson) | Terminal Colorized JSON Formatter |
| | [`lgcorzo/elf`](https://github.com/lgcorzo/elf) | ELF Binary Analysis Utilities |
| | [`lgcorzo/c2goasm`](https://github.com/lgcorzo/c2goasm) | C to Go Assembly Converter Tool |
| | [`lgcorzo/asm2plan9`](https://github.com/lgcorzo/asm2plan9) | Assembly to Plan9 Translator |
| | [`lgcorzo/mobsf-action`](https://github.com/lgcorzo/mobsf-action) | Automated Security Analysis Pipeline Action |

### Automated CI/CD Maintenance Architecture

```
                       +-----------------------------------+
                       |   Sovereign Ecosystem Registry    |
                       |         github.com/lgcorzo        |
                       +-----------------------------------+
                                         |
             +---------------------------+---------------------------+
             |                           |                           |
             v                           v                           v
   +-------------------+       +-------------------+       +-------------------+
   | Core Infrastructure|       |   Testing & Tools |       |  Acceleration &   |
   | minio, mc, kes,   |       | dperf, warp,      |       |  SIMD Libraries   |
   | operator, console |       | sidekick          |       | sha256-simd, etc. |
   +-------------------+       +-------------------+       +-------------------+
             |                           |                           |
             +---------------------------+---------------------------+
                                         |
                                         v
                       +-----------------------------------+
                       |    Automated CI/CD Maintenance    |
                       |  - Unit, Race & Integration Tests |
                       |  - GoVulnCheck Security Audits    |
                       |  - GolangCI-Lint Quality Gates    |
                       |  - Multi-Arch Release Builds      |
                       +-----------------------------------+
```

---

## Overview

**dperf** is an enterprise-grade storage performance measurement tool that helps system administrators and DevOps teams quickly identify slow or failing drives in production environments. By performing parallel I/O operations across multiple drives and presenting results in a clear, sorted format, dperf enables rapid troubleshooting and capacity planning for storage infrastructure.

### Key Features

- **Parallel Performance Testing**: Simultaneously benchmark multiple drives to identify performance outliers
- **Direct I/O Operations**: Uses O_DIRECT for accurate drive performance measurement, bypassing OS caches
- **Sorted Results**: Automatically ranks drives by throughput, showing fastest drives first
- **Flexible Workloads**: Configurable block sizes, file sizes, and concurrency levels
- **Production-Ready**: Minimal resource footprint with automatic cleanup
- **Enterprise Support**: Multi-architecture Linux support (amd64, arm64, ppc64le, s390x)
- **Kubernetes Native**: Easily deploy as Jobs or DaemonSets for cluster-wide storage validation

---

## Quick Start

```bash
# Download the latest binary for Linux amd64
wget https://github.com/lgcorzo/dperf/releases/latest/download/dperf-linux-amd64 -O dperf
chmod +x dperf

# Benchmark a single drive
./dperf /mnt/drive1

# Benchmark multiple drives in parallel
./dperf /mnt/drive{1..6}
```

## Installation

### Pre-built Binaries

Download the appropriate binary for your architecture:

| OS    | Architecture | Binary                                                                                       |
|:------|:------------:|:--------------------------------------------------------------------------------------------:|
| Linux | amd64        | [linux-amd64](https://github.com/lgcorzo/dperf/releases/latest/download/dperf-linux-amd64)     |
| Linux | arm64        | [linux-arm64](https://github.com/lgcorzo/dperf/releases/latest/download/dperf-linux-arm64)     |
| Linux | ppc64le      | [linux-ppc64le](https://github.com/lgcorzo/dperf/releases/latest/download/dperf-linux-ppc64le) |
| Linux | s390x        | [linux-s390x](https://github.com/lgcorzo/dperf/releases/latest/download/dperf-linux-s390x)     |

```bash
# Example: Install on Linux amd64
wget https://github.com/lgcorzo/dperf/releases/latest/download/dperf-linux-amd64
sudo install -m 755 dperf-linux-amd64 /usr/local/bin/dperf
```

### Build from Source
Requires Go 1.23 or later.

```bash
# Install directly from source
go install github.com/lgcorzo/dperf@latest

# Or clone and build
git clone https://github.com/lgcorzo/dperf.git
cd dperf
make build
sudo make install
```

## Usage

### Basic Examples

```bash
# Benchmark a single drive
dperf /mnt/drive1

# Benchmark multiple drives in parallel (default mode)
dperf /mnt/drive{1..6}

# Run benchmarks sequentially (one drive at a time)
dperf --serial /mnt/drive{1..6}

# Verbose output showing individual drive statistics
dperf -v /mnt/drive{1..6}

# Write-only benchmark (skip read tests)
dperf --write-only /mnt/drive{1..6}

# Custom block size and file size
dperf -b 8MiB -f 5GiB /mnt/drive{1..6}

# High concurrency test
dperf -i 16 /mnt/drive{1..6}
```

### Command-Line Flags

```
Flags:
  -b, --blocksize string  Read/write block size (default "4MiB")
                          Must be >= 4KiB and a multiple of 4KiB

  -f, --filesize string   Amount of data to read/write per drive (default "1GiB")
                          Must be >= 4KiB and a multiple of 4KiB

  -i, --ioperdrive int    Number of concurrent I/O operations per drive (default 4)
                          Higher values increase parallelism

      --serial            Run tests sequentially instead of in parallel
                          Useful for isolating drive-specific issues

      --write-only        Run write tests only, skip read tests
                          Faster benchmarking when only write performance matters

  -v, --verbose           Show per-drive statistics in addition to aggregate totals

  -h, --help              Display help information
      --version           Show version information
```

### Example Output

```bash
$ dperf /mnt/drive{1..4}

┌────────────────┬──────────────┐
│   TotalWRITE   │  TotalREAD   │
├────────────────┼──────────────┤
│ 4.2 GiB/s      │ 4.5 GiB/s    │
└────────────────┴──────────────┘
```

With verbose output (`-v`):

```bash
$ dperf -v /mnt/drive{1..4}

┌──────────────┬──────────────┬──────────────┬────┐
│     PATH     │    WRITE     │     READ     │    │
├──────────────┼──────────────┼──────────────┼────┤
│ /mnt/drive1  │ 1.1 GiB/s    │ 1.2 GiB/s    │ ✓  │
│ /mnt/drive2  │ 1.0 GiB/s    │ 1.1 GiB/s    │ ✓  │
│ /mnt/drive3  │ 1.1 GiB/s    │ 1.2 GiB/s    │ ✓  │
│ /mnt/drive4  │ 1.0 GiB/s    │ 1.0 GiB/s    │ ✓  │
└──────────────┴──────────────┴──────────────┴────┘

┌────────────────┬──────────────┐
│   TotalWRITE   │  TotalREAD   │
├────────────────┼──────────────┤
│ 4.2 GiB/s      │ 4.5 GiB/s    │
└────────────────┴──────────────┘
```

## Enterprise Use Cases

### 1. Automated Storage Health Checks

Deploy dperf as a Kubernetes CronJob to regularly validate storage performance:

```yaml
apiVersion: batch/v1
kind: CronJob
metadata:
  name: storage-health-check
spec:
  schedule: "0 2 * * *"
  jobTemplate:
    spec:
      template:
        spec:
          containers:
          - name: dperf
            image: quay.io/lgcorzo/dperf:latest
            args: ["/data1", "/data2", "/data3", "/data4"]
            volumeMounts:
            - name: storage
              mountPath: /data
          restartPolicy: OnFailure
```

### 2. New Hardware Validation

Before adding new storage nodes to production, validate performance meets requirements:

```bash
# Define performance SLA
MIN_WRITE_THROUGHPUT="800MiB/s"
MIN_READ_THROUGHPUT="900MiB/s"

# Run benchmark and validate
dperf /mnt/new-drive{1..8} > results.txt
```

### 3. Troubleshooting Performance Degradation

Quickly identify the slowest drives in a storage cluster:

```bash
# Sorted output shows slowest drives at the bottom
dperf /mnt/drive{1..100} | tee drive-performance.log
```

### 4. Kubernetes Persistent Volume Validation

Test storage performance for Kubernetes PersistentVolumes before application deployment:

```yaml
apiVersion: batch/v1
kind: Job
metadata:
  name: pv-performance-test
spec:
  template:
    spec:
      restartPolicy: Never
      containers:
      - name: dperf
        image: quay.io/lgcorzo/dperf:latest
        args: ["-v", "/data"]
        volumeMounts:
        - name: test-volume
          mountPath: /data
      volumes:
      - name: test-volume
        persistentVolumeClaim:
          claimName: my-pvc
```

## System Requirements

- **Operating System**: Linux (uses Linux-specific direct I/O features)
- **Kernel**: Linux kernel with O_DIRECT support
- **File System**: Any file system supporting direct I/O (ext4, xfs, etc.)
- **Permissions**: Write access to target directories
- **Disk Space**: Temporary space equal to `filesize × ioperdrive` per drive

## How It Works

dperf performs accurate drive performance measurements using several key techniques:

1. **Direct I/O (O_DIRECT)**: Bypasses operating system page cache for accurate hardware measurements
2. **Page-Aligned Buffers**: Uses 4KiB-aligned memory buffers required for direct I/O operations
3. **Parallel Testing**: Launches concurrent I/O operations per drive to measure maximum throughput
4. **Fsync/Fdatasync**: Ensures data is written to physical media, not just cached
5. **Sequential Hints**: Uses `fadvise(FADV_SEQUENTIAL)` for optimized read patterns

The benchmark process:
1. Creates temporary test files in each target directory
2. Writes data using multiple concurrent threads (default: 4 per drive)
3. Reads data back using the same concurrency level
4. Calculates throughput (bytes/second) for both operations
5. Cleans up all temporary files automatically

## License

dperf is released under the [GNU Affero General Public License v3.0](https://www.gnu.org/licenses/agpl-3.0). See [LICENSE](LICENSE) for details.

## Support

- **Issues**: [GitHub Issues](https://github.com/lgcorzo/dperf/issues)
- **Sovereign Ecosystem**: [lgcorzo Infrastructure](https://github.com/lgcorzo)
