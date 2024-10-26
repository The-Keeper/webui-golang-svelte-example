// See https://kit.svelte.dev/docs/types#app
// for information about these interfaces
declare global {
	namespace App {
		// interface Error {}
		// interface Locals {}
		// interface PageData {}
		// interface Platform {}
	}

	const webui: WebUI;

	interface WebUI {
		isConnected(): Promise<bool>;
	}
}

export {};
