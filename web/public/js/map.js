const ws = new MonkeSocket("/ws")
const click = document.getElementById("click")
const canvas = document.getElementById('map')
const context = canvas.getContext('2d')

const width = canvas.width
const height = canvas.height

const map = new Map()

drawPoint(context, 0, 0, 'red', 5)

ws.on("point:", (message) => {
  const data = JSON.parse(message)
  drawPoint(context, data.x, data.y, 'blue', 2);
})

// That's how you define the value of a pixel
function drawPoint(context, x, y, color, size) {
  context.beginPath();
  context.fillStyle = color;
  context.arc(x + width/2, y + height/2, size, 0 * Math.PI, 2 * Math.PI);
  context.fill()
}