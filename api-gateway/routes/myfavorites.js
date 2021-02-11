const express = require("express");
const router = express.Router();
const myFavoritesHandler = require("./handler/myfavorites");

router.get("/", myFavoritesHandler.get);
router.post("/", myFavoritesHandler.create);
router.delete("/:id", myFavoritesHandler.destroy);

module.exports = router;
