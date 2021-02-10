const apiAdapter = require("../../apiAdapter");
const { URL_SERVICE_EVENT, HOSTNAME } = process.env;
const api = apiAdapter(URL_SERVICE_EVENT);

module.exports = async (req, res) => {
  try {
    const event = await api.get("/api/event", {
      params: {
        ...req.query,
      },
    });

    const eventData = event.data;
    const firstPage = eventData.data.first_page_url.split("?").pop();
    const lastPage = eventData.data.last_page_url.split("?").pop();

    eventData.data.first_page_url = `${HOSTNAME}/event?${firstPage}`;
    eventData.data.last_page_url = `${HOSTNAME}/event?${lastPage}`;

    if (eventData.data.next_page_url) {
      const nextPage = eventData.data.next_page_url.split("?").pop();
      eventData.data.next_page_url = `${HOSTNAME}/event?${nextPage}`;
    }

    if (eventData.data.prev_page_url) {
      const prevPage = eventData.data.prev_page_url.split("?").pop();
      eventData.data.prev_page_url = `${HOSTNAME}/event?${prevPage}`;
    }

    eventData.data.path = `${HOSTNAME}/event`;

    return res.json(eventData);
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
