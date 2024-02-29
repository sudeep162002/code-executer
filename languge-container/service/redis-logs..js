const redis = require('redis');



export async function connect(){

    const client = redis.createClient(6379,'127.0.0.1');
   await client.connect();
   
     
   
   
   
    client.subscribe('code',(data)=>{
     console.log(data)
   });

}


