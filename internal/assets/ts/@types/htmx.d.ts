declare global {
    interface Window {
        htmx: {
            process: (element: HTMLElement) => void;
            // Add more methods if needed (e.g., trigger, find, etc.)
        };
    }
}

export { };
