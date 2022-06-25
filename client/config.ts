import { grpc } from "@improbable-eng/grpc-web";

// Import code-generated data structures.
import { NotePurchaseAgreementServiceClient } from "../proto/notePurchaseAgreement_pb_service";
import { 
  NotePuchaseAgreement, 
  GetNotePurchaseAgreementDocRequest, 
  GetNotePurchaseAgreementDocResponse,
} from "../proto/notePurchaseAgreement_pb";

import { NodeHttpTransport } from '@improbable-eng/grpc-web-node-http-transport';

const transport = NodeHttpTransport() as any;

const client = new NotePurchaseAgreementServiceClient("http://localhost:8020", {
  transport
});

export function getConfigs(): Promise<NotePuchaseAgreement.AsObject[]> {
  const request = new GetNotePurchaseAgreementDocRequest();

  const metadata = new grpc.Metadata();

  return new Promise((resolve, reject) => {
    client.getNotePurchaseAgreementDoc(request, metadata, function (err, response) {

      if (!!err) {
        reject(err);
      }

      const results: any = response?.toObject();

      if (results) {
        resolve(results.configsList);
        return;
      }

      resolve([]);
    });
  });
}