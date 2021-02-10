const express = require("express");
const router = express.Router();
const eventHandler = require("./handler/event");

router.post("/", eventHandler.create);
router.get("/", eventHandler.getAll);
router.get("/:id", eventHandler.get);
router.put("/:id", eventHandler.update);
router.delete("/:id", eventHandler.destroy);

module.exports = router;
