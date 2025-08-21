const amqp = require("amqplib");

const msg = { number: process.argv[2] };

connect();

async function connect() {
  try {
    const conn = await amqp.connect("amqp://localhost:5672"); 
    const channel = await conn.createChannel();
    const res = await channel.assertQueue("jobs");
    channel.sendToQueue("jobs", Buffer.from(JSON.stringify(msg)));
    console.log("Job sent successfully!");
  } catch(err) {
    console.error(err);
  }
}

