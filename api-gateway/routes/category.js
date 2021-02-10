const express = require("express");
const router = express.Router();
const categoryHandler = require("./handler/category");

router.get("/", categoryHandler.getAll);
router.get("/:id", categoryHandler.get);
router.post("/", categoryHandler.create);
router.put("/:id", categoryHandler.update);
router.delete("/:id", categoryHandler.destroy);

module.exports = router;
