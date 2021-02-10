const apiAdapter = require("../../apiAdapter");
const { URL_SERVICE_EVENT } = process.env;
const api = apiAdapter(URL_SERVICE_EVENT);

module.exports = async (req, res) => {
  try {
    const id = req.params.id;
    const category = await api.get(`/api/category/${id}`);
    return res.json(category.data);
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
