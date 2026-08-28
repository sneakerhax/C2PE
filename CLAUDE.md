# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Purpose

C2PE is a collection of experimental Command & Control (C2) and post-exploitation code samples, used for red team tradecraft research and to teach defenders how these techniques work so they can build detections against them. Each subdirectory is a self-contained, independent example — there is no shared build system, package, or entrypoint tying them together.

## Repository structure

- `Command_and_Control/` — C2 server/client/channel examples (HTTP, TCP, Discord-as-C2-channel, ngrok tunneling, in-memory execution via memfd, container backdooring). Each subfolder has its own README describing the specific technique and how to run it.
- `Post_Exploitation/` — standalone post-exploitation modules, organized by category (`General/C++`, `General/Go`, `Kubernetes`), each a single small program (host info, list files/processes, print env, print file contents, grab Kubernetes service account tokens).

There is no top-level build, lint, or test tooling — each example is built and run independently per its own README.

## Working with individual examples

**Go examples**: each has its own `go.mod` (module names and Go versions are inconsistent across subfolders — do not assume a shared module or Go version). Build with `go build <file>.go` from within that subfolder. Some (`async_ex`, `async_ex_to_discord`) have a separate `cli/` Go module for a companion CLI tool (built with `urfave/cli` + `olekukonko/tablewriter`) that talks to the server's HTTP API.

**Python examples** (`async_ex/server.py`, `async_ex_to_discord/server.py`, `python_stream_c2/`): require Python 3.7+. The `async_ex*` servers run via Gunicorn and are commonly built/run in Docker using the provided `Dockerfile`.

**No automated tests exist in this repo.** Verifying a change means building the client/server per its README and manually exercising the client-server flow (e.g. start server, connect client, send/receive a command).

## Architecture pattern across C2 examples

Most `Command_and_Control` examples share the same basic shape even though implementations differ:
- A **server** (Python/Gunicorn or Go) that accepts client callbacks/connections and queues commands per client/agent ID.
- A **client/agent** (Go, occasionally Python via PyInstaller) that checks in on an interval, executes queued commands, and returns output.
- Where present, a **CLI** (Go, in a `cli/` subfolder) that operates against the server's HTTP API to list clients, list queued commands, and add new commands (`list-clients`, `list-commands`, `add-command`).

Channel-specific variants (`go_tcp_c2`, `multi_tcp_ex`, `python_stream_c2`) implement the same client/server command loop directly over raw TCP sockets instead of HTTP.

## Working with this code

This repo intentionally contains offensive-security proof-of-concept code (backdoor payloads, reverse shells, in-memory execution, credential/token harvesting). When extending or modifying it, keep changes scoped to the stated research/defensive-education purpose of each module — see the top-level README.
