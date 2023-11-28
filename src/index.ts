import { Elysia, t } from "elysia";
import Redis from "ioredis"

export const app = new Elysia()
export const redis = new Redis({
    port: 6379,
    host: "redis"
})

app.get("/", async ({set}) => {

    set.status = 200
    return await redis.get("")
} )

app.get("/todo", async () => {
    const allTodos = await redis.get("todo") || []
    return allTodos
})

app.post("/todo", async ({ body, set }) => {
    try {
        // Retrieve existing todos from Redis and parse them
        const existingTodosStr = await redis.get("todo");
        const existingTodos = existingTodosStr ? JSON.parse(existingTodosStr) : [];

        // Add the new todo to the existing todos array
        existingTodos.push(body.content);

        // Update Redis with the updated todos
        await redis.set("todo", JSON.stringify(existingTodos));

        set.status = 200;
        return { message: "Todo added successfully" };
    } catch (err) {
        console.error(err);
        set.status = 500;
        return { error: "Internal Server Error" };
    }

}, {
    body: t.Object({
        content: t.String(),
    })
})

app.delete("/todo", async () => {
    try {
        const res = await redis.del("todo")
        return res
    } catch (err) {
        return err
    }
})

app.listen(3000, () => {
    console.log(
        `Bun Server is running at ${app.server?.hostname}:${app.server?.port}`
    );
});

