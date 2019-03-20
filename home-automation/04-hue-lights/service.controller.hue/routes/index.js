const express = require("express");
const dao = require("../dao");

const register = (app) => {
    app.use(express.json());

    // Request logger
    app.use((req, res, next) => {
        console.log(`${req.method} ${req.originalUrl} ${JSON.stringify(req.body)}`);
        next();
    });

    app.get("/device/:deviceId", (req, res) => {
        const device = dao.findByIdentifier(req.params.deviceId);
        res.json({data: device});
    });

    app.patch("/device/:deviceId", (req, res, next) => {
        const device = dao.findByIdentifier(req.params.deviceId);
        const state = req.body;

        dao.applyState(device, state)
            .then(() => {
                res.json({data: device});
            })
            .catch(next);
    });

    app.use(function (err, req, res, next) {
        console.error(err);
        res.status(500);
        res.json({message: err.message});
    });
};

exports = module.exports = { register };