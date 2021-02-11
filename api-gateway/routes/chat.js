const express = require("express");
const router = express.Router();
const chatHandler = require("./handler/chat");

router.post("/", chatHandler.create);
router.get("/", chatHandler.getAll);
router.get("/:id", chatHandler.get);
router.get("/filtered/:id", chatHandler.getCategory);

module.exports = router;
