import { grpc } from "@improbable-eng/grpc-web";

// Import code-generated data structures.
import { ConnectionsServiceClient } from "../proto/connection_pb_service";
import {
    ConnectionStatus,
    WatchConnectionStatusRequest,
    WatchConnectionStatusResponse,
} from "../proto/connection_pb";
import {
} from "./connection";

import { NodeHttpTransport } from '@improbable-eng/grpc-web-node-http-transport';

const transport = NodeHttpTransport() as any;

const client = new ConnectionsServiceClient("http://localhost:8020", {
    transport
});

export function watchConnectionStatus(watchData: Function, watchStatus: Function, watchEnd: Function): Promise<ConnectionStatus.AsObject | null> {
    const request = new WatchConnectionStatusRequest();

    const metadata = new grpc.Metadata();

    const connectionResponseStream = client.watchConnectionStatus(request, metadata);

    return new Promise((resolve, reject) => {

        connectionResponseStream.on("status", (status) => {
            console.log("watchConnectionStatus - stream status: ", status)
        });
        
        connectionResponseStream.on("data", function (response) {

            const results = response?.toObject();

            if (results && typeof results.connectionstatus !== "undefined") {
                resolve(results.connectionstatus);
                return;
            }

            resolve(null);
        });

    });
}