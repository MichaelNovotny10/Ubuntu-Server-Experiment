# Ubuntu Server Experiment

My personal home-lab journey: a Go HTTP server running on my Ubuntu Server laptop in my dorm room, exposed to the internet via Cloudflare Tunnel.

**Live site:** [server.michaelnovotny.dev](https://server.michaelnovotny.dev)

---

## What's inside

- **Go web server** — serves static files, battery status endpoint, graceful shutdown on SIGTERM
- **systemd service** — keeps the server running, auto-restarts on failure
- **Makefile** — build and deploy shortcuts

---

## Project structure

```
.
├── app.go                  — HTTP server, routing, graceful shutdown
├── deploy/
│   └── newserver.service   — systemd unit file (lives at /etc/systemd/system/)
├── Makefile                — build and deploy shortcuts
├── index/                  — static files served at /
├── changelog/              — static files served at /changelog
├── progress/               — static files served at /progress
├── projects/               — static files served at /projects
├── battery/                — frontend for the /battery endpoint (live acpi status)
└── go.mod
```

---

## First time run

Requirements: [Go](https://go.dev/dl/) and `make` installed.
Install Go:
```bash
sudo apt update && sudo apt install golang-go
```
Install make:
```bash
sudo apt update && sudo apt install make
```
Run it:
```bash
git clone https://github.com/MichaelNovotny10/Ubuntu-Server-Experiment.git
cd Ubuntu-Server-Experiment
make build
./app
```

Then open http://localhost:3000 in your browser.

If you don't want to install `make`:

```bash
go build -o app && ./app
```

---

## How to remove

If you set it up as a systemd service, stop and clean that up first:

```bash
sudo systemctl stop newserver
sudo systemctl disable newserver
sudo rm /etc/systemd/system/newserver.service
sudo systemctl daemon-reload
```

Then delete the project folder:

```bash
rm -rf /path/to/Ubuntu-Server-Experiment
```

---

## Local development

```bash
git clone https://github.com/MichaelNovotny10/Ubuntu-Server-Experiment.git
cd Ubuntu-Server-Experiment
go run app.go
```

Then open http://localhost:3000.

---

## Building and deploying

```bash
make deploy
```

This builds the binary and restarts the systemd service.

To install or update the systemd unit:

```bash
sudo cp deploy/newserver.service /etc/systemd/system/newserver.service
sudo systemctl daemon-reload
```

---

## License

MIT — do whatever you want with this code.

---

Made with love and passion in my dorm room.
