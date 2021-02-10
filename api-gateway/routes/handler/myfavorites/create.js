const apiAdapter = require("../../apiAdapter");
const { URL_SERVICE_EVENT } = process.env;
const api = apiAdapter(URL_SERVICE_EVENT);

module.exports = async (req, res) => {
  try {
    const userId = req.user.data.id;
    const eventId = req.body.event_id;
    const myfavorites = await api.post("/api/myfavorites", {
      user_id: userId,
      event_id: eventId,
    });
    return res.json(myfavorites.data);
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
