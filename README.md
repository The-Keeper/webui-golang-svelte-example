Building desktop app UIs using awesome [WebUI library](https://github.com/webui-dev/go-webui) and Svelte.

## Requirements
- Go 1.22 or later
- Bun

## Installation and running

Clone the repository:
```sh
git clone https://github.com/The-Keeper/webui-golang-svelte-example.git example --recursive
```

Install UI:
```sh
cd example/ui
bun install
```

Debug the UI:
```sh
bun run dev
```

Build and run the app:
```sh
bun run build
cd ..
go run .
```

