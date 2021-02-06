const jwt = require("jsonwebtoken");
const { JWT_SECRET } = process.env;

module.exports = async (req, res, next) => {
  const bearerHeader = req.headers["authorization"];
  console.log("token :", bearerHeader);
  if (bearerHeader) {
    const bearer = bearerHeader.split(" ");
    const token = bearer[1];

    jwt.verify(token, JWT_SECRET, function (err, decoded) {
      if (err) {
        return res.status(403).json({ message: err.message });
      }
      req.user = decoded;

      return next();
    });
  } else {
    res.status(403).send({ auth: false, message: "No token provided." });
  }
};
