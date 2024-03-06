// import * as redis from 'redis';

// const client = redis.createClient({
//     host: '127.0.0.1',
//     port: 6379,
// });
//     await new Promise<void>((resolve, reject) => {
//         client.on('error', (error) => {
//             console.error('Error connecting to Redis:', error);
//             reject(error);
//         });
//         client.on('ready', () => {
//             console.log('Connected to Redis');
//             resolve();
//         });
//     });
//     client.subscribe('code', (err, data) => {
//         if (err) {
//             console.error('Error subscribing to channel:', err);
//         } else {
//             console.log(data);
//         }
//     });
// }
