
# Contributing to SorobanAnchor Gate

Thank you for your interest in contributing to **SorobanAnchor Gate**! We are building an open-source, programmable bridge between Soroban smart contracts and Stellar's off-ramp anchors (SEPs).

Whether you are here through the **Drips Wave** program, fixing a bug, improving documentation, or adding new features, this guide outlines the workflow and standards expected across our monorepo.

---

## Table of Contents

- [Code of Conduct](#code-of-conduct)
- [Drips Wave Contributors](#drips-wave-contributors)
- [Monorepo Architecture](#monorepo-architecture)
- [Development Workflow](#development-workflow)
  - [1. Fork and Clone](#1-fork-and-clone)
  - [2. Branching Strategy](#2-branching-strategy)
  - [3. Local Development and Testing](#3-local-development-and-testing)
- [Commit Guidelines](#commit-guidelines)
- [Submitting a Pull Request](#submitting-a-pull-request-pr)
- [Getting Help](#getting-help)

---

## Code of Conduct

We are committed to providing a welcoming, inclusive, and harassment-free environment for all contributors.

Please:

- Treat everyone with respect.
- Give constructive feedback.
- Prioritize collaborative problem-solving.
- Respect different perspectives and experiences.
- Keep discussions focused on improving the project.

---

## Drips Wave Contributors

If you are contributing through the **Drips Wave** program, follow the contribution requirements associated with the relevant issue and program.

### 1. Find an Issue

Browse the repository's [Issues](https://github.com/<your-org-or-username>/soroban-anchor-gate/issues) tab and look for issues labeled:

- `wave:trivial` — 100 points
- `wave:medium` — 150 points
- `wave:high` — 200 points

Confirm that the issue is available and that you understand its requirements before beginning work.

### 2. Claiming an Issue

1. Leave a comment on the issue explaining your proposed approach.
2. Express your intention to work on the issue.
3. Wait for a maintainer to assign the issue to you.
4. Begin implementation after assignment.

Avoid working on an issue that has already been assigned to another contributor unless a maintainer gives permission.

### 3. Review Protocol

After your pull request is approved and merged, follow the applicable Drips Wave completion process.

If the program requires a two-way collaborator review, complete it within the specified period, including the 14-day period stated in the applicable program instructions.

Check the current Drips Wave rules for the exact requirements, deadlines, and point allocation.

---

## Monorepo Architecture

SorobanAnchor Gate is organized into three primary components:

| Directory | Technology | Purpose |
|---|---|---|
| `/contracts` | Rust / Soroban | Smart contracts and on-chain escrow logic |
| `/backend` | Go | Event listening, relayer services, and SEP integration |
| `/frontend` | Next.js / TypeScript / Tailwind CSS | User interface and wallet integration |

Contributors should keep changes focused on the component relevant to their assigned issue.

Changes affecting multiple components should explain how those components interact and include appropriate tests.

---

## Development Workflow

### 1. Fork and Clone

Fork the repository to your own GitHub account.

Clone your fork:

```bash
git clone https://github.com/<your-username>/soroban-anchor-gate.git
cd soroban-anchor-gate
```

Add the original repository as the upstream remote:

```bash
git remote add upstream https://github.com/<your-org-or-username>/soroban-anchor-gate.git
```

Verify your remotes:

```bash
git remote -v
```

Replace `<your-username>` and `<your-org-or-username>` with the appropriate GitHub usernames or organization names.

### 2. Branching Strategy

Always create a descriptive branch from the latest `main` branch.

Update your local repository:

```bash
git fetch upstream
git checkout main
git rebase upstream/main
```

Create a branch based on the type of change.

#### Features

```bash
git checkout -b feature/sep10-auth-token-caching
```

#### Bug Fixes

```bash
git checkout -b fix/event-listener-reconnect
```

#### Documentation

```bash
git checkout -b docs/update-rpc-setup
```

Use clear, descriptive branch names that communicate the purpose of your work.

### 3. Local Development and Testing

Before submitting changes, run the relevant tests and validation commands for the component you modified.

The commands below describe the expected development workflow. Some commands may require additional tooling or configuration depending on the project setup.

#### A. Smart Contracts (`/contracts`)

**Prerequisites:** Rust 1.74 or later, the WebAssembly target, and the Stellar CLI.

Navigate to the contracts directory:

```bash
cd contracts
```

Run unit tests:

```bash
cargo test
```

Check formatting:

```bash
cargo fmt --check
```

Run Clippy:

```bash
cargo clippy --all-targets -- -D warnings
```

Build the release WebAssembly binary:

```bash
cargo build --target wasm32-unknown-unknown --release
```

Ensure contract changes preserve authorization rules, state transitions, and expected behavior.

#### B. Backend Relayer (`/backend`)

**Prerequisites:** Go 1.22 or later.

Navigate to the backend directory:

```bash
cd backend
```

Download dependencies:

```bash
go mod download
```

Run unit and integration tests:

```bash
go test -v ./...
```

Run the race detector:

```bash
go test -race ./...
```

Verify code quality:

```bash
go vet ./...
```

If `golangci-lint` is installed and configured:

```bash
golangci-lint run
```

Backend changes should account for event processing, retry behavior, idempotency, error handling, and secure management of credentials.

#### C. Frontend (`/frontend`)

**Prerequisites:** Node.js 18 or later and npm or pnpm.

Navigate to the frontend directory:

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

Run the linter:

```bash
npm run lint
```

Build the application:

```bash
npm run build
```

Frontend changes should preserve accessibility, responsive behavior, wallet interaction safety, and clear transaction status feedback.

---

## Commit Guidelines

We use the **Conventional Commits** specification to maintain a clean and readable Git history.

### Commit Format

```text
<type>(<scope>): <short description>
```

### Allowed Types

| Type | Description |
|---|---|
| `feat` | A new feature or capability |
| `fix` | A bug fix |
| `docs` | Documentation updates |
| `test` | Adding or updating tests |
| `refactor` | Code changes that neither fix bugs nor add features |
| `chore` | Build scripts, CI pipelines, or dependency updates |

### Examples

```text
feat(backend): implement SEP-38 quote request client
```

```text
fix(contracts): handle escrow timeout timestamp arithmetic edge case
```

```text
docs(readme): add Docker Compose instructions for local Horizon
```

Keep commit messages concise and descriptive.

Avoid vague messages such as:

```text
update
```

```text
fix stuff
```

---

## Submitting a Pull Request (PR)

Before opening a pull request, ensure your changes are focused, tested, and ready for review.

### 1. Keep It Focused

Avoid mixing unrelated changes into a single pull request.

Each PR should address the specific issue assigned to you.

If your work requires changes across multiple components, explain why those changes are necessary.

### 2. Sync with Main

Fetch the latest upstream changes and rebase your branch:

```bash
git fetch upstream
git rebase upstream/main
```

Resolve any conflicts and rerun the relevant tests.

### 3. Push Your Branch

```bash
git push origin <your-branch-name>
```

If you have already pushed the branch and rebased it, you may need:

```bash
git push --force-with-lease origin <your-branch-name>
```

Use `--force-with-lease` carefully and only on your own feature branch.

### 4. Open a Pull Request

Open a pull request against the original repository's `main` branch.

Provide a clear description containing:

- **Problem:** What issue does this PR solve?
- **Summary:** What changes were made?
- **Implementation:** How were the changes implemented?
- **Testing:** What tests or validation commands were run?
- **Related issue:** Include the relevant issue number, for example, `Closes #12`.

### 5. CI Checks

Ensure all applicable automated checks pass.

Depending on the affected component, these may include:

- Rust contract tests
- Rust formatting and Clippy
- Go tests
- Go linting
- Frontend linting
- Frontend build
- Integration tests

Pull requests with failing checks may require changes before they can be merged.

### 6. Respond to Review Feedback

Review feedback is part of the contribution process.

Please:

- Respond respectfully.
- Ask questions when feedback is unclear.
- Make requested changes.
- Explain technical disagreements constructively.
- Keep the discussion focused on the code and project requirements.

---

## Security Contributions

Security-related changes require particular care because SorobanAnchor Gate interacts with smart contracts, wallets, and financial infrastructure.

When contributing security-sensitive code:

- Never commit private keys, seed phrases, passwords, or API secrets.
- Do not expose credentials in logs or error messages.
- Follow the project's authorization requirements.
- Validate untrusted inputs.
- Consider transaction replay and duplicate event processing.
- Avoid introducing unsafe assumptions about anchor responses.
- Add tests for security-sensitive behavior.

If you discover a vulnerability, do not disclose sensitive exploit details in a public issue.

Use the project's designated private security reporting channel when one is available.

---

## Documentation Contributions

Documentation improvements are welcome.

Examples include:

- README updates
- Installation instructions
- API documentation
- Smart contract explanations
- Architecture diagrams
- Troubleshooting guides
- Developer tutorials

Ensure documentation matches the actual implementation.

Do not document planned functionality as if it has already been implemented.

---

## Getting Help

If you run into issues, have architecture questions, or need clarification about Stellar or Soroban specifications:

1. Open a discussion in the repository's Discussions tab, if enabled.
2. Comment directly on the GitHub issue you are assigned to.
3. Review the project's existing documentation.
4. Consult the official Stellar developer documentation.

When asking for help, include:

- The component you are working on.
- The issue or feature involved.
- The command you ran.
- The error message or unexpected behavior.
- The steps you have already tried.

Please avoid sharing private keys, seed phrases, or sensitive credentials when requesting support.

---

## Thank You

Thank you for helping build **SorobanAnchor Gate**.

Every contribution—whether code, testing, documentation, security review, or constructive feedback—helps improve the project and strengthen the open-source Stellar ecosystem.
