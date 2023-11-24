import { Elysia } from "elysia";

const app = new Elysia().get("/", () => "Hello Elysaa and everyone").listen(3000);

console.log(
  `Server is running at ${app.server?.hostname}:${app.server?.port}`
);
