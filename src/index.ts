import { Elysia } from "elysia";
import Redis from "ioredis"

const app = new Elysia()
export const redis = new Redis({
    port: 6379,
    host:"redis"
})

//INFO: to test if redis is working
const redisTest = async () => {
    console.log("Testing beginning")
    try {
        await redis.set("mykey", "value is the key i inserted at the beginning of the code");
        redis.get("mykey", (err, result) => {
            if (err) {
                console.error(err);
            } else {
                console.log(result); // Prints "value"
            }
        });
    } catch (e) {
        console.error(e)
    }
    console.log("Testing end")
}

await redisTest()

//INFO: test ended!

app.get("/", () => "Welcome to this blog server bitch boy hhhhhhhhhh")
app.listen(3000, () => {
    console.log(
        `Server is running at ${app.server?.hostname}:${app.server?.port}`
    );
});

