import { Elysia, t } from "elysia";
import Redis from "ioredis"

export const app = new Elysia()
export const redis = new Redis({
    port: 6379,
    host: "redis"
})

app.get("/", () => "Welcome to this blog server bitch boy hhhhhhhhhh")

app.get("/todo", async () => {
    const allTodos = await redis.get("todo") || []
    return allTodos
})

app.post("/todo", async ({ body, set }) => {
    try {
        const todos = await redis.get("todo")
        const newtodos = [ ...todos! , body.content]
        const response = await redis.set("todo", JSON.stringify(newtodos))
        set.status = 200
        return response
    } catch (err) {
        set.status = 500
        return err
    }

}, {
    body: t.Object({
        content: t.String(),
    })
})

app.delete("/todo" , async () => {
    try{
        const res = await redis.del("todo")
        return res
    }catch(err) {
        return err
    }
})

app.listen(3000, () => {
    console.log(
        `Server is running at ${app.server?.hostname}:${app.server?.port}`
    );
});

