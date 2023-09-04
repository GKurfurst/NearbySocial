class PackedWebSocket {
    constructor(url) {
        this.ws = new window.WebSocket(url);
    }

    send(msg) {
        this.ws.send(msg);
    }

    onmessage(handler) {
        this.ws.onmessage = handler;
    }

    close() {
        this.ws.close();
    }
}

export default PackedWebSocket;
