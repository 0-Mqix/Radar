//websocket object for the go package/module i made 
class MonkeSocket {
    events 
    conn

    //the url is the http endpoint to make the connection
    //example: "localhost:8080/endpoint"
    constructor(url) {
        if (window["WebSocket"]) {
            this.events = new Map()
            this.conn = new WebSocket("ws://" + document.location.host + url)
        
            const _this = this;
            this.conn.onmessage = (e) => {
                var messages = e.data.split('\n')
                for (let i = 0; i < messages.length; i++) {
                    let data = messages[i]

                    let event = data.slice(0, data.indexOf(':') + 1)
                    let message = data.slice(data.indexOf(':') + 1)
            
                    let func = _this.events.get(event)
                    if (func != null) func(message)
                }   
            }
        }
    }

    //register an event with the event name and the function
    on(event, func) {
        this.events.set(event, func)
    }

    //registers the function that runs when the connection is made
    onOpen(func) {
        this.conn.onopen = (e) => func(e)
    }

    //registers the function that runs when the conection stops
    onClose(func) {
        this.conn.onclose = (e) => func(e)
    }

    //sends a command to the room the player is connected to with the event name and the message
    send(event, message) {
        this.conn.send(event+message)
    }
}