import { expect } from "chai";
import {
    WatchConnectionStatusRequest,
    WatchConnectionStatusResponse,
} from "./connection";

describe("grpc integration", () => {
    describe("configs", () => {
        it("should getConfigs", async () => {
            const configs = await getConfigs();

            expect(configs).to.not.equal(null);
            expect(configs.length).to.not.equal(0);

            const ids = configs.map(it => typeof it.id);

            expect(ids).to.have.members(["number"]);
        });
    });

    describe("healthcheck", () => {
        it("should recieve connection status", async () => {
            const configs = await getConfigs();

            expect(configs).to.not.equal(null);
            expect(configs.length).to.not.equal(0);

            const ids = configs.map(it => typeof it.id);

            expect(ids).to.have.members(["number"]);
        });
    });
});