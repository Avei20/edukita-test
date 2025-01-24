FROM oven/bun:latest

# create & set working directory
WORKDIR /app

# copy source files
COPY . .

# install dependencies
RUN bun install

# start app
RUN bun run build
EXPOSE 3000
CMD bun run start
