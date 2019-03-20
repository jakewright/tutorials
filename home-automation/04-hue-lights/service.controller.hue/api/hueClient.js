const axios = require("axios");

class HueClient {
    get hueUrl() {
        return `${this.host}/api/${this.username}`;
    }

    async fetchAllState() {
        const rsp = await axios.get(`${this.hueUrl}/lights`);

        const lights = {};

        for (const hueId in rsp.data) {
            lights[hueId] = {
                power: rsp.data[hueId].state.on,
                brightness: rsp.data[hueId].state.bri
            };
        }

        return lights;
    }

    async fetchState(hueId) {
        const rsp = await axios.get(`${this.hueUrl}/lights/${hueId}`);

        return {
            power: rsp.data.state.on,
            brightness: rsp.data.state.bri
        };
    }

    async applyState(hueId, state) {
        const hueState = {
            on: state.power,
            bri: state.brightness,
        };

        await axios.put(`${this.hueUrl}/lights/${hueId}/state`, hueState);
        return this.fetchState(hueId);
    }
}

const hueClient = new HueClient();
exports = module.exports = hueClient;