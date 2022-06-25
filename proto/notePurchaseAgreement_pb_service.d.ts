// package: notepurchaseagreement
// file: notePurchaseAgreement.proto

import * as notePurchaseAgreement_pb from "./notePurchaseAgreement_pb";
import {grpc} from "@improbable-eng/grpc-web";

type NotePurchaseAgreementServiceGetNotePurchaseAgreementDoc = {
  readonly methodName: string;
  readonly service: typeof NotePurchaseAgreementService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof notePurchaseAgreement_pb.GetNotePurchaseAgreementDocRequest;
  readonly responseType: typeof notePurchaseAgreement_pb.GetNotePurchaseAgreementDocResponse;
};

export class NotePurchaseAgreementService {
  static readonly serviceName: string;
  static readonly GetNotePurchaseAgreementDoc: NotePurchaseAgreementServiceGetNotePurchaseAgreementDoc;
}

export type ServiceError = { message: string, code: number; metadata: grpc.Metadata }
export type Status = { details: string, code: number; metadata: grpc.Metadata }

interface UnaryResponse {
  cancel(): void;
}
interface ResponseStream<T> {
  cancel(): void;
  on(type: 'data', handler: (message: T) => void): ResponseStream<T>;
  on(type: 'end', handler: (status?: Status) => void): ResponseStream<T>;
  on(type: 'status', handler: (status: Status) => void): ResponseStream<T>;
}
interface RequestStream<T> {
  write(message: T): RequestStream<T>;
  end(): void;
  cancel(): void;
  on(type: 'end', handler: (status?: Status) => void): RequestStream<T>;
  on(type: 'status', handler: (status: Status) => void): RequestStream<T>;
}
interface BidirectionalStream<ReqT, ResT> {
  write(message: ReqT): BidirectionalStream<ReqT, ResT>;
  end(): void;
  cancel(): void;
  on(type: 'data', handler: (message: ResT) => void): BidirectionalStream<ReqT, ResT>;
  on(type: 'end', handler: (status?: Status) => void): BidirectionalStream<ReqT, ResT>;
  on(type: 'status', handler: (status: Status) => void): BidirectionalStream<ReqT, ResT>;
}

export class NotePurchaseAgreementServiceClient {
  readonly serviceHost: string;

  constructor(serviceHost: string, options?: grpc.RpcOptions);
  getNotePurchaseAgreementDoc(
    requestMessage: notePurchaseAgreement_pb.GetNotePurchaseAgreementDocRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: notePurchaseAgreement_pb.GetNotePurchaseAgreementDocResponse|null) => void
  ): UnaryResponse;
  getNotePurchaseAgreementDoc(
    requestMessage: notePurchaseAgreement_pb.GetNotePurchaseAgreementDocRequest,
    callback: (error: ServiceError|null, responseMessage: notePurchaseAgreement_pb.GetNotePurchaseAgreementDocResponse|null) => void
  ): UnaryResponse;
}

