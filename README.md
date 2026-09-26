
# SorobanAnchor Gate

> Programmable Soroban-to-SEP Gateway & Automated Compliance Escrow

[![License](https://img.shields.io/badge/License-Apache_2.0-blue.svg)](LICENSE)
[![Soroban](https://img.shields.io/badge/Soroban-v22.0.0-purple.svg)](https://soroban.stellar.org/)
[![Go Version](https://img.shields.io/badge/Go-1.22+-00ADD8.svg)](https://go.dev/)
[![Next.js](https://img.shields.io/badge/Next.js-14+-black.svg)](https://nextjs.org/)

SorobanAnchor Gate is an open-source middleware protocol bridging smart contracts on Stellar (**Soroban**) with Stellar's regulated banking off-ramps (**Stellar Ecosystem Proposals / SEPs**). It enables decentralized protocols, DAOs, and escrow applications to disburse funds directly into real-world bank accounts and mobile money wallets without manual intervention or centralized custodial custody.

---

## Table of Contents

- [Architecture Overview](#architecture-overview)
- [Core Components](#core-components)
- [Repository Structure](#repository-structure)
- [Supported Stellar Standards](#supported-stellar-standards)
- [Prerequisites](#prerequisites)
- [Quickstart Guide](#quickstart-guide)
  - [1. Clone the Repository](#1-clone-the-repository)
  - [2. Smart Contract Build and Test](#2-smart-contract-build-and-test)
  - [3. Run the Go Relayer Engine](#3-run-the-go-relayer-engine)
  - [4. Run the Frontend Dashboard](#4-run-the-frontend-dashboard)
- [Contributing](#contributing)
- [License](#license)

---

## Architecture Overview

```text
               +-------------------------------------------+
               |         Next.js / React Frontend          |
               | (Freighter Kit, KYC Upload, Live Tracker) |
               +---------------------+---------------------+
                                     |
              +----------------------+----------------------+
              |                                             |
              v                                             v
+-------------------------------+             +-------------------------------+
|     Soroban Smart Contract    |             |      Go Relayer & Gateway     |
|   (ConditionalEscrow.wasm)    |             |       (anchor-gate-relay)     |
+---------------+---------------+             +---------------+---------------+
                |                                             |
                | Emits Event:                                | Pulls Event via RPC,
                | DisbursementAuthorized                      | Signs SEP-10 Auth Challenge
                v                                             v
+-----------------------------------------------------------------------------+
|                               Stellar Network                               |
|                  (Horizon, Soroban RPC, SAC USDC/EURC)                      |
+-----------------------------------------------------------------------------+
                                                              |
                                                              | Executes SEP-12 / SEP-31
                                                              v
                                              +-------------------------------+
                                              |        Stellar Anchors        |
                                              |   (Bank Transfer, Mobile Money)|
                                              +-------------------------------+
```

The system operates across three core decoupled layers:

### Smart Contracts (`/contracts`)

A Soroban smart contract suite written in Rust managing milestone-based escrow locks for Stellar Asset Contract (SAC) tokens.

### Backend Relayer (`/backend`)

A Go service that monitors Soroban RPC event streams, signs SEP-10 cryptographic challenges, submits SEP-12 customer profiles, and executes automated SEP-31/SEP-6 off-ramp disbursements.

### Frontend Application (`/frontend`)

A Next.js dashboard supporting wallet connections (Freighter/Stellar Wallet Kit), anchor discovery via `stellar.toml`, and real-time escrow tracking.

---

## Repository Structure

```text
soroban-anchor-gate/
├── contracts/                  # Soroban smart contracts (Rust)
│   ├── Cargo.toml
│   └── src/
│       ├── lib.rs              # Contract entrypoint & interface
│       ├── storage.rs          # Ledger state definitions
│       └── test.rs             # Unit & integration tests
├── backend/                    # Anchor relay daemon & API (Go)
│   ├── go.mod
│   ├── go.sum
│   ├── cmd/
│   │   └── relay/
│   │       └── main.go         # Worker entrypoint
│   └── internal/
│       ├── listener/           # Soroban RPC event listener
│       ├── sephandler/         # SEP-10, SEP-12, SEP-31 clients
│       └── store/              # Event idempotency store
├── frontend/                   # Next.js web application
│   ├── package.json
│   ├── src/
│   │   ├── app/                # App router pages
│   │   ├── components/         # UI components & Wallet modal
│   │   └── lib/                # Stellar SDK & Anchor utilities
│   └── tailwind.config.ts
├── .github/                    # CI/CD and Issue workflows
│   └── workflows/
│       └── ci.yml
├── CONTRIBUTING.md             # Contribution guidelines
├── LICENSE                     # Apache 2.0 License
└── README.md
```

---

## Supported Stellar Standards

| Standard | Description |
|---|---|
| SEP-1 | Stellar Info File (`stellar.toml` discovery) |
| SEP-10 | Stellar Web Authentication |
| SEP-12 | KYC API Integration |
| SEP-31 | Cross-Border Payout API |
| SEP-38 | Anchor RFQ / Quotes API |
| CAP-46-06 / Soroban | Smart Contract Execution Environment |

---

## Prerequisites

Ensure you have the following installed locally.

### Rust

Rust v1.74+ with the `wasm32-unknown-unknown` target:

```bash
rustup target add wasm32-unknown-unknown
```

### Stellar CLI

```bash
cargo install --locked stellar-cli --features opt
```

### Go

Go v1.22 or later.

### Node.js

Node.js v18 or later with npm or pnpm.

### Docker

Optional, for a local Soroban/PostgreSQL sandbox.

---

## Quickstart Guide

### 1. Clone the Repository

Replace `<your-org-or-username>` with the actual GitHub organization or username.

```bash
git clone https://github.com/<your-org-or-username>/soroban-anchor-gate.git
cd soroban-anchor-gate
```

### 2. Smart Contract Build and Test

Navigate to the contracts directory:

```bash
cd contracts
```

Run the contract tests:

```bash
cargo test
```

Build the contract for WebAssembly:

```bash
cargo build --target wasm32-unknown-unknown --release
```

### 3. Run the Go Relayer Engine

Navigate to the backend directory:

```bash
cd ../backend
```

Configure environment variables:

```bash
cp .env.example .env
```

Populate `.env` with your Stellar Testnet configuration:

```env
STELLAR_NETWORK_PASSPHRASE="Test SDF Network ; September 2015"
SOROBAN_RPC_URL="https://soroban-testnet.stellar.org"
RELAY_SECRET_KEY="S..."
TARGET_ANCHOR_DOMAIN="testanchor.stellar.org"
DATABASE_URL="postgres://user:pass@localhost:5432/anchorgate?sslmode=disable"
```

**Security Notice:** Never commit your `.env` file, private keys, or other secrets to version control. Use a dedicated test account for development.

Start the relayer service:

```bash
go run cmd/relay/main.go
```

### 4. Run the Frontend Dashboard

Open a new terminal and navigate to the frontend directory:

```bash
cd frontend
```

Install dependencies:

```bash
npm install
```

Start the development server:

```bash
npm run dev
```

Open the local development URL displayed by Next.js in your terminal.

---

## Contributing

We welcome community contributions.

Please review [CONTRIBUTING.md](CONTRIBUTING.md) for details on our coding standards, branch conventions, and testing requirements before opening a Pull Request.

---

## License

This project is licensed under the Apache License 2.0. See the [LICENSE](LICENSE) file for details.
