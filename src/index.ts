import { Elysia } from "elysia";

const app = new Elysia()
app.get("/", () => "Welcome to this blog server bitch boy hhhhhhhhhh")
app.listen(3000, () => {
    console.log(
        `Server is running at ${app.server?.hostname}:${app.server?.port}`
    );
});

