const express = require("express");
const router = express.Router();
const alumniHandler = require("./handler/alumni");

router.post("/", alumniHandler.create);
router.get("/", alumniHandler.getAll);
router.get("/:id", alumniHandler.getById);
router.put("/:id", alumniHandler.update);
router.delete("/:id", alumniHandler.destroy);

module.exports = router;
