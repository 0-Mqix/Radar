const ws = new MonkeSocket("/ws")
const click = document.getElementById("click")
const canvas = document.getElementById('map')
const context = canvas.getContext('2d')

const width = canvas.width
const height = canvas.height

const circlePoints = new Map()
const squarePoints = new Map()

//registers websocket event for "pointCircle:"
ws.on("pointCircle:", (message) => {
  const data = JSON.parse(message)
  circlePoints.set(data.angle, {
    x : data.x,
    y : data.y
  })
  drawScreen()
  drawPoint(context, data.x, data.y, 'blue', 8);
})

//registers websocket event for "pointSquare:"
ws.on("pointSquare:", (message) => {
  const data = JSON.parse(message)
  squarePoints.set(data.angle, {
    x : data.x,
    y : data.y
  })
  drawScreen()
  drawPoint(context, data.x, data.y, 'blue', 8);
})

//it calling this cleans the screen renders all the sqaure and the circle points
function drawScreen() {
  context.clearRect(0, 0, width, height);
  
  circlePoints.forEach((value, key, _) => {
    drawPoint(context, value.x, value.y, 'purple', 5);
  })

  squarePoints.forEach((value, key, _) => {
    drawPoint(context, value.x, value.y, 'green', 5);
  })
  
  drawPoint(context, 0, 0, 'red', 10);
}

//draws a circle with the given cordinates, color, and size
function drawPoint(context, x, y, color, size) {
  context.beginPath();
  context.fillStyle = color;
  context.arc(x + width/2, y + height/2, size, 0 * Math.PI, 2 * Math.PI);
  context.fill()
}