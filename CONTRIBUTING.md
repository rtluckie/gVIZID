# Contributing

Thanks for the interest in helping. This project is intentionally lightweight and the goal is to keep behavior consistent across ports.

## Principles

- Keep the interface consistent across ports (Go/Rust/Python).
- Prefer small, reviewable PRs over large refactors.
- Update docs whenever behavior, flags, or formats change.

## Suggested workflow

1. Open an issue or start a discussion for changes that affect ID format, encoding, or CLI behavior.
2. Fork and create a feature branch.
3. Add tests when possible, especially for codec and timestamp behavior.
4. Keep commit messages conventional and descriptive.
5. Submit a PR and note any cross-port implications.

## What to include in PRs

- A clear description of what changed and why
- Updated docs in `docs/` when relevant
- Any CLI changes documented in `docs/cli.md`
