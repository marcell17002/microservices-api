const apiAdapter = require("../../apiAdapter");
const { URL_SERVICE_ALUMNI } = process.env;
const api = apiAdapter(URL_SERVICE_ALUMNI);

module.exports = async (req, res) => {
  try {
    const info = await api.post("/alumni-information", req.body);
    return res.json(info.data);
  } catch (error) {
    if (error.code === "ECONNREFUSED") {
      return res.status(500).json({
        status: "error",
        message: "service unavailable",
      });
    }
    const { status, data } = error.response;
    return res.status(status).json(data);
  }
};
