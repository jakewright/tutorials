const axios = require("axios");
const hueClient = require("../api/hueClient");

let devices = [];

const findByHueId = (hueId) => {
    return devices.find(device => device.attributes["hue_id"] === hueId);
};

const findByIdentifier = (identifier) => {
    return devices.find(device => device.identifier === identifier);
};

const fetchAllState = async () => {
    const rsp = await axios.get("http://service.registry.device/devices");
    devices = rsp.data.data;

    const hueIdToState = await hueClient.fetchAllState();

    for (const hueId in hueIdToState) {
        const device = findByHueId(hueId);
        if (device === undefined) continue;

        Object.assign(device, hueIdToState[hueId]);
    }
};

const applyState = async (device, state) => {
    const newState = await hueClient.applyState(device.attributes["hue_id"], state);
    Object.assign(device, newState);
};

exports = module.exports = { findByIdentifier, fetchAllState, applyState };
