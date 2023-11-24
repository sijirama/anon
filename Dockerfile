FROM oven/bun

WORKDIR /app

COPY package.json .
COPY bun.lockb .

RUN bun install

COPY src src
COPY tsconfig.json .

ENV NODE_ENV development 
CMD ["bun", "src/index.ts"]

EXPOSE 3000
