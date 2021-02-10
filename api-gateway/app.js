require("dotenv").config();

const express = require("express");
const path = require("path");
const cookieParser = require("cookie-parser");
const logger = require("morgan");

const indexRouter = require("./routes/index");
const mediaRouter = require("./routes/media");
const eventRouter = require("./routes/event");
const chatRouter = require("./routes/chat");
const usersRouter = require("./routes/users");
const categoryRouter = require("./routes/category");
const myFavoritesRouter = require("./routes/myfavorites");
const refreshTokenRouter = require("./routes/refreshToken");

const verifyToken = require("./middlewares/verifyToken");

const app = express();

app.use(logger("dev"));
app.use(express.json({ limit: "50mb" }));
app.use(express.urlencoded({ extended: false, limit: "50mb" }));
app.use(cookieParser());
app.use(express.static(path.join(__dirname, "public")));

app.use("/", indexRouter);
app.use("/users", usersRouter);
app.use("/media", mediaRouter);
app.use("/event", eventRouter);
app.use("/chat", chatRouter);
app.use("/category", categoryRouter);
app.use("/myfavorites", verifyToken, myFavoritesRouter);

app.use("/refresh-token", refreshTokenRouter);

module.exports = app;
