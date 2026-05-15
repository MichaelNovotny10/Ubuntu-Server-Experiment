# Ubuntu Server Experiment

> My personal home‑lab journey – running an Ubuntu server with a Go backend, Cloudflare Tunnel, and a simple static site for changelog and progress tracking.

🔗 **Live site:** [server.michaelnovotny.dev](https://server.michaelnovotny.dev)

---

## 🚀 What’s inside

- **Go web server** – serves static files for `/`, `/changelog`, `/progress`
- **Ubuntu 22.04 / 24.04** – hosted on my local hardware
- **Cloudflare Tunnel** – exposes the server without opening public IP ports
- **Static HTML/CSS** – fully responsive, dark theme, mobile‑friendly
- **Systemd service** – auto‑restarts the Go binary on failure

---

## 📁 Project structure

```

.
├── main.go                – entry point, HTTP handlers
├── index/                 – root page (/, style.css, favicon)
├── changelog/             – /changelog page
├── progress/              – /progress page
├── go.mod
└── README.md

```

---

## 🛠️ Local development

1. **Clone the repo**
   ```bash
   git clone https://github.com/MichaelNovotny10/Ubuntu-Server-Experiment.git
   cd Ubuntu-Server-Experiment
   ```

1. Run the server
   ```bash
   go run main.go
   ```
   Then open http://localhost:3000
2. Build for production
   ```bash
   go build -o app
   sudo systemctl restart newserver   # (if using the systemd service)
   ```

---

📝 Changelog & progress

All notable changes and ongoing experiments are documented live at:

· /changelog
· /progress

---

🔧 Future ideas

· Having fun! (I have no idea what exactly to do right now)

---

📄 License

MIT – do whatever you want with this code.

---

Made with Love and Passion in my dorm room.
