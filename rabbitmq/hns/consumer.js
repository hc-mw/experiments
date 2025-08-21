const amqp = require("amqplib")
connect();
async function connect() {
  try {
    const conn = await amqp.connect("amqp://localhost:5672");
    const channel = await conn.createChannel();
    const res = await channel.assertQueue("jobs");
    channel.consume("jobs", message => {
      console.log(message.content.toString())
      channel.ack(message);
    })
    console.log("waiting for messages...")
  } catch(err) {
    console.error(err);
  }
}
