const apiAdapter = require("../../apiAdapter");
const { URL_SERVICE_CHAT } = process.env;
const api = apiAdapter(URL_SERVICE_CHAT);

module.exports = async (req, res) => {
  try {
    const id = req.params.id;
    const chat = await api.get(`/chat/postsCategory/${id}`);
    return res.json(chat.data);
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
