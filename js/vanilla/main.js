const logLine = '10.244.0.20 - - [11/Dec/2024:10:44:24 +0000] "GET /wp-login.php HTTP/1.1" 200 4445';
const regex = '/^(?<ipAddress>\d+\.\d+\.\d+\.\d+)\s+-\s+-\s+\[(?<timestamp>\d{2}\/\w+\/\d{4}:\d{2}:\d{2}:\d{2}\s+\+\d{4})\]\s+"(?<method>GET|POST|PUT|DELETE)\s+(?<path>[^"\s]*)\s+(?<protocol>HTTP\/[\d.]+)"\s+(?<statusCode>\d{3})\s+(?<responseSize>\d+)$/';
const result = logLine.match(new RegExp(regex));
console.log('Does it match?', !!result);
if (result) {
    console.log('Captured groups:', result.groups);
}
