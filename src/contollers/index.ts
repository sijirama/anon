import {Handler} from "elysia"
import { redis } from ".."


export const getAllTodos = async (handler:Handler) => {
    const allTodos = await redis.get("todo")
    return allTodos
}
