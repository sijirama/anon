import { t } from 'elysia'
import { app } from '..'
import { redis } from '..'

//TODO: get all todos
app.get("/todo", async () => {
    await redis.get("todo")
})

//TODO: update todos
//TODO: delete todos

//TODO: create todo
app.post("/todo", async ({ body, set }) => {
    try {
        const allTodos = await redis.get("todo")
        const existingTodos = JSON.parse(allTodos!) || [];
        existingTodos.push(body.content);
        await redis.set("todo", JSON.stringify(existingTodos));
        set.status=200
        return existingTodos
    } catch (e) { 
        set.status = 500
        return e
    }
}, {
    body: t.Object({
        content: t.String(),
    })
})
