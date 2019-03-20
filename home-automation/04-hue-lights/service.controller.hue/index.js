const axios = require("axios");
const express = require("express");
const dao = require("./dao");
const hueClient = require("./api/hueClient");
const routes = require("./routes");

const port = 80;

axios.get("http://service.config/read/service.controller.hue")
    .then(rsp => {
        hueClient.host = rsp.data.hueBridge.host;
        hueClient.username = rsp.data.hueBridge.username;

        return dao.fetchAllState();
    })
    .then(() => {
        const app = express();
        routes.register(app);
        app.listen(port, () => console.log(`Listening on port ${port}`));
    })
    .catch(err => {
        console.error("Error initialising service", err);
    });









